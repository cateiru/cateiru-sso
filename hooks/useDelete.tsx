import {useToast} from '@chakra-ui/react';
import {useSetRecoilState} from 'recoil';
import {deleteAccount} from '../utils/api/logout';
import {UserState} from '../utils/state/atom';

const useDelete = (): (() => Promise<void>) => {
  const setUser = useSetRecoilState(UserState);
  const toast = useToast();

  const f = async () => {
    try {
      await deleteAccount();
      setUser(null);
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

  return f;
};

export default useDelete;
