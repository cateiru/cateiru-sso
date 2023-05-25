'use client';

import {Box} from '@chakra-ui/react';
import {DeleteAccount} from './DeleteAccount';
import {NoticeSetting} from './NoticeSetting';
import {OtpSetting} from './OtpSetting';
import {PasskeySetting} from './PasskeySetting';

export const AccountSetting = () => {
  return (
    <Box>
      <NoticeSetting />
      <OtpSetting />
      <PasskeySetting />
      <DeleteAccount />
    </Box>
  );
};
