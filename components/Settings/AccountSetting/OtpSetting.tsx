import {Skeleton} from '@chakra-ui/react';
import useSWR from 'swr';
import {userOtpFeather} from '../../../utils/swr/featcher';
import {Error} from '../../Common/Error/Error';
import {OtpDisableText} from './OtpDisableText';
import {OtpEnableText} from './OtpEnableText';

export const OtpSetting = () => {
  const {data, error} = useSWR('/v2/user/otp', userOtpFeather);

  if (error) {
    return <Error {...error} />;
  }

  if (!data) {
    return <Skeleton w="100%" h="40px" />;
  }

  if (data.enable && data.modified) {
    return <OtpEnableText modified={new Date(data.modified)} />;
  }

  // OTP無効化
  return <OtpDisableText />;
};
