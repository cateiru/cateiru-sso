import React from 'react';
import {useSetRecoilState} from 'recoil';
import {useGetUserInfo} from '../../hooks/useGetUserInfo';
import cookieValue from '../../utils/cookie';
import {UserState} from '../../utils/state/atom';

const Me: React.FC<{children: React.ReactNode}> = props => {
  const get = useGetUserInfo();
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

  return <>{props.children}</>;
};

export default Me;
