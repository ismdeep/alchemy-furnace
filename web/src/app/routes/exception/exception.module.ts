import {NgModule} from '@angular/core';
import {SharedModule} from '@shared';
import {Exception403Component} from './403.component';
import {Exception404Component} from './404.component';
import {Exception500Component} from './500.component';
import {ChartComponent} from './chart/chart.component';
import {ExceptionRoutingModule} from './exception-routing.module';
import {ExceptionTriggerComponent} from './trigger.component';

const COMPONENTS = [Exception403Component, Exception404Component, Exception500Component, ExceptionTriggerComponent, ChartComponent];
const COMPONENTS_NO_ROUTE = [];

@NgModule({
  imports: [SharedModule, ExceptionRoutingModule],
  declarations: [...COMPONENTS, ...COMPONENTS_NO_ROUTE],
  entryComponents: COMPONENTS_NO_ROUTE,
  exports: [ChartComponent],
})
export class ExceptionModule {
}
