import {useToast} from '@chakra-ui/react';
import React from 'react';
import {useSetRecoilState} from 'recoil';
import {createVerify} from '../utils/api/create';
import {CTState} from '../utils/state/atom';

const useVerify = (): [(t: string) => void, boolean, boolean, boolean] => {
  const [isKeep, setIsKeep] = React.useState(false);
  const [load, setLoad] = React.useState(true);
  const [err, setError] = React.useState(false);
  const setCT = useSetRecoilState(CTState);
  const toast = useToast();

  const verify = (t: string) => {
    const f = async () => {
      try {
        const resp = await createVerify(t);
        setCT(resp.client_token);
        setIsKeep(resp.keep_this_page);
        setLoad(false);
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

  return [verify, isKeep, load, err];
};

export default useVerify;
