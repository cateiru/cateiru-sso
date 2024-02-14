import {useSetAtom} from 'jotai';
import {UserState} from '../../utils/state/atom';
import {useRequest} from './useRequest';

interface Returns {
  logout: () => Promise<void>;
}

export const useLogout = (): Returns => {
  const {request} = useRequest('/account/logout');
  const setUser = useSetAtom(UserState);

  const logout = async () => {
    const res = await request({
      method: 'POST',
    });

    if (res) {
      // FedCMのために、ブラウザにログアウト状態を伝える
      // まだ提案段階の使用なのでanyで無理やり適用している
      // ref. https://github.com/fedidcg/login-status
      // ref2. https://developers.google.com/privacy-sandbox/blog/fedcm-chrome-120-updates?hl=ja
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      const login = (navigator as any).login;
      if (typeof login !== 'undefined') {
        login.setStatus('logged-out');
      }

      setUser(null);
    }
  };

  return {
    logout,
  };
};
