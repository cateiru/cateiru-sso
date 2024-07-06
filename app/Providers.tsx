// app/providers.tsx
'use client';

import {ChakraBaseProvider} from '@chakra-ui/react';
import {Frame} from '../components/Common/Frame/Frame';
import {ReCaptcha} from '../components/Common/ReCaptcha';
import {theme, toastOptions} from '../utils/theme';

export const Providers = ({children}: {children: React.ReactNode}) => (
  <ChakraBaseProvider
    theme={theme}
    toastOptions={{defaultOptions: toastOptions}}
  >
    <ReCaptcha>
      <Frame>{children}</Frame>
    </ReCaptcha>
  </ChakraBaseProvider>
);
