import {CanActivateFn, Router} from '@angular/router';
import {inject} from "@angular/core";
import {AuthenticationService} from "../service/authentication.service";
import {map, tap} from "rxjs";

export const authenticationGuard: CanActivateFn = (route, state) => {
  let authenticationService = inject(AuthenticationService)
  let router = inject(Router)

  return authenticationService.isAuthenticated.pipe(
    map((authenticated: boolean) => {
      if (authenticated) {
        return true
      } else {
        router.navigate(['/login'])
        return false
      }
    })
  )
}
