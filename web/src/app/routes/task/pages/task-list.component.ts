import {ChangeDetectorRef, Component, OnInit, OnDestroy} from '@angular/core';
import {FormBuilder} from '@angular/forms';
import {Router} from '@angular/router';
import {ModalHelper, _HttpClient, DrawerHelper, TitleService} from '@delon/theme';
import {NzMessageService} from 'ng-zorro-antd';
import {tap} from 'rxjs/operators';
import {TaskEditComponent} from "../components/task-edit.component";
import format from 'date-fns/format';
import {TriggerEditComponent} from "../components/trigger-edit.component";
import {TaskDetailComponent} from "./task-detail.component";
import {RunDetailComponent} from "./run-detail.component";
import * as moment from "moment";

@Component({
  selector: 'tasks',
  templateUrl: './task-list.component.html',
  styleUrls: ['./task-list.component.scss']
})
export class TaskListComponent implements OnInit, OnDestroy {
  constructor(
    private http: _HttpClient,
    public message: NzMessageService,
    private cdr: ChangeDetectorRef,
    private fb: FormBuilder,
    private router: Router,
    private modalHelper: ModalHelper,
    private drawerHelper: DrawerHelper,
    public title: TitleService,
  ) {
  }

  loading = false;

  ngOnInit() {
    this.title.setTitle('Tasks - Alchemy Furnace')
    this.getData();
  }

  ngOnDestroy() {
  }

  tasks = [];

  getData() {
    this.loading = true;
    this.http.get(`/api/v1/tasks`).pipe(tap(() => (this.loading = false))).subscribe(
      (res) => {
        this.tasks = res.data;
      },
      () => {
        this.loading = false;
      },
    );
  }

  formatTime(t) {
    return format(new Date(t), 'yyyy-MM-dd  HH:mm:ss');
  }

  formatFromNow(t) {
    return moment(t).fromNow()
  }

  create() {
    this.drawerHelper.create('Create', TaskEditComponent, {}, {size: document.body.clientWidth * 0.618}).subscribe(() => {
      this.getData();
    });
  }

  editTask(item) {
    this.drawerHelper.create('Edit', TaskEditComponent, {record: item}, {size: document.body.clientWidth * 0.618}).subscribe(() => {
      this.getData();
    });
  }

  deleteTask(item) {
    this.http.delete(`/api/v1/tasks/${item.id}`).subscribe(() => {
      this.message.success('Deleted');
      this.getData();
    });
  }

  showTaskDetail(item) {
    this.modalHelper.create(TaskDetailComponent, {id: item.id}).subscribe(() => {
    })
  }

  editTrigger(taskInfo, item) {
    this.modalHelper.create(TriggerEditComponent, {
      task_id: taskInfo.id,
      trigger_id: item.id,
      record: item
    }, {size: 800}).subscribe(() => {
      this.getData()
    })
  }

  addTrigger(task) {
    this.modalHelper.create(TriggerEditComponent, {
      task_id: task.id,
      trigger_id: 0
    }, {size: 800}).subscribe(() => {
      this.getData()
    });
  }

  deleteTrigger(taskInfo, item) {
    this.http.delete(`/api/v1/tasks/${taskInfo.id}/triggers/${item.id}`).subscribe(() => {
      this.getData()
    })
  }

  runTrigger(taskInfo, e) {
    this.http.post(`/api/v1/tasks/${taskInfo.id}/triggers/${e.id}/runs`).subscribe(() => {
      this.message.success("Success")
    })
  }

  showLog(taskID, e) {
    this.modalHelper.create(RunDetailComponent, {id: taskID, run_id: e.id}, {size: 'lg'}).subscribe(() => {
    })
  }
}
