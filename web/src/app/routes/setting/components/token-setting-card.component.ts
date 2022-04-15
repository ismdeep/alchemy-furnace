import {Component, OnDestroy, OnInit} from "@angular/core";
import {_HttpClient, DrawerHelper, ModalHelper} from "@delon/theme";
import {TokenEditFormComponent} from "./token-edit-form.component";
import format from 'date-fns/format';

@Component({
  selector: 'token-setting-card',
  templateUrl: './token-setting-card.component.html',
})
export class TokenSettingCardComponent implements OnInit, OnDestroy {
  constructor(
    private http: _HttpClient,
    public modalHelper: ModalHelper,
    private drawerHelper: DrawerHelper,
  ) {
  }

  ngOnDestroy(): void {
  }

  ngOnInit(): void {
    this.loadData()
  }

  // data
  tokens = []
  loadData() {
    this.http.get(`/api/v1/tokens`).subscribe((res) => {
      if (res.data) {
        this.tokens = res.data
      } else {
        this.tokens = []
      }
    })
  }

  formatTime(t) {
    return format(new Date(t), 'yyyy-MM-dd  HH:mm:ss');
  }

  add() {
    this.drawerHelper.create('Generate new token', TokenEditFormComponent, {}, {size: document.body.clientWidth * 0.618}).subscribe(() => {
      this.loadData()
    });
  }

  edit(token) {
    this.drawerHelper.create('Edit token', TokenEditFormComponent, {record: token}, {size: document.body.clientWidth * 0.618}).subscribe(() => {
      this.loadData()
    })
  }

  delete(node) {
    this.http.delete(`/api/v1/tokens/${node.id}`).subscribe(() => {
      this.loadData()
    })
  }
}