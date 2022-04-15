import {Component, OnInit} from '@angular/core';
import {_HttpClient} from '@delon/theme';
import {NzMessageService} from 'ng-zorro-antd/message';
import {NzDrawerRef} from 'ng-zorro-antd';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'node-edit-form',
  templateUrl: './node-edit-form.component.html',
})
export class NodeEditFormComponent implements OnInit {

  // params
  record = null;

  // form
  nodeEditForm: FormGroup;
  editorOptions = {
    theme: 'vs-dark',
    language: 'text',
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
    this.nodeEditForm = this.fb.group({
      name: [null, [Validators.required]],
      host: [null, [Validators.required]],
      port: [22, [Validators.required]],
      username: [null, [Validators.required]],
      ssh_key: [null],
    });
    if (this.record) {
      console.log(this.record)
      this.nodeEditForm.patchValue(this.record)
    }
  }

  submitForm() {
    let client = this.http.post(`/api/v1/nodes`, this.nodeEditForm.value)
    if (this.record != null) {
      client = this.http.put(`/api/v1/nodes/${this.record.id}`, this.nodeEditForm.value)
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
