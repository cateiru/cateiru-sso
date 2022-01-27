import {useToast} from '@chakra-ui/react';
import {useRecoilValue} from 'recoil';
import {createInfo} from '../utils/api/create';
import {CTState} from '../utils/state/atom';

const useCreateInfo = (): ((
  firstName: string,
  lastName: string,
  userName: string,
  theme: string
) => void) => {
  const toast = useToast();
  const ct = useRecoilValue(CTState);

  const info = (
    firstName: string,
    lastName: string,
    userName: string,
    theme: string
  ) => {
    const f = async () => {
      try {
        if (ct) {
          createInfo(ct, firstName, lastName, userName, theme);
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

    f();
  };

  return info;
};

export default useCreateInfo;
