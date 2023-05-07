import {Box, Heading} from '@chakra-ui/react';
import React from 'react';
import {RegisterAccount} from './RegisterAccount';

export const RegisterAccountPage = () => {
  return (
    <Box mt="3rem" w={{base: '96%', md: '700px'}} mx="auto">
      <Heading textAlign="center" mb="1rem">
        アカウント登録
      </Heading>
      <RegisterAccount />
    </Box>
  );
};
