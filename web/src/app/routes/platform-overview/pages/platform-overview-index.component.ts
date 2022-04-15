import {Component, OnInit} from '@angular/core';
import {_HttpClient, ModalHelper, TitleService} from '@delon/theme';
import {NzMessageService} from 'ng-zorro-antd';

@Component({
  selector: 'platform-overview',
  templateUrl: './platform-overview-index.component.html',
  styleUrls: ['./platform-overview-index.component.css'],
})
export class PlatformOverviewIndexComponent implements OnInit {

  constructor(
    private http: _HttpClient,
    private modal: ModalHelper,
    private message: NzMessageService,
    public title: TitleService,
  ) {

  }

  ngOnInit() {
    this.title.setTitle('Welcome - Alchemy Furnace')
  }
}
