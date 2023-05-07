import React from 'react';
import {GoogleReCaptchaProvider} from 'react-google-recaptcha-v3';

export const ReCaptcha = React.memo<{children: React.ReactNode}>(
  ({children}) => {
    return (
      <GoogleReCaptchaProvider
        reCaptchaKey={
          process.env.NEXT_PUBLIC_RE_CAPTCHA ?? 'empty_recaptcha_key'
        }
        scriptProps={{
          async: false,
          defer: false,
          appendTo: 'head',
          nonce: undefined,
        }}
      >
        {children}
      </GoogleReCaptchaProvider>
    );
  }
);

ReCaptcha.displayName = 'ReCaptcha';
