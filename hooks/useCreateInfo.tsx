import {useToast} from '@chakra-ui/react';
import {useRecoilValue} from 'recoil';
import {useSetRecoilState} from 'recoil';
import {createInfo} from '../utils/api/create';
import {CTState} from '../utils/state/atom';
import {UserState} from '../utils/state/atom';

const useCreateInfo = (): ((
  firstName: string,
  lastName: string,
  userName: string,
  theme: string,
  password: string
) => Promise<void>) => {
  const toast = useToast();
  const ct = useRecoilValue(CTState);
  const setUser = useSetRecoilState(UserState);

  const info = async (
    firstName: string,
    lastName: string,
    userName: string,
    theme: string,
    password: string
  ) => {
    try {
      if (ct) {
        const user = await createInfo(
          ct,
          firstName,
          lastName,
          userName,
          theme,
          password
        );
        setUser(user);
      } else {
        throw new Error('あれ？トークンがありませんよ');
      }
    } catch (error) {
      if (error instanceof Error) {
        toast({
          title: error.message,
          status: 'error',
          isClosable: true,
          duration: 9000,
        });
      }
    }
  };

  return info;
};

export default useCreateInfo;
