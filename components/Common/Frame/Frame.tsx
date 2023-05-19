import {Box, Flex} from '@chakra-ui/react';
import React from 'react';
import {useSession} from '../useSession';
import {Footer} from './Footer';
import {Header} from './Header';

interface Props {
  children: React.ReactNode;
}

export const Frame = React.memo<Props>(props => {
  useSession();

  return (
    <Box minHeight="100vh">
      <Box>
        <Header />
        {props.children}
      </Box>
      <Footer />
    </Box>
  );
});

Frame.displayName = 'Frame';
