import { PlatformLocation } from '@angular/common';
import { ChangeDetectionStrategy, Component, Inject, ChangeDetectorRef } from '@angular/core';
import { Router } from '@angular/router';
import { DA_SERVICE_TOKEN, ITokenService } from '@delon/auth';
import { SettingsService, _HttpClient } from '@delon/theme';

@Component({
  selector: 'header-user',
  styles: [
    `
      ::ng-deep .alain-default__nav-item {
        color: #000000a6 !important;
      }
    `,
  ],
  template: `
    <div class="alain-default__nav-item d-flex align-items-center px-sm" nz-dropdown nzPlacement="bottomRight" [nzDropdownMenu]="userMenu">
      {{ settings?.user.username }}
    </div>
    <nz-dropdown-menu #userMenu="nzDropdownMenu">
      <div nz-menu class="width-sm">
        <div nz-menu-item (click)="gotoPreference()">
          <i nz-icon nzType="setting" class="mr-sm"></i> Preference
        </div>
        <div nz-menu-item (click)="logout()">
          <i nz-icon nzType="logout" class="mr-sm"></i> Sign out
        </div>
      </div>
    </nz-dropdown-menu>
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HeaderUserComponent {
  constructor(
    public settings: SettingsService,
    private router: Router,
    public http: _HttpClient,
    @Inject(DA_SERVICE_TOKEN) private tokenService: ITokenService,
    private location: PlatformLocation,
    public cdr: ChangeDetectorRef,
  ) {}

  ngOnInit() {
    setTimeout(() => {
      this.cdr.markForCheck();
      this.cdr.detectChanges();
    }, 500);
  }

  logout() {
    sessionStorage.clear();
    localStorage.clear();
    window.location.reload();
  }

  gotoPreference() {
    this.router.navigate([`/setting`]).then()
  }
}
