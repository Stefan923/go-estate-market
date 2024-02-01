import { TestBed } from '@angular/core/testing';
import { CanActivateFn } from '@angular/router';

import { preAuthenticationGuard } from './pre-authentication.guard';

describe('preAuthenticationGuard', () => {
  const executeGuard: CanActivateFn = (...guardParameters) => 
      TestBed.runInInjectionContext(() => preAuthenticationGuard(...guardParameters));

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});
