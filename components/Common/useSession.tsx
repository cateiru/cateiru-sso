import {useToast} from '@chakra-ui/react';
import cookie from 'cookie';
import React from 'react';
import {useRecoilState} from 'recoil';
import {api} from '../../utils/api';
import {config} from '../../utils/config';
import {UserState} from '../../utils/state/atom';
import {ErrorSchema, ErrorUniqueMessage} from '../../utils/types/error';
import {UserMeSchema} from '../../utils/types/user';

export const useSession = () => {
  const [user, setUser] = useRecoilState(UserState);
  const toast = useToast();

  React.useEffect(() => {
    // 未ログイン
    if (typeof user === 'undefined') {
      const isLogin =
        cookie.parse(document.cookie)[config.loginStateCookieName] === '1';

      // クッキーが存在していない場合はログインしていないのでnullにしてなにもしない
      if (!isLogin) {
        setUser(null);
        return;
      }

      fetch(api('/v2/user/me'), {
        method: 'GET',
        credentials: 'include',
        mode: 'cors',
      })
        .then(async res => {
          if (res.ok) {
            const parsedUserMe = UserMeSchema.safeParse(await res.json());
            if (parsedUserMe.success) {
              setUser(parsedUserMe.data);
            }
          } else {
            const error = ErrorSchema.safeParse(await res.json());
            if (error.success) {
              toast({
                title: 'ログインに失敗しました',
                description:
                  ErrorUniqueMessage[error.data.unique_code] ??
                  error.data.message,
                status: 'error',
                duration: 9000,
                isClosable: true,
              });
            }
            setUser(null);
          }
        })
        .catch(e => {
          if (e instanceof Error) {
            toast({
              title: 'ログインに失敗しました',
              description: e.message,
              status: 'error',
              duration: 9000,
              isClosable: true,
            });
          }
        });
    }
  }, [user]);
};
