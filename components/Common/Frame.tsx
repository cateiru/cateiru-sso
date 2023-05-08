import {Box, Flex} from '@chakra-ui/react';
import React from 'react';
import {Footer} from './Footer';
import {Header} from './Header';
import {useSession} from './useSession';

interface Props {
  children: React.ReactNode;
}

export const Frame = React.memo<Props>(props => {
  useSession();

  return (
    <Flex flexDirection="column" minHeight="100vh">
      <Box>
        <Header />
        {props.children}
      </Box>
      <Footer />
    </Flex>
  );
});

Frame.displayName = 'Frame';
