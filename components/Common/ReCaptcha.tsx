import {useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {GoogleReCaptchaProvider} from 'react-google-recaptcha-v3';

export const ReCaptcha: React.FC<{children: React.ReactNode}> = ({
  children,
}) => {
  const theme = useColorModeValue('light', 'dark');

  return (
    <GoogleReCaptchaProvider
      reCaptchaKey={process.env.NEXT_PUBLIC_RE_CAPTCHA ?? 'empty_recaptcha_key'}
      scriptProps={{
        async: false,
        defer: false,
        appendTo: 'head',
        nonce: undefined,
      }}
      container={{
        parameters: {
          theme: theme,
        },
      }}
    >
      {children}
    </GoogleReCaptchaProvider>
  );
};
