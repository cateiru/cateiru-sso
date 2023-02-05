import {atom} from 'recoil';
import {OTPState, UserInfo} from './types';

// const localStorageEffect =
//   (key: string) =>
//   ({setSelf, onSet}: any) => {
//     if (process.browser) {
//       const savedValue = localStorage.getItem(key);
//       if (savedValue !== null) {
//         setSelf(JSON.parse(savedValue));
//       }

//       onSet((newValue: DefaultValue | string) => {
//         if (newValue instanceof DefaultValue) {
//           localStorage.removeItem(key);
//         } else {
//           localStorage.setItem(key, JSON.stringify(newValue));
//         }
//       });
//     }
//   };

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

export const LoadState = atom<boolean>({
  key: 'Load',
  default: false,
});

export const NoLoginState = atom<boolean>({
  key: 'noLogin',
  default: false,
});

export const OTPEnableState = atom<OTPState>({
  key: 'otpEnable',
  default: OTPState.Loading,
});
