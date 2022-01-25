import {useToast} from '@chakra-ui/react';
import React from 'react';
import getUserInfo, {UserInfo} from '../utils/api/userInfo';

export const useGetUserInfo = (): [
  () => void,
  UserInfo | undefined,
  boolean
] => {
  const toast = useToast();
  const [user, setUser] = React.useState<UserInfo>();
  const [err, setError] = React.useState(false);

  const get = () => {
    const f = async () => {
      try {
        const user = await getUserInfo();
        setUser(user);
      } catch (error) {
        if (error instanceof Error) {
          setError(true);
          toast({
            title: 'エラー',
            description: error.message,
            status: 'error',
            isClosable: true,
          });
        }
      }
    };

    f();
  };

  return [get, user, err];
};
