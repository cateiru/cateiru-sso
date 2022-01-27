import {API} from './api';

/**
 * logout
 */
export async function logout() {
  const api = new API();
  api.get();

  await api.connect('/logout');
}

/**
 * delete account
 */
export async function deleteAccount() {
  const api = new API();
  api.delete();

  await api.connect('/logout');
}
