import {Inject, Injectable, Injector} from '@angular/core';
import {
  ActivatedRoute,
  ActivatedRouteSnapshot,
  CanActivate,
  Params,
  Router,
  RouterStateSnapshot,
  UrlTree
} from '@angular/router';
import {StartupService} from '@core';
import {ReuseTabService} from '@delon/abc';
import {DA_SERVICE_TOKEN, ITokenService} from '@delon/auth';
import {SettingsService, _HttpClient} from '@delon/theme';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class LoginGuard implements CanActivate {
  constructor(
    private http: _HttpClient,
    private activeRoute: ActivatedRoute,
    private reuseTabService: ReuseTabService,
    @Inject(DA_SERVICE_TOKEN) private tokenService: ITokenService,
    private startupSrv: StartupService,
    private settingService: SettingsService,
    private injector: Injector,
    private router: Router,
  ) {
  }

  private goTo(url: string) {
    setTimeout(() => this.injector.get(Router).navigateByUrl(url));
  }

  // tslint:disable-next-line:max-line-length
  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot,
  ): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    let paramToken = window.location.href.split('token=')[1];
    if (paramToken != undefined) {
      this.tokenService.set({token: paramToken});
      window.location.href = '/';
      return true;
    }
    const token = window.localStorage.getItem('_token');
    this.reuseTabService.clear();
    this.http.get(`/api/v1/my/profile`).subscribe((res: any) => {
      this.settingService.setUser(res?.data);
    });
    return true;
  }
}
