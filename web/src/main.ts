import {enableProdMode, ViewEncapsulation} from '@angular/core';
import {platformBrowserDynamic} from '@angular/platform-browser-dynamic';
import {AppModule} from './app/app.module';
import {preloaderFinished} from '@delon/theme';

preloaderFinished();
enableProdMode();

const bootstrap = () => {
  return platformBrowserDynamic().bootstrapModule(AppModule, {
    defaultEncapsulation: ViewEncapsulation.Emulated,
  }).then((res) => {
    if ((window as any).appBootstrap) {
      (window as any).appBootstrap()
    }
    return res;
  });
};

bootstrap().then()
