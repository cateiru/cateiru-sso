import {ChakraProvider} from '@chakra-ui/react';
import {AppProps} from 'next/app';
import {useRouter} from 'next/router';
import nprogress from 'nprogress';
import React from 'react';
import {RecoilRoot} from 'recoil';
import {ReCaptcha} from '../components/Common/ReCaptcha';
import {pageview} from '../utils/ga/gtag';
import {theme} from '../utils/theme';

import '../public/nprogress.css';

nprogress.configure({showSpinner: false, speed: 400, minimum: 0.25});

const App = ({Component, pageProps}: AppProps) => {
  const router = useRouter();

  React.useEffect(() => {
    const handleRouteChange = (url: string) => {
      pageview(url);
    };
    router.events.on('routeChangeComplete', handleRouteChange);

    router.events.on('routeChangeStart', () => {
      nprogress.start();
    });

    router.events.on('routeChangeComplete', () => {
      nprogress.done();
    });

    router.events.on('routeChangeError', () => {
      nprogress.done();
    });
    return () => {
      router.events.off('routeChangeComplete', handleRouteChange);
    };
  }, [router.events]);

  return (
    <RecoilRoot>
      <ChakraProvider theme={theme}>
        <ReCaptcha>
          <Component {...pageProps} />
        </ReCaptcha>
      </ChakraProvider>
    </RecoilRoot>
  );
};

export default App;
