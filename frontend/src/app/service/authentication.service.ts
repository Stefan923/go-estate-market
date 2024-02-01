import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {BehaviorSubject, distinctUntilChanged, map, Observable, tap} from "rxjs";
import {UserDetail} from "../model/user-detail";
import {TokenService} from "./token.service";
import {RegisterDetail} from "../model/register-detail";
import {BaseResponse} from "../model/base-response";
import {AuthDetail} from "../model/auth-detail";
import {Router} from "@angular/router";

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private static readonly USER_DETAIL_KEY: string = "user_detail"

  private static readonly LOGIN_PATH: string = "/auth/login"
  private static readonly REGISTER_PATH: string = "/auth/register"

  private currentUserSubject = new BehaviorSubject<UserDetail | null>(null);

  public currentUser = this.currentUserSubject.asObservable().pipe(distinctUntilChanged())
  public isAuthenticated = this.currentUser.pipe(map(user => !!user))

  constructor(
    private readonly httpClient: HttpClient,
    private readonly tokenService: TokenService,
    private readonly router: Router
  ) {
    let userDetail = this.getUserDetail()

    if (userDetail !== null) {
      this.currentUserSubject.next(this.getUserDetail())
    }
  }

  login(credentials: { email: string; password: string; }): Observable<BaseResponse<AuthDetail>> {
    return this.httpClient
      .post<BaseResponse<AuthDetail>>(AuthenticationService.LOGIN_PATH, credentials)
      .pipe(tap((response: BaseResponse<AuthDetail>) => this.onAuthenticationResponse(response)))
  }

  register(registerDetail: RegisterDetail): Observable<BaseResponse<AuthDetail>> {
    return this.httpClient
      .post<BaseResponse<AuthDetail>>(AuthenticationService.REGISTER_PATH, registerDetail)
      .pipe(tap((response: BaseResponse<AuthDetail>) => this.onAuthenticationResponse(response)))
  }

  logout(): void {
    this.destroyUserDetail()
    this.tokenService.destroyTokenDetail();
    this.currentUserSubject.next(null);

    this.router.navigate(['/login'])
  }

  onAuthenticationResponse(response: BaseResponse<AuthDetail>): void {
    if (response.success) {
      this.tokenService.saveTokenDetail(response.result.tokenDetail)
      this.currentUserSubject.next(response.result.userDetail)
      this.saveUserDetail(response.result.userDetail)
    }
  }

  private getUserDetail(): UserDetail {
    return window.localStorage[AuthenticationService.USER_DETAIL_KEY]
  }

  private saveUserDetail(userDetail: UserDetail): void {
    window.localStorage[AuthenticationService.USER_DETAIL_KEY] = userDetail
  }

  private destroyUserDetail(): void {
    window.localStorage.removeItem(AuthenticationService.USER_DETAIL_KEY)
  }
}
