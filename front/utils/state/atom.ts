import {atom} from 'recoil';

export const CTState = atom<string>({
  key: 'CT',
  default: '',
});

export const CreateNextState = atom<boolean>({
  key: 'CreateNext',
  default: false,
});
