import {useSetRecoilState} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {useRequest} from './useRequest';

interface Returns {
  logout: () => Promise<void>;
}

export const useLogout = (): Returns => {
  const {request} = useRequest('/v2/account/logout');
  const setUser = useSetRecoilState(UserState);

  const logout = async () => {
    const res = await request({
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      setUser(null);
    }
  };

  return {
    logout,
  };
};
