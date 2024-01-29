import { TestBed } from '@angular/core/testing';
import {HttpInterceptor} from '@angular/common/http';

import { ApiInterceptor } from './api.interceptor';

describe('ApiInterceptor', () => {
  const interceptor: HttpInterceptor = new ApiInterceptor();

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(interceptor).toBeTruthy();
  });
});
