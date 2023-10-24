'use client';

import {usePathname, useRouter} from 'next/navigation';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {config} from '../../utils/config';
import {UserState} from '../../utils/state/atom';

interface Props {
  isLoggedIn?: boolean;
  redirectTo?: string; // 未ログイン時のリダイレクト先
  redirectQuery?: boolean; // リダイレクト時に`redirect`クエリを渡す
  isStaff?: boolean; // スタッフの場合のみchildrenを表示させる
  noRedirect?: boolean; // 未認証でもリダイレクトしないでchildrenを表示する
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

    if (
      !props.noRedirect &&
      (!!props.isLoggedIn === !!user || props.isStaff === !user?.is_staff)
    ) {
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
