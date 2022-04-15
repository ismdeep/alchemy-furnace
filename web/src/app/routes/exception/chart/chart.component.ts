import {AfterViewInit, ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {_HttpClient} from '@delon/theme';
import Axios from 'axios';
import {webSocket} from 'rxjs/webSocket';

setTimeout(() => {
  document.title = '服务状态监控面板-UOS';
}, 100);

@Component({
  selector: 'app-monitor-charts',
  templateUrl: './chart.component.html',
  styles: [
    `
      ::ng-deep .alain-fullscreen {
        padding: 0;
      }

      .background {
        background-color: #f5f7fa;
        width: 100%;
      }

      i {
        font-size: 24px;
      }

      ::ng-deep .g2-card__top {
        display: flex;
        justify-content: center;
      }

      ::ng-deep .g2-card__meta {
        display: flex;
        justify-content: center;
      }

      ::ng-deep .trend {
        display: flex;
        justify-content: center;
      }

      ::ng-deep .g2-card__avatar {
        margin-right: 0px;
      }

      ::ng-deep .pt-lg {
        padding-top: 0px !important;
      }

      {
      }
    `,
  ],
})
export class ChartComponent implements OnInit {
  constructor(public http: _HttpClient,public cdr: ChangeDetectorRef, private fb: FormBuilder) {
    this.validateForm = this.fb.group({
      app: [''],
      service: [''],
      health_status: [''],
      ts: ['5'],
    });
  }

  validateForm: FormGroup;
  serviceName; // 服务名称
  status;
  data: any;
  statusData: any[] = [];
  offlineChartData;
  webSite;
  appName;
  cost;
  count;
  countTotal;
  countMaintain;
  countErr;
  totaldata: any[] = [];
  timeDate: any[] = [];
  theme = {
    height: 100,
  };

  getInit;

  socket(url, options?) {
    const opt = [];
    Object.keys(options).forEach((e) => {
      opt.push(`${e}=${options[e]}`);
    });
    return webSocket('ws://10.20.22.172:9811' + url + '?' + opt.join('&'));
  }

  ngOnInit() {
    this.getInitData().then((e) => {
      this.serviceName = this.appName.filter((v) => v.value == '')[0].sub_service;
      this.validateForm.patchValue({service: ''});
    });
  }

  getInitData() {
    return new Promise((resolve) => {
      this.http.get('/api/v1/service/health/app/list').subscribe((res) => {
        this.appName = res.data.data;
        // 添加全部
        this.appName = this.appName.map((e) => {
          return {name: e.name, value: e.value, sub_service: ['', ...e.sub_service]};
        });
        this.appName.unshift({name: '全部', value: '', sub_service: ['']});
        resolve(this.appName);
        this.getData();
        this.getCode();
        this.getTotal();
        this.reqTime();
      });
    });
  }

  getData() {
    // this.socket('/api/v1/service/health/request/cost', ...this.validateForm.value).subscribe((res) => {});
    // 平均响应时间

    this.http.get('/api/v1/service/health/request/cost', {params: this.validateForm.value}).subscribe((res) => {
      this.cost = res.data.data;
    });
    // 正常数
    this.validateForm.value.health_status = 1;
    this.http.get('/api/v1/service/health/request/count', {params: this.validateForm.value}).subscribe((res) => {
      this.count = res.data.data;
    });

    // 维护数
    this.validateForm.value.health_status = 2;
    this.http.get('/api/v1/service/health/request/count', {params: this.validateForm.value}).subscribe((res) => {
      this.countMaintain = res.data.data;
    });
    // 错误数
    this.validateForm.value.health_status = -1;
    this.http.get('/api/v1/service/health/request/count', {params: this.validateForm.value}).subscribe((res) => {
      this.countErr = res.data.data;
    });

    // 请求总数
    delete this.validateForm.value.health_status;
    this.http.get('/api/v1/service/health/request/count', {params: this.validateForm.value}).subscribe((res) => {
      this.countTotal = res.data.data;
    });
  }

  getCode() {
    this.http.get('/api/v1/service/health/request/status_code/statistics', {params: this.validateForm.value}).subscribe((res) => {
      const statusData: any[] = [];
      res.data.data.sort((a, b) => {
        return a.Status - b.Status;
      });
      res.data.data.forEach((e) => {
        statusData.push({
          x: `状态码${e.Status}`,
          y: e.Count,
        });
      });
      this.statusData = statusData;
    });
  }

  getTotal() {
    this.http.get('/api/v1/service/health/request/count/per_minute', {params: this.validateForm.value}).subscribe((res) => {
      const totaldata = [];
      const data = res.data.data;
      data.pop();
      data.forEach((e) => {
        totaldata.push({
          time: e.Time,
          y1: e.Count,
          y2: null,
        });
      });
      this.totaldata = totaldata;
    });
  }

  reqTime() {
    this.http.get('/api/v1/service/health/request/avg_cost/per_minute', {params: this.validateForm.value}).subscribe((res) => {
      const reqtime = [];
      res.data.data.forEach((e) => {
        reqtime.push({
          time: e?.Time,
          // tslint:disable-next-line: no-bitwise
          y1: ~~(~~e?.Cost as number).toFixed(2),
          y2: null,
        });
      });
      this.timeDate = reqtime;
    });
  }

  changeTime() {
    this.getInitData();
  }

  getService() {
    this.getInitData();
  }

  getAppName(e) {
    this.serviceName = this.appName.filter((v) => v.value == e)[0].sub_service;
    this.getInitData().then((e) => {
      this.validateForm.patchValue({service: ''});
    });
  }

  reast() {
    this.validateForm.patchValue({app: ''});
    this.validateForm.patchValue({service: ''});
  }
}
