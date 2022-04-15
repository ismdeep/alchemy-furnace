import { ExceptionModule } from '@delon/abc/exception';
import { GlobalFooterModule } from '@delon/abc/global-footer';
import { NoticeIconModule } from '@delon/abc/notice-icon';
import { PageHeaderModule } from '@delon/abc/page-header';
import { ResultModule } from '@delon/abc/result';
// import { ReuseTabModule } from '@delon/abc/reuse-tab';
import { SEModule } from '@delon/abc/se';
import { SidebarNavModule } from '@delon/abc/sidebar-nav';
import { STModule } from '@delon/abc/st';
import { SVModule } from '@delon/abc/sv';
import { G2BarModule } from '@delon/chart/bar';
import { G2CardModule } from '@delon/chart/card';
import { G2MiniBarModule } from '@delon/chart/mini-bar';
import { G2MiniProgressModule } from '@delon/chart/mini-progress';
import { G2TimelineModule } from '@delon/chart/timeline';
import { TrendModule } from '@delon/chart/trend';

export const SHARED_DELON_MODULES = [
  PageHeaderModule,
  ResultModule,
  ExceptionModule,
  NoticeIconModule,
  SidebarNavModule,
  GlobalFooterModule,
  STModule,
  SEModule,
  SVModule,
  G2BarModule,
  G2MiniBarModule,
  G2CardModule,
  G2MiniProgressModule,
  TrendModule,
  G2TimelineModule,
  // ReuseTabModule,
];
