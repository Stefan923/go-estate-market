import {CanActivateFn, Router} from '@angular/router';
import {inject} from "@angular/core";
import {AuthenticationService} from "../service/authentication.service";
import {map, tap} from "rxjs";

export const preAuthenticationGuard: CanActivateFn = (route, state) => {
  let authenticationService = inject(AuthenticationService)

  return authenticationService.isAuthenticated.pipe(
    map((authenticated: boolean) => {
      return !authenticated
    })
  )
}
