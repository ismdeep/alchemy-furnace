import {NgModule} from '@angular/core';
import {SharedModule} from '@shared';
import {MonacoEditorModule} from 'ngx-monaco-editor';
import {TaskListComponent} from "./pages/task-list.component";
import {TaskEditComponent} from './components/task-edit.component';
import {TaskRoutingModule} from './task-routing.module';
import {NzNoAnimationModule, NzTreeModule} from "ng-zorro-antd";
import {TaskDetailComponent} from "./pages/task-detail.component";
import {RunDetailComponent} from "./pages/run-detail.component";
import {TriggerEditComponent} from "./components/trigger-edit.component";

const COMPONENTS = [
  TaskListComponent,
  TaskDetailComponent,
  RunDetailComponent,
];

const COMPONENTS_NO_ROUTE = [
  TaskEditComponent,
  TriggerEditComponent,
];

@NgModule({
  imports: [SharedModule, TaskRoutingModule, MonacoEditorModule.forRoot(), NzTreeModule, NzNoAnimationModule],
  declarations: [...COMPONENTS, ...COMPONENTS_NO_ROUTE],
  entryComponents: COMPONENTS_NO_ROUTE,
})
export class TaskModule {
}
