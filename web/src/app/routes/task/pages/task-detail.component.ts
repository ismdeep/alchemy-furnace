import {ChangeDetectorRef, Component, OnDestroy, OnInit, ViewChild, Input} from '@angular/core';
import {FormBuilder} from '@angular/forms';
import {ActivatedRoute, Router} from '@angular/router';
import {STColumn, STComponent} from '@delon/abc/st';
import {ModalHelper, _HttpClient, TitleService} from '@delon/theme';
import {NzMessageService} from 'ng-zorro-antd';
import {TriggerEditComponent} from "../components/trigger-edit.component";
import * as moment from "moment";
import {RunDetailComponent} from "./run-detail.component";

@Component({
  selector: 'task-detail',
  templateUrl: './task-detail.component.html',
})
export class TaskDetailComponent implements OnInit, OnDestroy {

  // params
  @Input()
  TaskID: string;

  constructor(
    private http: _HttpClient,
    public message: NzMessageService,
    private cdr: ChangeDetectorRef,
    private fb: FormBuilder,
    private activatedRouter: ActivatedRoute,
    private router: Router,
    private modalHelper: ModalHelper,
    public title: TitleService,
  ) {
  }


  intervalInstance = null

  ngOnInit() {
    if (!this.TaskID) {
      this.TaskID = this.activatedRouter.snapshot.params['id']
    }
    this.loadTriggerList()
    this.loadRunList()
    this.intervalInstance = setInterval(() => {
      this.loadRunList()
    }, 1000)
  }

  ngOnDestroy() {
    if (this.intervalInstance != null) {
      clearInterval(this.intervalInstance)
    }
  }

  @ViewChild('st', {static: false}) st: STComponent;
  columns: STColumn[] = [
    {title: 'Status', render: 'exit_code'},
    {title: 'Trigger', index: 'trigger_name'},
    {
      title: 'Created At', index: 'created_at',
      format: (item) => {
        return moment(item.created_at).fromNow()
      },
    },
    {
      title: 'Time Elapse',
      index: 'time_elapse_second',
      format: (item) => {
        let s = item.time_elapse_second
        if (s < 60) {
          return `${s} s`
        }

        s = (s / 60).toFixed(1)
        if (s < 60) {
          return `${s} m`
        }

        s = (s / 60).toFixed(1)
        if (s < 24) {
          return `${s} h`
        }

        s = (s / 24).toFixed(1)
        return `${s} d`
      }
    },
    {title: 'Ops', render: 'ops'}
  ];

  trigger_list = [];

  loadTriggerList() {
    this.http.get(`/api/v1/tasks/${this.TaskID}/triggers`).subscribe((res) => {
      this.trigger_list = res.data
    })
  }

  run_list = [];

  loadRunList() {
    this.http.get(`/api/v1/tasks/${this.TaskID}/runs`).subscribe((res) => {
      this.run_list = res.data.list
    })
  }

  runTrigger(e) {
    console.log(e)
    this.http.post(`/api/v1/tasks/${this.TaskID}/triggers/${e.id}/runs`).subscribe(() => {
      this.message.success("success")
    })
  }

  showLog(e) {
    this.modalHelper.create(RunDetailComponent, {id: this.TaskID, run_id: e.id}, {size: 'lg'}).subscribe(() => {})
  }

  addTrigger() {
    this.modalHelper.createStatic(TriggerEditComponent, {
      task_id: this.TaskID,
      trigger_id: 0
    }, {size: 800}).subscribe(() => {
      this.loadTriggerList()
    });
  }

  editTrigger(e) {
    this.modalHelper.createStatic(TriggerEditComponent, {
      task_id: this.TaskID,
      trigger_id: e.id,
      record: e
    }, {size: 800}).subscribe(() => {
      this.loadTriggerList()
    })
  }

  deleteTrigger(e) {
    console.log(e)
    this.http.delete(`/api/v1/tasks/${this.TaskID}/triggers/${e.id}`).subscribe(() => {
      this.loadTriggerList()
    })
  }
}
