import { ChangeDetectorRef, Component, forwardRef, Input, OnDestroy } from '@angular/core';
import { NG_VALUE_ACCESSOR } from '@angular/forms';
import { _HttpClient } from '@delon/theme';
import { Subject } from 'rxjs';
import { debounceTime, distinctUntilChanged, switchMap } from 'rxjs/operators';

@Component({
  selector: 'app-search',
  templateUrl: 'search.component.html',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => AppSearchComponent),
      multi: true,
    },
  ],
})
export class AppSearchComponent implements OnDestroy {
  options: string[] = [];
  @Input() placeholder: string;
  @Input() type;
  @Input() maxLength;
  @Input() name;

  search$ = new Subject<string>();
  constructor(http: _HttpClient, cdr: ChangeDetectorRef) {
    this.search$
      .pipe(
        debounceTime(500),
        distinctUntilChanged(),
        // tslint:disable-next-line: no-bitwise
        switchMap((q) => http.post('/api/v1/portainer/service/getports', { port: ~~q },
        null)),
      )
      .subscribe((res: any) => {
        this.options = res.data;
        cdr.detectChanges();
      });
  }

  value = 0;

  ngOnDestroy(): void {
    this.search$.unsubscribe();
  }

  onChange: (_: any) => void = (_: any) => {};

  onTouched: () => void = () => {};

  updateChanges() {
    this.onChange(this.value);
  }

  writeValue(value: number): void {
    this.value = value;
    this.updateChanges();
  }

  registerOnChange(fn: any): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }
}
