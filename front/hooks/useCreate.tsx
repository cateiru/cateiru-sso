import {useToast} from '@chakra-ui/react';
import React from 'react';
import {createTemp} from '../utils/api/create';

export const useCreateTemp = (): [
  (mail: string, password: string, recaptcha: string) => void,
  string
] => {
  const toast = useToast();
  const [clientToken, setClientToken] = React.useState('');

  const create = (mail: string, password: string, recaptcha: string) => {
    const f = async () => {
      try {
        const token = await createTemp(mail, password, recaptcha);
        setClientToken(token);
      } catch (error) {
        if (error instanceof Error) {
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

  return [create, clientToken];
};
