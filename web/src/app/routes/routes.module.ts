import {NgModule} from '@angular/core';

import {SharedModule} from '@shared';
import {CallbackComponent} from './callback/callback.component';
import {UserLockComponent} from './passport/lock/lock.component';
import {UserLoginComponent} from './passport/login/login.component';
import {UserRegisterResultComponent} from './passport/register-result/register-result.component';
import {UserRegisterComponent} from './passport/register/register.component';
import {RouteRoutingModule} from './routes-routing.module';

const COMPONENTS = [
  UserLoginComponent,
  UserRegisterComponent,
  UserRegisterResultComponent,
  CallbackComponent,
  UserLockComponent,
];
const COMPONENTS_NO_ROUTE = [];

@NgModule({
  imports: [SharedModule, RouteRoutingModule],
  declarations: [...COMPONENTS, ...COMPONENTS_NO_ROUTE],
  entryComponents: COMPONENTS_NO_ROUTE,
})
export class RoutesModule {
}
