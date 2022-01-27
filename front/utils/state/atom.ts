import {atom} from 'recoil';
import {UserInfo} from './types';

export const CTState = atom<string>({
  key: 'CT',
  default: '',
});

export const CreateNextState = atom<boolean>({
  key: 'CreateNext',
  default: false,
});

export const UserState = atom<UserInfo | null | undefined>({
  key: 'User',
  default: undefined,
});
