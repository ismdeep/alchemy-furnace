import {ChangeDetectorRef, Component, OnInit, OnDestroy, forwardRef} from '@angular/core';
import {FormBuilder} from '@angular/forms';
import {Router} from '@angular/router';
import {ModalHelper, _HttpClient, DrawerHelper, TitleService} from '@delon/theme';
import {NzMessageService} from 'ng-zorro-antd';
import {tap} from 'rxjs/operators';
import {TaskEditComponent} from "../components/task-edit.component";
import {TriggerEditComponent} from "../components/trigger-edit.component";
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
  intervalInstance;
  ngOnInit() {
    this.title.setTitle('Tasks - Alchemy Furnace')
    this.getData();
    if (this.intervalInstance) {
      clearInterval(this.intervalInstance)
    }
    this.intervalInstance = setInterval(() => {
      this.getData()
    }, 1000)
  }

  ngOnDestroy() {
    if (this.intervalInstance) {
      clearInterval(this.intervalInstance)
    }
  }

  tasks = []
  tasks_loaded : boolean = false
  runs_map = {}
  getData() {
    this.loading = true;
    this.http.get(`/api/v1/tasks`).pipe(tap(() => (this.loading = false))).subscribe(
      (res) => {
        if (!this.tasks_loaded) {
          this.tasks = res.data;
          this.tasks_loaded = true
        }
        for (let i = 0; i < res.data.length; i++) {
          let task = res.data[i]
          if (!task.triggers) {
            continue
          }
          for (let j = 0; j < task.triggers.length; j++) {
            let trigger = task.triggers[j]
            this.runs_map[trigger.id] = trigger.recent_runs
          }
        }
      },
      () => {
        this.loading = false;
      },
    );
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
    this.modalHelper.create(RunDetailComponent, {id: taskID, run_id: e.id}, {size: document.body.clientWidth * 0.8}).subscribe(() => {
    })
  }

  getRunColor(status,exit_code) {
    if (status === 0 || status === 1) {
      return "yellow"
    }
    if (status === 2 && exit_code === 0) {
      return "green"
    }
    if (status === 2 && exit_code !== 0) {
      return "red"
    }
    if (status === 3) {
      return "gray"
    }
    return "gray"
  }
}
