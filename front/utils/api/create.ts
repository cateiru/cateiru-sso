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
 * @param {string} password - user password
 * @param {string} recaptcha - reCAPTCHA token
 * @returns {string} - client_token
 */
export async function createTemp(
  mail: string,
  password: string,
  recaptcha: string
): Promise<string> {
  const api = new API();
  api.post(
    JSON.stringify({mail: mail, password: password, re_chaptcha: recaptcha})
  );

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
