import {Component, OnInit} from '@angular/core';
import {_HttpClient} from '@delon/theme';
import {NzMessageService} from 'ng-zorro-antd/message';
import {NzDrawerRef} from 'ng-zorro-antd';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'task-edit',
  templateUrl: './task-edit.component.html',
})
export class TaskEditComponent implements OnInit {

  // params
  record = null;

  // form
  taskEditForm: FormGroup;
  editorOptions = {
    theme: 'vs-dark',
    language: 'shell',
    minimap: {
      enabled: false
    },
    scrollBeyondLastLine: false,
    mouseWheelZoom: true,
    links: true,
  };


  constructor(
    private fb: FormBuilder,
    private msgSrv: NzMessageService,
    public http: _HttpClient,
    private ref: NzDrawerRef,
  ) {
  }

  ngOnInit(): void {
    this.taskEditForm = this.fb.group({
      name: [null, [Validators.required]],
      bash_content: [null, [Validators.required]],
    });
    if (this.record) {
      console.log(this.record)
      this.taskEditForm.patchValue({
        name: this.record.name,
        bash_content: this.record.bash
      })
    }
  }

  submitForm() {
    let client = this.http.post(`/api/v1/tasks`, this.taskEditForm.value)
    if (this.record != null) {
      client = this.http.put(`/api/v1/tasks/${this.record.id}`, this.taskEditForm.value)
    }
    client.subscribe(() => {
      this.msgSrv.success('Saved')
      this.ref.close(true)
    })
  }

  close() {
    this.ref.close(true)
  }
}
