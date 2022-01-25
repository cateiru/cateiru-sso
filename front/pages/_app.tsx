import {ChakraProvider} from '@chakra-ui/react';
import type {AppProps} from 'next/app';
import Router, {useRouter} from 'next/router';
import nprogress from 'nprogress';
import {useEffect} from 'react';
import {RecoilRoot} from 'recoil';
import Font from '../components/common/Font';
import Page from '../components/common/Page';
import {GA_TRACKING_ID, pageview} from '../utils/ga/gtag';
import theme from '../utils/theme/theme';
import 'nprogress/nprogress.css';

nprogress.configure({showSpinner: false, speed: 400, minimum: 0.25});

const App = ({Component, pageProps}: AppProps) => {
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
        <Font />
        <Page>
          <Component {...pageProps} />
        </Page>
      </ChakraProvider>
    </RecoilRoot>
  );
};

export default App;
