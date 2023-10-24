import {atom} from 'jotai';
import {UserMe} from '../types/user';

export function atomWithBroadcast<Value = undefined | null | object>(
  key: string,
  initializeValue: Value
) {
  const baseAtom = atom<Value>(initializeValue);

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const listeners = new Set<(event: MessageEvent<Value>) => void>();
  const channel = new BroadcastChannel(key);
  channel.onmessage = event => {
    listeners.forEach(l => l(event));
  };

  const broadcastAtom = atom<
    Value,
    [{isEvent: boolean; value: Value}],
    unknown
  >(
    get => get(baseAtom),
    (get, set, update) => {
      set(baseAtom, update.value);

      if (!update.isEvent) {
        const data = get(baseAtom);
        // undefinedの場合、別タブ側でも`/user/me`をfetchしてしまうのでnullを渡す
        if (typeof data === 'undefined') {
          channel.postMessage(null);
        } else {
          channel.postMessage(data);
        }
      }
    }
  );

  broadcastAtom.onMount = setAtom => {
    const listener = (event: MessageEvent<Value>) => {
      setAtom({isEvent: true, value: event.data});
    };
    listeners.add(listener);
    return () => {
      listeners.delete(listener);
    };
  };

  const returnedAtom = atom<Value, [Value], unknown>(
    get => get(broadcastAtom),
    (_, set, update) => {
      set(broadcastAtom, {isEvent: false, value: update});
    }
  );

  return returnedAtom;
}

export const UserState = atomWithBroadcast<UserMe | null | undefined>(
  'user',
  undefined
);
