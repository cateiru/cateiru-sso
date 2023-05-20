'use client';

import {usePathname, useRouter} from 'next/navigation';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';

interface Props {
  isLoggedIn?: boolean;
  redirectTo?: string;
  redirectQuery?: boolean;
  children: React.ReactNode;
}

export const Session: React.FC<Props> = props => {
  const user = useRecoilValue(UserState);
  const router = useRouter();
  const pathname = usePathname();
  const [ok, setOk] = React.useState(false);

  React.useEffect(() => {
    if (typeof user === 'undefined') return;

    let url = props.redirectTo ?? '/';
    if (props.redirectQuery) {
      url += `?redirect=${pathname}`;
    }

    if (!!props.isLoggedIn === !!user) {
      router.replace(url);
      return;
    }

    setOk(true);
  }, [user]);

  return <>{ok && props.children}</>;
};
