import {Component, OnInit, ViewChild, Inject} from '@angular/core';
import {NzMessageService} from 'ng-zorro-antd';
import {SFComponent, SFSchema} from '@delon/form';
import {_HttpClient} from '@delon/theme';
import {Router} from '@angular/router';
import {DA_SERVICE_TOKEN, ITokenService} from '@delon/auth';

@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.less'],
})
export class IndexComponent implements OnInit {
  @ViewChild('sf') sf: SFComponent;

  constructor(
    private message: NzMessageService,
    private http: _HttpClient,
    private router: Router,
    @Inject(DA_SERVICE_TOKEN) private tokenService: ITokenService,
  ) {
  }

  schema: SFSchema = {
    properties: {
      username: {
        type: 'string',
        title: 'User',
        ui: {
          placeholder: 'Please input username',
          errors: {
            required: 'Please input username',
          },
        },
      },
      password: {
        type: 'string',
        title: 'Pass',
        ui: {
          type: 'password',
          placeholder: 'Please input password',
          errors: {
            required: 'Please input password',
          },
        },
      },
    },
    required: ['username', 'password'],
  };

  ngOnInit(): void {
  }

  isLogging = false
  login_text = 'Login'

  submit(e) {
    this.isLogging = true;
    this.http.post(`/api/v1/sign-in`, e).subscribe((res) => {
      setTimeout(() => {
        this.isLogging = false
        this.login_text = 'Login success'
        this.message.success('Login successfully');
        this.tokenService.set({token: res.data});
        window.location.href = '/'
      }, 1000)
    }, () => {
      setTimeout(() => {
        this.isLogging = false
      }, 1000)
    });
  }
}
