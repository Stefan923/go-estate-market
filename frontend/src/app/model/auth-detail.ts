import {UserDetail} from "./user-detail";
import {TokenDetail} from "./token-detail";

export interface AuthDetail {
  userDetail: UserDetail,
  tokenDetail: TokenDetail,
}
