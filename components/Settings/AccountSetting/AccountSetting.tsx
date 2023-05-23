'use client';

import {Box, Skeleton} from '@chakra-ui/react';
import useSWR from 'swr';
import {userAccountCertificatesFeather} from '../../../utils/swr/featcher';
import {Error} from '../../Common/Error/Error';
import {SettingCard} from '../SettingCard';
import {OtpSetting} from './OtpSetting';

export const AccountSetting = () => {
  const {data, error} = useSWR(
    '/v2/account/certificates',
    userAccountCertificatesFeather
  );

  if (error) {
    return <Error {...error} />;
  }

  if (!data) {
    // TODO
    return <Skeleton w="100%" h="40px" />;
  }

  return (
    <Box mt="2rem">
      <SettingCard title="二段階認証">
        <OtpSetting
          settingOtp={data.otp}
          otpModified={data.otp_modified}
          settingPassword={data.password}
        />
      </SettingCard>
    </Box>
  );
};
