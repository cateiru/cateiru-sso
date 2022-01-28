import {UserInfo} from '../state/types';
import {API} from './api';

interface CreateResponse {
  client_token: string;
}

interface VerifyResponse {
  keep_this_page: boolean;
  client_token: string;
}

/**
 * @param {string} mail - user email
 * @param {string} recaptcha - reCAPTCHA token
 * @returns {string} - client_token
 */
export async function createTemp(
  mail: string,
  recaptcha: string
): Promise<string> {
  const api = new API();
  api.post(JSON.stringify({mail: mail, re_captcha: recaptcha}));

  const resp = (await (await api.connect('/create')).json()) as CreateResponse;

  return resp['client_token'];
}

/**
 * @param {string} mailToken - メールアドレスに送信されるトークン
 * @returns {VerifyResponse} - ClientTokenとこのページでやるかの選択
 */
export async function createVerify(mailToken: string): Promise<VerifyResponse> {
  const api = new API();
  api.post(JSON.stringify({mail_token: mailToken}));

  return (await (await api.connect('/create/verify')).json()) as VerifyResponse;
}

/**
 * @param {string} clientToken - client token
 * @param {string} firstName - 名前
 * @param {string} lastName - 名字
 * @param {string} userName - ユーザ名
 * @param {string} theme - テーマ
 * @param {string} password - パスワード
 * @returns {UserInfo} - ユーザ情報。 /meで取得したときと同じです
 */
export async function createInfo(
  clientToken: string,
  firstName: string,
  lastName: string,
  userName: string,
  theme: string,
  password: string
): Promise<UserInfo> {
  const api = new API();
  api.post(
    JSON.stringify({
      client_token: clientToken,
      first_name: firstName,
      last_name: lastName,
      user_name: userName,
      theme: theme,
      avatar_url: '', // これはいらなくね？
      password: password,
    })
  );

  return (await (await api.connect('/create/info')).json()) as UserInfo;
}
