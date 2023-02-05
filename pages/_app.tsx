import {ChakraProvider} from '@chakra-ui/react';
import type {AppProps} from 'next/app';
import Router, {useRouter} from 'next/router';
import nprogress from 'nprogress';
import {useEffect} from 'react';
import React from 'react';
import {GoogleReCaptchaProvider} from 'react-google-recaptcha-v3';
import {RecoilRoot} from 'recoil';
import Load from '../components/common/Load';
import Me from '../components/common/Me';
import Page from '../components/common/Page';
import {GA_TRACKING_ID, pageview} from '../utils/ga/gtag';
import theme from '../utils/theme/theme';

import '@fontsource/noto-sans-jp';
import 'nprogress/nprogress.css';

nprogress.configure({showSpinner: false, speed: 400, minimum: 0.25});

const reCAPTCHA = process.env.NEXT_PUBLIC_RE_CAPTCHA;

const App = ({Component, pageProps}: AppProps) => {
  const [scriptProps] = React.useState<{
    nonce?: string;
    defer?: boolean;
    async?: boolean;
    appendTo?: 'head' | 'body';
    id?: string;
  }>({
    async: false,
    defer: false,
    appendTo: 'head',
  });

  const router = useRouter();

  useEffect(() => {
    if (!GA_TRACKING_ID) return;

    const handleRouteChange = (url: string) => {
      pageview(url);
    };
    router.events.on('routeChangeComplete', handleRouteChange);
    return () => {
      router.events.off('routeChangeComplete', handleRouteChange);
    };
  }, [router.events]);

  Router.events.on('routeChangeStart', () => {
    nprogress.start();
  });

  Router.events.on('routeChangeComplete', () => {
    nprogress.done();
  });

  Router.events.on('routeChangeError', () => {
    nprogress.done();
  });

  return (
    <RecoilRoot>
      <ChakraProvider theme={theme}>
        <GoogleReCaptchaProvider
          reCaptchaKey={reCAPTCHA}
          language="ja"
          scriptProps={scriptProps}
        >
          <Load />
          <Me>
            <Page>
              <Component {...pageProps} />
            </Page>
          </Me>
        </GoogleReCaptchaProvider>
      </ChakraProvider>
    </RecoilRoot>
  );
};

export default App;
