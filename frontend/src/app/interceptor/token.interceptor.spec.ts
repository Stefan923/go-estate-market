import { TestBed } from '@angular/core/testing';
import {HttpInterceptor} from '@angular/common/http';

import { TokenInterceptor } from './token.interceptor';
import {TokenService} from "../service/token.service";

describe('TokenInterceptor', () => {
  const interceptor: HttpInterceptor = new TokenInterceptor(new TokenService());

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(interceptor).toBeTruthy();
  });
});
