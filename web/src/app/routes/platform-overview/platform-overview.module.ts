import {NgModule} from '@angular/core';
import {SharedModule} from '@shared';
import {PlatformOverviewRoutingModule} from './platform-overview-routing.module';
import {PlatformOverviewIndexComponent} from './pages/platform-overview-index.component';
import {NgxEchartsModule} from 'ngx-echarts';

const COMPONENTS = [PlatformOverviewIndexComponent];
const COMPONENTS_NO_ROUTE = [PlatformOverviewIndexComponent];

@NgModule({
  imports: [SharedModule, PlatformOverviewRoutingModule, NgxEchartsModule],
  declarations: [...COMPONENTS, ...COMPONENTS_NO_ROUTE],
})
export class PlatformOverviewModule {
}
