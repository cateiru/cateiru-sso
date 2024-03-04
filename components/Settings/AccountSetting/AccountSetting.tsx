'use client';

import {Box} from '@chakra-ui/react';
import {DeleteAccount} from './DeleteAccount';
import {OtpSetting} from './OtpSetting';
import {WebAuthn} from './WebAuthn';

export const AccountSetting = () => {
  return (
    <Box>
      <OtpSetting />
      <WebAuthn />
      <DeleteAccount />
    </Box>
  );
};
