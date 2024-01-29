import {Routes} from '@angular/router';
import {AuthenticationService} from "./service/authentication.service";
import {inject} from "@angular/core";
import {map} from "rxjs";
import {AppComponent} from "./app.component";
import AuthenticationComponent from "./authentication/authentication.component";
import {HomeComponent} from "./home/home.component";
import {ListPropertiesComponent} from "./properties/list-properties/list-properties.component";
import {NewPropertyComponent} from "./properties/new-property/new-property.component";

export const routes: Routes = [
  {
    path: "",
      component: HomeComponent,
  },
  {
    path: "login",
    component: AuthenticationComponent,
    canActivate: [
      () => inject(AuthenticationService).isAuthenticated.pipe(map((isAuth) => !isAuth)),
    ],
  },
  {
    path: "register",
    component: AuthenticationComponent,
    canActivate: [
      () => inject(AuthenticationService).isAuthenticated.pipe(map((isAuth) => !isAuth)),
    ],
  },
  {
    path: "properties",
    component: ListPropertiesComponent,
    canActivate: [() => inject(AuthenticationService).isAuthenticated],
    children: [
      {
        path: "new",
        component: NewPropertyComponent,
        canActivate: [() => inject(AuthenticationService).isAuthenticated],
      },
    ],
  },
]
