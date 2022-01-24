import {useToast} from '@chakra-ui/react';
import React from 'react';
import {createVerify} from '../utils/api/create';

const useVerify = (): [(t: string) => void, boolean, string, boolean] => {
  const [isKeep, setIsKeep] = React.useState(false);
  const [token, setToken] = React.useState('');
  const [err, setError] = React.useState(false);
  const toast = useToast();

  const verify = (t: string) => {
    const f = async () => {
      try {
        const resp = await createVerify(t);
        setToken(resp.client_token);
        setIsKeep(resp.keep_this_page);
      } catch (error) {
        setError(true);
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

  return [verify, isKeep, token, err];
};

export default useVerify;
