import {UserInfo} from '../state/types';
import {API} from './api';
import decodeErrorCode from './errorCode';

/**
 * ユーザ情報を取得する
 *
 * @returns {UserInfo} - ユーザのデータ
 */
export default async function getUserInfo(): Promise<UserInfo | null> {
  const api = new API();
  api.get();

  const response = await api.connectNoErr('/me');

  // 403が返るときはログインできないことなのでnullを返す
  if (response.status === 403) {
    return null;
  }

  if (!response.ok) {
    const resp = await response.json();
    if (resp['code'] !== 1) {
      throw new Error(decodeErrorCode(resp['code']));
    }
    throw new Error(`${resp['status_code']}::ID:${resp['error_id']}`);
  }

  return (await response.json()) as UserInfo;
}
