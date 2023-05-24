'use client';

import {Box} from '@chakra-ui/react';
import {OtpSetting} from './OtpSetting';
import {PasskeySetting} from './PasskeySetting';

export const AccountSetting = () => {
  return (
    <Box mt="2rem">
      <OtpSetting />
      <PasskeySetting />
    </Box>
  );
};
