import {UserInfo} from '../state/types';
import {API} from './api';

export const getAllUsers = async () => {
  const api = new API();

  api.get();

  const resp = await api.connect('/admin/user');

  return (await resp.json()) as UserInfo[];
};
