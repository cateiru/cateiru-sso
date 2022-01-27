import {useToast} from '@chakra-ui/react';
import {useSetRecoilState} from 'recoil';
import getUserInfo from '../utils/api/userInfo';
import {UserState} from '../utils/state/atom';

export const useGetUserInfo = (): (() => void) => {
  const toast = useToast();
  const setUser = useSetRecoilState(UserState);

  const get = () => {
    const f = async () => {
      try {
        const user = await getUserInfo();
        setUser(user);
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
    };

    f();
  };

  return get;
};
