import {API} from './api';

export interface LoginLogResponse {
  this_device: boolean;
  is_logout: boolean;
  last_login_date: string;
  date: string;
  access_id: string;
  ip_address: string;
  user_agent: string;
  is_sso: boolean;
  sso_publickey: string;
  user_id: string;
}

export const getLoginLog = async (
  limit: number | undefined
): Promise<LoginLogResponse[]> => {
  const api = new API();

  api.get();

  let path = '/user/history/login';

  if (typeof limit !== 'undefined') {
    path += `?limit=${limit}`;
  }

  const response = await api.connect(path);

  const body = (await response.json()) as LoginLogResponse[];

  return body.sort((a, b) => {
    // このデバイスの履歴は無条件にトップに持ってくる
    if (a.this_device) {
      return -1;
    } else if (b.this_device) {
      return 1;
    }

    let current: number;
    let target: number;

    if (a.is_logout) {
      current = Date.parse(a.date);
    } else {
      current = Date.parse(a.last_login_date);
    }

    if (b.is_logout) {
      target = Date.parse(b.date);
    } else {
      target = Date.parse(b.last_login_date);
    }

    return target - current;
  });
};
