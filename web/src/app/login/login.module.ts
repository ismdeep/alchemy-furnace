import {NgModule} from '@angular/core';
import {SharedModule} from '@shared';
import {LoginRoutingModule} from './login-routing.module';
import {IndexComponent} from './index/index.component';

const COMPONENTS = [];
const COMPONENTS_NO_ROUTE = [];

@NgModule({
  imports: [
    SharedModule,
    LoginRoutingModule
  ],
  declarations: [
    ...COMPONENTS,
    ...COMPONENTS_NO_ROUTE,
    IndexComponent
  ],
})
export class LoginModule {
}
