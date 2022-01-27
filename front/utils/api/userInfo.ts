import {UserInfo} from '../state/types';
import {API} from './api';
/**
 * ユーザ情報を取得する
 *
 * @returns {UserInfo} - ユーザのデータ
 */
export default async function getUserInfo(): Promise<UserInfo> {
  const api = new API();
  api.get();

  return (await (await api.connect('/me')).json()) as UserInfo;
}
