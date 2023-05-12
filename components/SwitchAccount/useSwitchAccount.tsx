import {useToast} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import nProgress from 'nprogress';
import {useSetRecoilState} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {useRequest} from '../Common/useRequest';

interface Returns {
  switch: (id: string, name: string) => void;
}

export const useSwitchAccount = (): Returns => {
  const setUser = useSetRecoilState(UserState);
  const {request} = useRequest('/v2/account/switch', {
    errorCallback: () => {
      nProgress.done();
    },
  });
  const router = useRouter();
  const toast = useToast();

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
        setTimeout(async () => {
          setUser(undefined);
          toast({
            title: `ユーザー ${name} にログインしました`,
            status: 'success',
          });
          await router.push('/');
        }, 500);
        return;
      }
      nProgress.done();
    };
    f();
  };

  return {
    switch: switchAccount,
  };
};