import { Injectable } from '@angular/core';
import {TokenDetail} from "../model/token-detail";

@Injectable({
  providedIn: 'root'
})
export class TokenService {
  private static readonly ACCESS_TOKEN_KEY: string = "access_token"
  private static readonly REFRESH_TOKEN_KEY: string = "access_token"
  private static readonly ACCESS_TOKEN_EXPIRE_TIME_KEY: string = "access_token_expire_time"
  private static readonly REFRESH_TOKEN_EXPIRE_TIME_KEY: string = "refresh_token_expire_time"

  constructor() {}

  getAccessToken(): string {
    return window.localStorage[TokenService.ACCESS_TOKEN_KEY]
  }

  getRefreshToken(): string {
    return window.localStorage[TokenService.REFRESH_TOKEN_KEY]
  }

  getAccessTokenExpireTime(): string {
    return window.localStorage[TokenService.ACCESS_TOKEN_EXPIRE_TIME_KEY]
  }

  getRefreshTokenExpireTime(): string {
    return window.localStorage[TokenService.REFRESH_TOKEN_EXPIRE_TIME_KEY]
  }

  saveTokenDetail(tokenDetail: TokenDetail): void {
    window.localStorage[TokenService.ACCESS_TOKEN_KEY] = tokenDetail.accessToken
    window.localStorage[TokenService.REFRESH_TOKEN_KEY] = tokenDetail.refreshToken
    window.localStorage[TokenService.ACCESS_TOKEN_EXPIRE_TIME_KEY] = tokenDetail.accessTokenExpireTime
    window.localStorage[TokenService.REFRESH_TOKEN_EXPIRE_TIME_KEY] = tokenDetail.refreshTokenExpireTime
  }

  destroyTokenDetail(): void {
    window.localStorage.removeItem(TokenService.ACCESS_TOKEN_KEY)
    window.localStorage.removeItem(TokenService.REFRESH_TOKEN_KEY)
    window.localStorage.removeItem(TokenService.ACCESS_TOKEN_EXPIRE_TIME_KEY)
    window.localStorage.removeItem(TokenService.REFRESH_TOKEN_EXPIRE_TIME_KEY)
  }
}
