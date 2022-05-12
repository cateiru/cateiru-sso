import {useToast} from '@chakra-ui/react';
import {useSetRecoilState} from 'recoil';
import {ping} from '../utils/api/ping';
import getUserInfo from '../utils/api/userInfo';
import {UserState} from '../utils/state/atom';

export const useGetUserInfo = (): (() => void) => {
  const toast = useToast();
  const setUser = useSetRecoilState(UserState);

  const beforeUnload = (event: Event) => {
    event.preventDefault();

    // 必須: https://developer.mozilla.org/ja/docs/Web/API/Window/beforeunload_event
    // > しかし、すべてのブラウザーがこのメソッドに対応しているわけではなく、
    // > 一部はイベントハンドラーに古い方法二つのうちの一つを実装するよう求めていることに注意してください。
    // >
    // > - イベントの returnValue プロパティに文字列を代入する
    // > - イベントハンドラーから文字列を返す
    event.returnValue = true;
  };

  const get = () => {
    const f = async () => {
      try {
        window.addEventListener('beforeunload', beforeUnload);

        // cookieが消えないようにするため初回ping出す
        await ping();

        const user = await getUserInfo();
        setUser(user);

        // TODO: ユーザーアクセシビリティが低下するので
        // if (user) {
        //   switch (user.theme) {
        //     case 'dark':
        //       setColorMode('dark');
        //       break;
        //     case 'light':
        //       setColorMode('light');
        //       break;
        //     default:
        //       break;
        //   }
        // }
      } catch (error) {
        if (error instanceof Error) {
          setUser(null);
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }
      window.removeEventListener('beforeunload', beforeUnload);
    };

    f();
  };

  return get;
};
