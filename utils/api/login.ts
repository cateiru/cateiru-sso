import {API} from './api';

interface LoginResponse {
  is_otp: boolean;
  otp_token: string;
}

/**
 * ログインする
 *
 * @param {string} mail - メールアドレス
 * @param {string} password - パスワード
 * @param {string} recaptcha - reCAPTCHAのトークン
 * @returns {string | undefined} - OTPが必要な場合専用のトークンを返す
 */
export async function login(
  mail: string,
  password: string,
  recaptcha: string
): Promise<string | undefined> {
  const api = new API();

  api.post(
    JSON.stringify({mail: mail, password: password, re_captcha: recaptcha})
  );

  const resp = (await (await api.connect('/login')).json()) as LoginResponse;

  if (resp) {
    return resp.otp_token;
  }

  return undefined;
}

/**
 * OTPを入力してログインする
 *
 * @param {string} passcode - OTPパスコード
 * @param {string} otpToken - OTP トークン
 */
export async function loginOTP(
  passcode: string,
  otpToken: string
): Promise<void> {
  const api = new API();

  api.post(JSON.stringify({passcode: passcode, otp_token: otpToken}));

  await api.connect('/login/onetime');
}
