import { Pipe, PipeTransform } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';

@Pipe({
  name: 'time',
})
export class TimePipe implements PipeTransform {
  transform(value: any, ...args: unknown[]): unknown {
    if (value < 1000) {
      return '小于1秒';
    }
    value = value / 1000;
    let h: any = Math.floor(value / 3600);
    h = h === 0 ? '' : h + '小时';
    let m: any = Math.floor(value / 60);
    m = m === 0 ? '' : m + '分';
    const s = Math.floor(value % 60);
    return h + m + s.toString() + '秒';
  }
}

@Pipe({
  name: 'name',
})
export class NamePipe implements PipeTransform {
  transform(value: any, ...args: unknown[]): unknown {
    switch (value) {
      case 'app_name':
        return '应用名称';
      case 'description':
        return '描述';
      case 'full_name':
        return '全名';
      case 'gitlab_url':
        return 'gitlab地址';
      case 'principal_name':
        return '责任人';
      case 'project_name':
        return '项目名';
      default:
        return value;
    }
  }
}

@Pipe({
  name: 'html',
})
export class HtmlPipe implements PipeTransform {
  constructor(private domSanitizer: DomSanitizer) {}

  transform(html: string) {
    return this.domSanitizer.bypassSecurityTrustHtml(html);
  }
}

@Pipe({
  name: 'status',
})
export class StatusPipe implements PipeTransform {
  transform(str: any, emun?: any) {
    for (let i = 0; i < emun?.length; i++) {
      const keylist = [];
      for (const key in emun[i]) {
        keylist.push(key);
      }
      if (str == emun[i][keylist[0]]) {
        return emun[i][keylist[1]];
      }
      if (str == emun[i][keylist[1]]) {
        return emun[i][keylist[0]];
      }
    }
  }
}

export const Pipes = [TimePipe, NamePipe, HtmlPipe, StatusPipe];
