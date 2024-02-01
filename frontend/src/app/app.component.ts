import {Component, inject} from '@angular/core';
import {RouterLink, RouterOutlet} from '@angular/router';
import {AuthenticationService} from "./service/authentication.service";
import {TokenService} from "./service/token.service";
import {HttpClientModule} from "@angular/common/http";
import {map} from "rxjs";
import {AsyncPipe} from "@angular/common";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    RouterLink,
    AsyncPipe,
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  protected readonly authenticationService: AuthenticationService
  constructor() {
    this.authenticationService = inject(AuthenticationService)
  }

  protected readonly map = map;
  protected readonly console = console;
}
