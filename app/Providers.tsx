// app/providers.tsx
'use client';

import {ChakraBaseProvider} from '@chakra-ui/react';
import {RecoilRoot} from 'recoil';
import {Frame} from '../components/Common/Frame/Frame';
import {ReCaptcha} from '../components/Common/ReCaptcha';
import {theme, toastOptions} from '../utils/theme';

export const Providers = ({children}: {children: React.ReactNode}) => (
  <RecoilRoot>
    <ChakraBaseProvider
      theme={theme}
      toastOptions={{defaultOptions: toastOptions}}
    >
      <ReCaptcha>
        <Frame>{children}</Frame>
      </ReCaptcha>
    </ChakraBaseProvider>
  </RecoilRoot>
);
