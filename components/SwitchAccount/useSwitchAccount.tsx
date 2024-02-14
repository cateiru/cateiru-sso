import {useToast} from '@chakra-ui/react';
import {useSetAtom} from 'jotai';
import {useRouter, useSearchParams} from 'next/navigation';
import nProgress from 'nprogress';
import React from 'react';
import {formatRedirectUrl} from '../../utils/format';
import {UserState} from '../../utils/state/atom';
import {useRequest} from '../Common/useRequest';

nProgress.configure({showSpinner: false, speed: 400, minimum: 0.25});

interface Returns {
  switch: (id: string, name: string) => void;
  loading: string | null;
  redirect: (id: string) => void;
}

export const useSwitchAccount = (): Returns => {
  const setUser = useSetAtom(UserState);
  const {request} = useRequest('/account/switch', {
    errorCallback: () => {
      nProgress.done();
    },
  });
  const router = useRouter();
  const params = useSearchParams();
  const toast = useToast();
  const [loading, setLoading] = React.useState<null | string>(null);

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
        setLoading(id);

        setTimeout(async () => {
          setUser(undefined);
          toast({
            title: `ユーザー ${name} にログインしました`,
            status: 'success',
          });
          onRedirect();
          setLoading(null);
        }, 500);
        return;
      }
      nProgress.done();
    };
    f();
  };

  const onRedirect = () => {
    const redirectTo = params.get('redirect_to');
    if (typeof redirectTo === 'string') {
      router.push(formatRedirectUrl(redirectTo));
    } else {
      router.push('/profile');
    }
  };

  const redirect = (id: string) => {
    nProgress.start();
    setLoading(id);

    onRedirect();

    setLoading(null);
    nProgress.done();
  };

  return {
    loading,
    switch: switchAccount,
    redirect,
  };
};
