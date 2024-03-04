import {useToast} from '@chakra-ui/react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {config} from '../../utils/config';

interface Returns {
  getRecaptchaToken: () => Promise<string | undefined>;
}

export const useRecaptcha = (): Returns => {
  const {executeRecaptcha} = useGoogleReCaptcha();
  const toast = useToast();

  const getRecaptchaToken = async (): Promise<string | undefined> => {
    if (typeof config.reCaptchaKey === 'undefined') {
      return 'dummy-recaptcha-key';
    }

    if (!executeRecaptcha) {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return undefined;
    }

    try {
      return await executeRecaptcha();
    } catch {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return undefined;
    }
  };

  return {
    getRecaptchaToken,
  };
};
