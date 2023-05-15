import {type AtomEffect, atom} from 'recoil';
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

const tabId = Math.random().toString(32).substring(2);
const broadcastEffect =
  <T>(key: string): AtomEffect<T> =>
  ({setSelf, onSet}) => {
    const bc = new BroadcastChannel(key);
    bc.addEventListener('message', event => {
      if (event.data.id !== tabId) {
        setSelf(event.data.value);
      }
    });

    onSet(newValue => {
      bc.postMessage({
        id: tabId,
        value: newValue,
      });
    });
  };

export const UserState = atom<UserMe | null | undefined>({
  key: 'User',
  default: undefined,
  effects: [broadcastEffect('user')],
});
