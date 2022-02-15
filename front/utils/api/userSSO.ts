import {API} from './api';

export interface ServiceLogInfo {
  client_id: string;
  name: string;
  service_icon: string;
  logs: ServiceLog[];
}

export interface ServiceLog {
  log_id: string;
  accept_date: string;
  client_id: string;
  user_id: string;
}

export const getUserSSO = async (): Promise<ServiceLogInfo[]> => {
  const api = new API();

  api.get();

  const resp = await api.connect('/user/oauth');

  return (await resp.json()) as ServiceLogInfo[];
};
