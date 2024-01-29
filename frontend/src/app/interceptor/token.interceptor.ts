import {HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from '@angular/common/http';
import {Injectable} from "@angular/core";
import {Observable} from "rxjs";
import {TokenService} from "../service/token.service";

@Injectable({ providedIn: "root" })
export class TokenInterceptor implements HttpInterceptor {
  constructor(
    private readonly tokenService: TokenService
  ) {}

  intercept(
    req: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    const token = this.tokenService.getAccessToken()

    const request = req.clone({
      setHeaders: {
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
      },
    })
    return next.handle(request)
  }
}
