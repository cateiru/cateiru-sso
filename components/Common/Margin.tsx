import {Box} from '@chakra-ui/react';
import React from 'react';

export const Margin: typeof Box = props => {
  return (
    <Box w={{base: '98%', md: '650px'}} mx="auto">
      {props.children}
    </Box>
  );
};
