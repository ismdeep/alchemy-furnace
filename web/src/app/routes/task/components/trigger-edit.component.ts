import {Component, OnInit, ViewChild} from '@angular/core';
import {SFComponent, SFSchema, SFTextareaWidgetSchema, SFUISchema} from '@delon/form';
import {_HttpClient} from '@delon/theme';
import {NzMessageService} from 'ng-zorro-antd/message';
import {NzModalRef} from 'ng-zorro-antd/modal';

@Component({
  selector: 'trigger-edit',
  templateUrl: './trigger-edit.component.html',
})
export class TriggerEditComponent implements OnInit {
  record: any = {};
  i: any;
  @ViewChild('sf') sf: SFComponent;
  schema: SFSchema = {
    properties: {
      name: {type: 'string', title: 'Trigger name'},
      cron: {type: 'string', title: 'Cron'},
      environment: {
        type: 'string',
        title: 'Environment',
        ui: {
          widget: 'textarea',
          autosize: {minRows: 2, maxRows: 4},
        } as SFTextareaWidgetSchema
      },
    },
    required: ['name'],
  };

  ui: SFUISchema = {
    '*': {},
  };

  task_id: any;
  trigger_id: any;

  constructor(private modal: NzModalRef, private msgSrv: NzMessageService, public http: _HttpClient) {
  }

  ngOnInit(): void {
    console.log(this.task_id)
    this.i = [this.record].map((e) => ({
      name: e.name,
      cron: e.cron,
      environment: e.environment,
    }))[0];
  }


  saving = false

  save(value: any) {
    this.saving = true
    let client = this.http.post(`/api/v1/tasks/${this.task_id}/triggers`, value)
    if (this.trigger_id != 0) {
      client = this.http.put(`/api/v1/tasks/${this.task_id}/triggers/${this.trigger_id}`, value)
    }
    client.subscribe(() => {
      this.msgSrv.success('Saved');
      this.modal.close(true);
      this.saving = false
    }, () => {
      this.saving = false
    });
  }


  close() {
    this.modal.destroy();
  }
}
