import {Inject, Injectable, Injector} from '@angular/core';
import {ACLService} from '@delon/acl';
import {DA_SERVICE_TOKEN, ITokenService} from '@delon/auth';
import {ALAIN_I18N_TOKEN, MenuService, SettingsService, TitleService, _HttpClient} from '@delon/theme';
import {TranslateService} from '@ngx-translate/core';
import {zip} from 'rxjs';
import {catchError} from 'rxjs/operators';
import {I18NService} from '../i18n/i18n.service';
import {NzIconService} from 'ng-zorro-antd/icon';
import {ICONS} from '../../../style-icons';
import {ICONS_AUTO} from '../../../style-icons-auto';
import {AppData} from "../app-data";

/**
 * Used for application startup
 * Generally used to get the basic data of the application, like: Menu Data, User Data, etc.
 */
@Injectable()
export class StartupService {
  constructor(
    iconSrv: NzIconService,
    private menuService: MenuService,
    private translate: TranslateService,
    @Inject(ALAIN_I18N_TOKEN) private i18n: I18NService,
    private settingService: SettingsService,
    private aclService: ACLService,
    private titleService: TitleService,
    @Inject(DA_SERVICE_TOKEN) private tokenService: ITokenService,
    private httpClient: _HttpClient,
    private injector: Injector,
  ) {
    iconSrv.addIcon(...ICONS_AUTO, ...ICONS);
  }

  private viaHttp(resolve: any, reject: any) {
    resolve(null);
  }

  load(): Promise<any> {
    // only works with promises
    // https://github.com/angular/angular/issues/15088
    return new Promise((resolve, reject) => {
      // http
      this.viaHttp(resolve, reject);
      // mock：请勿在生产环境中这么使用，viaMock 单纯只是为了模拟一些数据使脚手架一开始能正常运行
      // this.viaMockI18n(resolve, reject);
    });
  }
}
