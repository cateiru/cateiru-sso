'use client';

import {Box} from '@chakra-ui/react';
import {DeleteAccount} from './DeleteAccount';
import {NoticeSetting} from './NoticeSetting';
import {OtpSetting} from './OtpSetting';
import {WebAuthn} from './WebAuthn';

export const AccountSetting = () => {
  return (
    <Box>
      <NoticeSetting />
      <OtpSetting />
      <WebAuthn />
      <DeleteAccount />
    </Box>
  );
};
