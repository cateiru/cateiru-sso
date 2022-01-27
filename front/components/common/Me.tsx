import React from 'react';
import {useSetRecoilState} from 'recoil';
import {useGetUserInfo} from '../../hooks/useGetUserInfo';
import cookieValue from '../../utils/cookie';
import {UserState} from '../../utils/state/atom';

const Me: React.FC = props => {
  const [get, user, err] = useGetUserInfo();
  const setUser = useSetRecoilState(UserState);

  React.useEffect(() => {
    if (process.browser) {
      const refresh = cookieValue('refresh-token');

      if (refresh) {
        get();
      } else {
        setUser(null);
      }
    }
  }, []);

  React.useEffect(() => {
    if (user) {
      setUser(user);
    } else if (err) {
      setUser(null);
    }
  }, [user, err]);

  return <>{props.children}</>;
};

export default Me;
