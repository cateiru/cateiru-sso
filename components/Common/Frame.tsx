import {Box} from '@chakra-ui/react';
import React from 'react';
import {Header} from './Header';
import {useSession} from './useSession';

interface Props {
  children: React.ReactNode;
}

export const Frame = React.memo<Props>(props => {
  useSession();

  return (
    <Box w="100%" h="100%">
      <Header />
      {props.children}
    </Box>
  );
});

Frame.displayName = 'Frame';
