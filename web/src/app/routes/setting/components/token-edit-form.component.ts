import {Component, OnDestroy, OnInit} from '@angular/core';
import {_HttpClient} from '@delon/theme';
import {NzMessageService} from 'ng-zorro-antd/message';
import {NzDrawerRef} from 'ng-zorro-antd';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'token-edit-form',
  templateUrl: './token-edit-form.component.html',
})
export class TokenEditFormComponent implements OnInit, OnDestroy {

  // params
  record = null;

  // form
  tokenEditForm: FormGroup;
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
    this.tokenEditForm = this.fb.group({
      name: [null, [Validators.required]],
    });
    if (this.record) {
      console.log(this.record)
      this.tokenEditForm.patchValue(this.record)
    }
  }

  ngOnDestroy() {
    this.ref.close(true)
  }

  tokenKey = ""
  submitForm() {
    if (this.record != null) {
      this.http.put(`/api/v1/tokens/${this.record.id}`, this.tokenEditForm.value).subscribe(() => {
        this.msgSrv.success('Token name changed successfully.')
      })
    } else {
      this.http.post(`/api/v1/tokens`, this.tokenEditForm.value).subscribe((res) => {
        this.tokenKey = res.data.token_key
        this.msgSrv.success('Token generated successfully.')
      })
    }
  }

  close() {
    this.ref.close(true)
  }
}
