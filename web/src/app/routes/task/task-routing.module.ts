import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {TaskListComponent} from "./pages/task-list.component";
import {TaskDetailComponent} from "./pages/task-detail.component";
import {RunDetailComponent} from "./pages/run-detail.component";

const routes: Routes = [
  {path: '', component: TaskListComponent},
  {path: ':id', component: TaskDetailComponent},
  {path: ':id/runs/:run_id', component: RunDetailComponent},
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class TaskRoutingModule {
}
