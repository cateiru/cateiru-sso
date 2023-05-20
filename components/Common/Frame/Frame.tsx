import {Box, Flex} from '@chakra-ui/react';
import React from 'react';
import {useSession} from '../useSession';
import {Body} from './Body';
import {Footer} from './Footer';
import {Header} from './Header';

interface Props {
  children: React.ReactNode;
}

export const Frame = React.memo<Props>(props => {
  useSession();

  return (
    <Body>
      <Flex flexDirection="column" minHeight="100vh">
        <Box w="100%">
          <Header />
          {props.children}
        </Box>
        <Footer />
      </Flex>
    </Body>
  );
});

Frame.displayName = 'Frame';
