import {Box} from '@chakra-ui/react';
import React from 'react';

export const Margin: React.FC<{children: React.ReactNode}> = ({children}) => {
  return (
    <Box w={{base: '98%', md: '650px'}} mx="auto">
      {children}
    </Box>
  );
};
