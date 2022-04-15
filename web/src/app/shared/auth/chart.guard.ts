import { Inject, Injectable, Injector } from '@angular/core';
import { ActivatedRoute, ActivatedRouteSnapshot, CanActivate, Params, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { StartupService } from '@core';
import { ReuseTabService } from '@delon/abc';
import { DA_SERVICE_TOKEN, ITokenService } from '@delon/auth';
import { SettingsService, _HttpClient } from '@delon/theme';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ChartGuard implements CanActivate {
  constructor(private injector: Injector) {}

  private goTo(url: string) {
    setTimeout(() => this.injector.get(Router).navigateByUrl(url));
  }
  // tslint:disable-next-line:max-line-length
  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot,
  ): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    return true;
  }
}
