import {Routes} from '@angular/router';
import AuthenticationComponent from "./authentication/authentication.component";
import {HomeComponent} from "./home/home.component";
import {ListPropertiesComponent} from "./properties/list-properties/list-properties.component";
import {NewPropertyComponent} from "./properties/new-property/new-property.component";
import {authenticationGuard} from "./guard/authentication.guard";
import {preAuthenticationGuard} from "./guard/pre-authentication.guard";

export const routes: Routes = [
  {
    path: "",
    component: HomeComponent,
    canActivate: [authenticationGuard],
  },
  {
    path: "login",
    component: AuthenticationComponent,
    canActivate: [preAuthenticationGuard],
  },
  {
    path: "register",
    component: AuthenticationComponent,
    canActivate: [preAuthenticationGuard],
  },
  {
    path: "properties",
    component: ListPropertiesComponent,
    canActivate: [authenticationGuard],
    children: [
      {
        path: "new",
        component: NewPropertyComponent,
      },
    ],
  },
]
