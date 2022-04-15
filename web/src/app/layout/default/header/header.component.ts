import {ChangeDetectionStrategy, ChangeDetectorRef, Component} from '@angular/core';
import {SettingsService, _HttpClient} from '@delon/theme';
import {Router} from '@angular/router';
import {AppData} from "../../../core/app-data";

@Component({
  selector: 'layout-header',
  templateUrl: './header.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
  styles: [
    `.ant-menu-horizontal {
        line-height: 66px;
    }

    .alain-default__nav li:active {
        border: none;
    }

    .menu_doc {
        padding: 0 40px;
        cursor: pointer;
    }

    li.menu_doc:hover {
        color: #1890ff;
        display: block;
        border-bottom: 2px solid #1890ff
    }

    .menu_doc .anticon {
        min-width: 14px;
        margin-right: 10px;
        font-size: 14px;
    }
    `,
  ],
})
export class HeaderComponent {
  searchToggleStatus: boolean;
  menu = [];

  constructor(public cdr: ChangeDetectorRef, public router: Router, public settings: SettingsService) {
    this.settings.setLayout('collapsed', false);
  }

  ngOnInit() {
    this.menu = AppData.menu[0].children
    this.cdr.detectChanges();
  }

  go(link: string, index?: number) {
    if (index != undefined) {
      this.router.navigate([link || this.menu[index].children[0].link]).then(() => {
      });
    } else {
      this.router.navigate([link]).then(() => {
      });
    }
  }
}
