'use client';

import {Box} from '@chakra-ui/react';
import {DeleteAccount} from './DeleteAccount';
import {OtpSetting} from './OtpSetting';
import {PasskeySetting} from './PasskeySetting';
import {UserSetting} from './UserSetting';

export const AccountSetting = () => {
  return (
    <Box>
      <UserSetting />
      <OtpSetting />
      <PasskeySetting />
      <DeleteAccount />
    </Box>
  );
};
