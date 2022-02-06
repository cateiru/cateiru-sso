import {UserInfo} from '../state/types';
import {API} from './api';

export const getAllUsers = async () => {
  const api = new API();

  api.get();

  const resp = await api.connect('/admin/user');

  return (await resp.json()) as UserInfo[];
};

export const getUsers = async (id: string) => {
  const api = new API();

  api.get();

  const resp = await api.connect(`/admin/user?id=${id}`);

  return (await resp.json()) as UserInfo[];
};

export const role = async (enable: boolean, role: string, userId: string) => {
  const api = new API();

  api.post(
    JSON.stringify({
      action: enable ? 'enable' : 'disable',
      role: role,
      user_id: userId,
    })
  );

  await api.connect('/admin/role');
};

export const deleteUser = async (userId: string) => {
  const api = new API();

  api.delete();

  await api.connect(`/admin/user?id=${userId}`);
};
