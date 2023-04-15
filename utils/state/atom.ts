import {atom} from 'recoil';
import {UserMe} from '../types/user';

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

export const UserState = atom<UserMe | null | undefined>({
  key: 'User',
  default: undefined,
});
