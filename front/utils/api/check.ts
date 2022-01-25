import {API} from './api';

interface CheckUserNameResponse {
  exist: boolean;
}

/**
 * @param {string} userName - ユーザ名
 * @returns {boolean} - 存在するかどうか
 */
export async function checkUserName(userName: string): Promise<boolean> {
  const api = new API();

  api.get();

  const resp = (await (
    await api.connect(`/check/username?name=${userName}`)
  ).json()) as CheckUserNameResponse;

  return resp.exist;
}
