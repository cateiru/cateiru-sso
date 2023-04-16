import {useRouter} from 'next/router';
import React from 'react';
import {UserMe} from '../../utils/types/user';
import {useSession} from './useSession';

interface Props {
  children: (user: UserMe) => React.ReactNode;
}

export const Session: React.FC<Props> = props => {
  const {user, noLogin} = useSession();
  const router = useRouter();

  if (noLogin) {
    return null;
  }

  if (!user) {
    router.replace(`/login?redirect=${router.asPath}`);
    return null;
  }

  return <>{props.children(user)}</>;
};
