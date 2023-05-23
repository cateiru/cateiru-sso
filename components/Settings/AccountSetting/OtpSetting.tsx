import React from 'react';
import {OtpDisableText} from './OtpDisableText';
import {OtpEnableText} from './OtpEnableText';
import {OtpImpossible} from './OtpImpossible';

interface Props {
  settingPassword: boolean;
  settingOtp: boolean;
  otpModified: string | null;
}

export const OtpSetting: React.FC<Props> = props => {
  // パスワード設定していない場合
  if (!props.settingPassword) {
    return <OtpImpossible />;
  }

  // OTP設定済み
  if (props.settingOtp && props.otpModified) {
    return <OtpEnableText modified={new Date(props.otpModified)} />;
  }

  // OTP無効化
  return <OtpDisableText />;
};
