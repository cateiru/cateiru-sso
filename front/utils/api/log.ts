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

  return body.sort((a, b) => Date.parse(b.date) - Date.parse(a.date));
};
