import {API} from './api';

export const sendForget = async (mail: string) => {
  const api = new API();

  api.post(JSON.stringify({mail: mail}));

  await api.connect('/password/forget');
};

export const acceptPassword = async (token: string, newPW: string) => {
  const api = new API();

  api.post(JSON.stringify({forget_token: token, new_password: newPW}));

  await api.connect('/password/forget/accept');
};
