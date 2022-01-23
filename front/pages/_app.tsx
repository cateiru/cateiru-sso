import {ChakraProvider} from '@chakra-ui/react';
import type {AppProps} from 'next/app';
import {RecoilRoot} from 'recoil';
import Font from '../components/common/Font';
import Page from '../components/common/Page';
import theme from '../utils/theme/theme';

const App = ({Component, pageProps}: AppProps) => {
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
