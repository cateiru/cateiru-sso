import React from 'react';
import {GoogleReCaptchaProvider} from 'react-google-recaptcha-v3';
import {config} from '../../utils/config';

export const ReCaptcha = React.memo<{children: React.ReactNode}>(
  ({children}) => {
    if (typeof config.reCaptchaKey === 'undefined') {
      return <>{children}</>;
    }

    return (
      <GoogleReCaptchaProvider
        reCaptchaKey={config.reCaptchaKey ?? ''}
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
