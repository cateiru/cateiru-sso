import {UserInfo} from '../state/types';
import {API} from './api';

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
