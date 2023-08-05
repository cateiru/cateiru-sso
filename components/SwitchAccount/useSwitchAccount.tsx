import {useToast} from '@chakra-ui/react';
import {useRouter, useSearchParams} from 'next/navigation';
import nProgress from 'nprogress';
import React from 'react';
import {useSetRecoilState} from 'recoil';
import {formatRedirectUrl} from '../../utils/format';
import {UserState} from '../../utils/state/atom';
import {useRequest} from '../Common/useRequest';

nProgress.configure({showSpinner: false, speed: 400, minimum: 0.25});

interface Returns {
  switch: (id: string, name: string) => void;
  loading: boolean;
}

export const useSwitchAccount = (): Returns => {
  const setUser = useSetRecoilState(UserState);
  const {request} = useRequest('/v2/account/switch', {
    errorCallback: () => {
      nProgress.done();
    },
  });
  const router = useRouter();
  const params = useSearchParams();
  const toast = useToast();
  const [loading, setLoading] = React.useState(false);

  const switchAccount = (id: string, name: string) => {
    const f = async () => {
      nProgress.start();

      const form = new FormData();
      form.append('user_id', id);

      const res = await request({
        method: 'POST',
        credentials: 'include',
        mode: 'cors',
        body: form,
      });

      if (res) {
        setLoading(true);

        setTimeout(async () => {
          setUser(undefined);
          toast({
            title: `ユーザー ${name} にログインしました`,
            status: 'success',
          });
          const redirectTo = params.get('redirect_to');
          if (typeof redirectTo === 'string') {
            router.push(formatRedirectUrl(redirectTo));
          } else {
            router.push('/profile');
          }
          setLoading(false);
        }, 500);
        return;
      }
      nProgress.done();
    };
    f();
  };

  return {
    loading,
    switch: switchAccount,
  };
};
