import {Component, OnInit} from '@angular/core';
import {NgOptimizedImage} from "@angular/common";
import {ActivatedRoute, Router} from "@angular/router";
import LoginComponent from "./login/login.component";
import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import RegisterComponent from "./register/register.component";

@Component({
  selector: 'app-authentication',
  standalone: true,
  imports: [
    NgOptimizedImage,
    LoginComponent,
    RegisterComponent,
  ],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  templateUrl: './authentication.component.html',
  styleUrl: './authentication.component.css'
})
export default class AuthenticationComponent implements OnInit {
  protected readonly AuthenticationType = AuthenticationType;
  protected authenticationType: AuthenticationType = AuthenticationType.LOGIN

  constructor(
    private readonly route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    this.authenticationType = this.route.snapshot.url.at(-1)!.path === "login" ?
      AuthenticationType.LOGIN : AuthenticationType.REGISTER
  }
}

enum AuthenticationType {
  LOGIN,
  REGISTER
}
