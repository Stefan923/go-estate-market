import {HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from '@angular/common/http';
import {Injectable} from "@angular/core";
import {Observable} from "rxjs";

@Injectable({
  providedIn: "root"
})
export class ApiInterceptor implements HttpInterceptor {
  intercept(
    req: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    const apiReq = req.clone({ url: `http://localhost:8080/api/v1${req.url}` });
    return next.handle(apiReq);
  }
}
