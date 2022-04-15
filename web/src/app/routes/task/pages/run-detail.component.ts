import {Component, OnDestroy, OnInit} from '@angular/core';
import {_HttpClient, TitleService} from '@delon/theme';
import {NzMessageService} from 'ng-zorro-antd';
import {webSocket} from 'rxjs/webSocket';
import {PlatformLocation} from "@angular/common";
import {Subject} from "rxjs";

@Component({
  selector: 'run-detail',
  templateUrl: './run-detail.component.html',
})
export class RunDetailComponent implements OnInit, OnDestroy {
  private id: string; // Task ID
  private run_id: string; // Run ID

  constructor(
    private http: _HttpClient,
    public location: PlatformLocation,
    public message: NzMessageService,
    public title: TitleService,
  ) {
  }

  ngOnInit() {
    this.loadData()
  }

  ngOnDestroy() {
  }

  getWebSocketUrlBase() {
    let url = this.location.href;
    if (url.indexOf("https://") == 0) {
      let s = url.substr(8, url.length)
      return `wss://${s.split('/')[0]}`
    }
    if (url.indexOf("http://") == 0) {
      let s = url.substr(7, url.length)
      return `ws://${s.split('/')[0]}`
    }
    return ''
  }

  logContent = ''

  running = true;

  loadData() {
    this.http.get(`/api/v1/tasks/${this.id}/runs/${this.run_id}`).subscribe((res) => {
      console.log(res)
    })

    let wsBase = this.getWebSocketUrlBase()
    if (wsBase == '') {
      this.message.error('access denied')
      return
    }

    const closeEvent = new Subject();
    closeEvent.subscribe(() => {
      console.log('closed')
      this.running = false
    })

    this.running = true;
    webSocket({
      url: `${wsBase}/api/v1/tasks/${this.id}/runs/${this.run_id}/log`,
      deserializer: (msg) => msg,
      closeObserver: closeEvent
    }).subscribe((msg) => {
      this.logContent = this.logContent + msg.data
    })
  }
}
