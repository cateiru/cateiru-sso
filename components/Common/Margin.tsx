import {Box} from '@chakra-ui/react';
import React from 'react';

export const Margin: React.FC<{children: React.ReactNode}> = ({children}) => {
  return (
    <Box w={{base: '96%', md: '700px'}} mx="auto">
      {children}
    </Box>
  );
};
