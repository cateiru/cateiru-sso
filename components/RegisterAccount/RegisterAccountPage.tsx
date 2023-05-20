'use client';

import {Heading} from '@chakra-ui/react';
import React from 'react';
import {Margin} from '../Common/Margin';
import {RegisterAccount} from './RegisterAccount';

export const RegisterAccountPage = () => {
  return (
    <Margin>
      <Heading textAlign="center" mb="1rem" mt="3rem">
        アカウント登録
      </Heading>
      <RegisterAccount />
    </Margin>
  );
};
