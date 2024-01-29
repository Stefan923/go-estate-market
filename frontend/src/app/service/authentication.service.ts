import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {BehaviorSubject, distinctUntilChanged, map, Observable, tap} from "rxjs";
import {User} from "../model/user";
import {TokenDetail} from "../model/token-detail";
import {TokenService} from "./token.service";
import {RegisterDetail} from "../model/register-detail";
import {BaseResponse} from "../model/base-response";

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private static readonly LOGIN_PATH: string = "/auth/login"
  private static readonly REGISTER_PATH: string = "/auth/register"

  private currentUserSubject = new BehaviorSubject<User | null>(null);

  public currentUser = this.currentUserSubject.asObservable().pipe(distinctUntilChanged())
  public isAuthenticated = this.currentUser.pipe(map(user => !!user))

  constructor(
    private readonly httpClient: HttpClient,
    private readonly tokenService: TokenService,
  ) { }

  login(credentials: {
    email: string;
    password: string;
  }): Observable<BaseResponse<TokenDetail>> {
    return this.httpClient
      .post<BaseResponse<TokenDetail>>(AuthenticationService.LOGIN_PATH, credentials)
      .pipe(tap((response: BaseResponse<TokenDetail>) => this.onAuthenticationResponse(response)))
  }

  register(registerDetail: RegisterDetail): Observable<BaseResponse<TokenDetail>> {
    return this.httpClient
      .post<BaseResponse<TokenDetail>>(AuthenticationService.REGISTER_PATH, registerDetail)
      .pipe(tap((response: BaseResponse<TokenDetail>) => this.onAuthenticationResponse(response)))
  }

  logout(): void {
    this.tokenService.destroyTokenDetail();
    this.currentUserSubject.next(null);
  }

  onAuthenticationResponse(response: BaseResponse<TokenDetail>): void {
    this.tokenService.saveTokenDetail(response.result)
  }
}
