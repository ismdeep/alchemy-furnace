import {NgModule} from '@angular/core';
import {SharedModule} from '@shared';
import {MonacoEditorModule} from 'ngx-monaco-editor';
import {SettingRoutingModule} from './setting-routing.module';
import {NzNoAnimationModule, NzTreeModule} from "ng-zorro-antd";
import {SettingComponent} from "./pages/setting.component";
import {TokenSettingCardComponent} from "./components/token-setting-card.component";
import {TokenEditFormComponent} from "./components/token-edit-form.component";

const COMPONENTS = [
  SettingComponent
];

const COMPONENTS_NO_ROUTE = [
  TokenSettingCardComponent,
  TokenEditFormComponent,
];

@NgModule({
  imports: [SharedModule, SettingRoutingModule, MonacoEditorModule.forRoot(), NzTreeModule, NzNoAnimationModule],
  declarations: [...COMPONENTS, ...COMPONENTS_NO_ROUTE],
  entryComponents: COMPONENTS_NO_ROUTE,
})
export class SettingModule {
}
