import {API} from './api';

interface CheckUserNameResponse {
  exist: boolean;
}

interface OTPMeResponse {
  enable: boolean;
}

/**
 * @param {string} userName - ユーザ名
 * @returns {boolean} - 存在するかどうか
 */
export async function checkUserName(userName: string): Promise<boolean> {
  const api = new API();

  api.get();

  const resp = (await (
    await api.connect(`/check/username?name=${userName.toLowerCase()}`)
  ).json()) as CheckUserNameResponse;

  return resp.exist;
}

/**
 * OTPが設定されているか確認する
 *
 * @returns {boolean} - OTPが設定されているか
 */
export async function isEnableOTP(): Promise<boolean> {
  const api = new API();

  api.get();

  const resp = (await (
    await api.connect('/user/otp/me')
  ).json()) as OTPMeResponse;

  return resp.enable;
}
