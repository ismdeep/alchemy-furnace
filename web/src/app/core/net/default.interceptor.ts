import {
  HttpClient,
  HttpErrorResponse,
  HttpEvent,
  HttpHandler,
  HttpInterceptor,
  HttpRequest,
  HttpResponse,
  HttpResponseBase,
} from '@angular/common/http';
import {Injectable, Injector, Inject} from '@angular/core';
import {Router} from '@angular/router';
import {DA_SERVICE_TOKEN, ITokenService} from '@delon/auth';
import {_HttpClient} from '@delon/theme';
import {NzMessageService, NzModalRef, NzModalService, NzNotificationService} from 'ng-zorro-antd/';
import {Observable, of, throwError} from 'rxjs';
import {catchError, debounceTime, mergeMap} from 'rxjs/operators';
import {StoreService} from 'src/app/shared/store/store.service';

const CODE_MESSAGE = {
  200: '服务器成功返回请求的数据。',
  201: '新建或修改数据成功。',
  202: '一个请求已经进入后台排队（异步任务）。',
  204: '删除数据成功。',
  400: '发出的请求有错误，服务器没有进行新建或修改数据的操作。',
  401: '用户没有权限（令牌、用户名、密码错误）。',
  403: '用户得到授权，但是访问是被禁止的。',
  404: '发出的请求针对的是不存在的记录，服务器没有进行操作。',
  406: '请求的格式不可得。',
  410: '请求的资源被永久删除，且不会再得到的。',
  422: '当创建一个对象时，发生一个验证错误。',
  500: '服务器发生错误，请检查服务器。',
  502: '网关错误。',
  503: '服务不可用，服务器暂时过载或维护。',
  504: '网关超时。',
};

/**
 * 默认HTTP拦截器，其注册细节见 `app.module.ts`
 */
@Injectable()
export class DefaultInterceptor implements HttpInterceptor {
  constructor(
    private injector: Injector,
    private http: HttpClient,
    private msg: NzMessageService,
    private modal: NzModalService,
    public store: StoreService,
    private router: Router,
    @Inject(DA_SERVICE_TOKEN) private tokenService: ITokenService,
  ) {
  }

  private get notification(): NzNotificationService {
    return this.injector.get(NzNotificationService);
  }

  confirmModal: NzModalRef;

  private goTo(url: string) {
    setTimeout(() => this.injector.get(Router).navigateByUrl(url));
  }

  private checkStatus(ev: HttpResponseBase) {
    if ((ev?.status >= 200 && ev?.status < 300) || ev?.status === 401) {
      return;
    }
    const errText = CODE_MESSAGE[ev.status] || ev.statusText;
    this.notification.error(`请求错误 ${ev.status}: ${ev.url}`, errText);
  }

  private handleData(ev: HttpResponseBase): Observable<any> {
    // 可能会因为 `throw` 导出无法执行 `_HttpClient` 的 `end()` 操作
    if (ev.status > 0) {
      this.injector.get(_HttpClient).end();
    }

    this.checkStatus(ev);

    // 业务处理：一些通用操作
    switch (ev.status) {
      case 200:
        // 业务层级错误处理，以下是假定restful有一套统一输出格式（指不管成功与否都有相应的数据格式）情况下进行处理
        // 则以下代码片断可直接适用
        if (ev instanceof HttpResponse) {
          const body: any = ev.body;
          switch (body.code) {
            case 2:
            case 403:
              (this.injector.get(DA_SERVICE_TOKEN) as ITokenService).clear();
              const url = window.location.href;
              if (url.includes('/monitor/status')) {
                this.goTo('fullscreen/chart');
              } else {
                window.localStorage.clear();
                this.tokenService.clear();
                window.location.href = '/sign-in';
              }
              break;
            case 1:
              // 终止抛出数据 不会触发订阅
              return throwError(ev);
            case 0:
              // 重新修改 `body` 内容为 `response` 内容，对于绝大多数场景已经无须再关心业务状态码
              // return of(new HttpResponse(Object.assign(event, { body: body.response })));
              // return throwError({});
              return of(ev);
          }
        }
        break;
      case 401:
        this.notification.error(`未登录或登录已过期，请重新登录。`, ``);
        // 清空 token 信息
        (this.injector.get(DA_SERVICE_TOKEN) as ITokenService).clear();
        const url = window.location.href;
        // 此路由页面 不做跳转登录
        if (url.includes('/monitor/status')) {
          this.goTo('fullscreen/chart');
        } else {
          this.router.navigate(['/login']).then();
          window.localStorage.clear();
          this.tokenService.clear();
          window.location.href = '/api/v1/oauth/login?redirect_url=' + encodeURIComponent(url);
        }
        break;
      case 403:
      case 404:
        break;
      case 500:
        break;
      default:
        if (ev instanceof HttpErrorResponse) {
          console.warn('未可知错误，大部分是由于后端不支持CORS或无效配置引起', ev);
        }
        break;
    }
    if (ev instanceof HttpErrorResponse) {
      return throwError(ev);
    } else {
      return of(ev);
    }
  }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    // 统一加上服务端前缀
    let url = req.url;
    // 若请求地址为 ./assets或http打头 则请求完整url
    // if (url.startsWith('./assets')) {
    // } else if (!url.startsWith('https://') && !url.startsWith('http://')) {
    //   // url = url.includes('/api/v1/unit/') ? environment.SERVER_URL2 + url : environment.SERVER_URL + url;
    //   // url = environment.SERVER_URL + url;
    // }
    const newReq = req.clone({
      // 添加请求头处理
      headers: req.headers.set('Authorization', localStorage.getItem('_token') ? 'Bearer ' + JSON.parse(localStorage.getItem('_token')).token : ''),
      url,
    });
    return next.handle(newReq).pipe(
      mergeMap((event: any) => {
        // 允许统一对请求错误处理
        if (event instanceof HttpResponseBase) {
          return this.handleData(event);
        }
        // 若一切都正常，则后续操作
        return of(event);
      }),
      catchError((err: HttpErrorResponse) => this.handleData(err)),
      // delay(10000),
      // timeout(10000),
    );
  }
}
