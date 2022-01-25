import {API} from './api';

export interface UserInfo {
  first_name: string;
  last_name: string;
  user_name: string;
  mail: string;
  theme: string;
  avatar_url: string;
  user_id: string;
}

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
