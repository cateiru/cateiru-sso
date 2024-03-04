import {SkeletonText, Text, useDisclosure} from '@chakra-ui/react';
import React from 'react';
import useSWR from 'swr';
import {userAccountCertificatesFeather} from '../../../../utils/swr/account';
import {ErrorText} from '../../../Common/Error/Error';
import {SettingCard} from '../../SettingCard';
import {OtpDisableText} from './OtpDisableText';
import {OtpEnableText} from './OtpEnableText';
import {OtpImpossible} from './OtpImpossible';
import {OtpRegister} from './OtpRegister';

export const OtpSetting = () => {
  const {isOpen, onOpen, onClose} = useDisclosure();

  const Main = () => {
    const {data, error} = useSWR(
      '/account/certificates',
      userAccountCertificatesFeather
    );

    const C = () => {
      if (error) {
        return <ErrorText {...error} />;
      }

      if (!data) {
        return <SkeletonText title="二段階認証" />;
      }

      // パスワード設定していない場合
      if (!data.password) {
        return <OtpImpossible />;
      }

      // OTP設定済み
      if (data.otp && data.otp_updated_at) {
        return <OtpEnableText updatedAt={new Date(data.otp_updated_at)} />;
      }

      // OTP無効化
      return <OtpDisableText onOpen={onOpen} />;
    };

    return (
      <SettingCard
        title="二段階認証"
        description={
          <>
            二段階認証を使用することでアカウントのセキュリティを向上させることができます。
            <br />
            パスワードを設定している場合のみ有効化できます。
          </>
        }
      >
        <C />
      </SettingCard>
    );
  };

  return (
    <>
      <Main />
      <OtpRegister isOpen={isOpen} onClose={onClose} />
    </>
  );
};
