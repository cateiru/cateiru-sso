'use client';

import {Box} from '@chakra-ui/react';
import {SettingCard} from '../SettingCard';
import {OtpSetting} from './OtpSetting';

export const AccountSetting = () => {
  return (
    <Box mt="2rem">
      <SettingCard title="二段階認証">
        <OtpSetting />
      </SettingCard>
    </Box>
  );
};
