import {useToast} from '@chakra-ui/react';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {useSetRecoilState} from 'recoil';
import {createTemp} from '../utils/api/create';
import {CTState} from '../utils/state/atom';

export const useCreateTemp = (): [(mail: string) => void, boolean] => {
  const toast = useToast();
  const [err, setError] = React.useState(false);
  const setCT = useSetRecoilState(CTState);
  const {executeRecaptcha} = useGoogleReCaptcha();

  const create = (mail: string) => {
    const f = async () => {
      if (!executeRecaptcha) {
        return;
      }

      const recaptcha = await executeRecaptcha();

      try {
        const token = await createTemp(mail, recaptcha);
        setCT(token);
      } catch (error) {
        if (error instanceof Error) {
          setError(true);
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

  return [create, err];
};
