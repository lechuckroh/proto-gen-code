import { ErrorType, UserStatus } from "./constants";

export const MsgErrorType: { [code: number]: string } = {
  [ErrorType.ERROR_NONE]: "ERROR_NONE",
  [ErrorType.ERROR_API_ENDPOINT_UNKNOWN]: "ERROR_API_ENDPOINT_UNKNOWN",
  [ErrorType.ERROR_INTERNAL_UNKNOWN]: "ERROR_INTERNAL_UNKNOWN",
  [ErrorType.ERROR_NOT_AUTHENTICATED]: "ERROR_NOT_AUTHENTICATED",
  [ErrorType.ERROR_NOT_AUTHORIZED]: "ERROR_NOT_AUTHORIZED",
  [ErrorType.ERROR_USER_NOT_FOUND]: "ERROR_USER_NOT_FOUND",
};
export const MsgUserStatus: { [code: number]: string } = {
  [UserStatus.USER_STATUS_OK]: "USER_STATUS_OK",
  [UserStatus.USER_STATUS_REMOVED]: "USER_STATUS_REMOVED",
};

