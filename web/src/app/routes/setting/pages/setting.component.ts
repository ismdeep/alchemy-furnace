import {Component, OnDestroy} from "@angular/core";
import {_HttpClient, ModalHelper, TitleService} from "@delon/theme";

@Component({
  selector: 'setting',
  templateUrl: 'setting.component.html',
})
export class SettingComponent implements OnDestroy {

  constructor(
    private http: _HttpClient,
    private modalHelper: ModalHelper,
    public title: TitleService,
  ) {
  }

  ngOnInit() {
    this.title.setTitle("Setting - Alchemy Furnace")
  }

  ngOnDestroy(): void {
  }
}
