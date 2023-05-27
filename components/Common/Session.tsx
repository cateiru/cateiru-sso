'use client';

import {usePathname, useRouter} from 'next/navigation';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {config} from '../../utils/config';
import {UserState} from '../../utils/state/atom';

interface Props {
  isLoggedIn?: boolean;
  redirectTo?: string;
  redirectQuery?: boolean;
  isStaff?: boolean;
  children: React.ReactNode;
}

export const Session: React.FC<Props> = props => {
  const user = useRecoilValue(UserState);
  const router = useRouter();
  const pathname = usePathname();
  const [ok, setOk] = React.useState(false);

  React.useEffect(() => {
    if (typeof user === 'undefined') return;

    let url = props.redirectTo ?? (user ? '/profile' : '/');
    if (props.redirectQuery) {
      url += `?redirect=${pathname}`;
    }

    if (!!props.isLoggedIn === !!user || props.isStaff === !user?.is_staff) {
      if (config.mode === 'development') {
        console.log('Redirect to: ', url);
      }

      router.replace(url);
      return;
    }

    setOk(true);
  }, [user]);

  return <>{ok && props.children}</>;
};
