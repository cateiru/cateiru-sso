import {UserInfo} from '../state/types';
import {API} from './api';

interface ChangeMailVerifyResponse {
  new_mail: string;
}

export const changeUser = async (
  firstName: string | undefined,
  lastName: string | undefined,
  userName: string | undefined,
  theme: string | undefined
): Promise<UserInfo> => {
  const api = new API();

  api.post(
    JSON.stringify({
      first_name: firstName,
      last_name: lastName,
      user_name: userName,
      theme: theme,
    })
  );

  const response = await api.connect('/user/info');

  return (await response.json()) as UserInfo;
};

export const changeMail = async (newMail: string) => {
  const api = new API();

  api.post(JSON.stringify({type: 'change', new_mail: newMail}));

  await api.connect('/user/mail');
};

export const changeMailVerify = async (token: string): Promise<string> => {
  const api = new API();

  api.post(JSON.stringify({type: 'verify', mail_token: token}));

  const resp = await api.connect('/user/mail');

  return ((await resp.json()) as ChangeMailVerifyResponse).new_mail;
};

export const changePassword = async (oldPW: string, newPW: string) => {
  const api = new API();

  api.post(JSON.stringify({new_password: newPW, old_password: oldPW}));

  await api.connect('/user/password');
};
