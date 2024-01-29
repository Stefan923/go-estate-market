import {ValidationError} from "./validation-error";

export interface BaseResponse<T> {
  result: T,
  success: boolean,
  validationErrors: ValidationError[],
  error: string,
}
