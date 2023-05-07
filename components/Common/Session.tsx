import {useRouter} from 'next/router';
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
  const [ok, setOk] = React.useState(false);

  React.useEffect(() => {
    if (!router.isReady) return;
    if (ok) return;
    if (typeof user === 'undefined') return;

    let url = props.redirectTo ?? '/';
    if (props.redirectQuery) {
      url += `?redirect=${router.asPath}`;
    }

    if (!!props.isLoggedIn === !!user) {
      router.replace(url);
      return;
    }

    setOk(true);
  }, [router.isReady, ok, user]);

  return <>{ok && props.children}</>;
};
