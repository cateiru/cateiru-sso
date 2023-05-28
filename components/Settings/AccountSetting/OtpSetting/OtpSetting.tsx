import {useDisclosure} from '@chakra-ui/react';
import React from 'react';
import useSWR from 'swr';
import {userAccountCertificatesFeather} from '../../../../utils/swr/featcher';
import {Error} from '../../../Common/Error/Error';
import {SettingCard} from '../../SettingCard';
import {SettingCardSkelton} from '../../SettingCardSkelton';
import {OtpDisableText} from './OtpDisableText';
import {OtpEnableText} from './OtpEnableText';
import {OtpImpossible} from './OtpImpossible';
import {OtpRegister} from './OtpRegister';

export const OtpSetting = () => {
  const {isOpen, onOpen, onClose} = useDisclosure();

  const Main = () => {
    const {data, error} = useSWR(
      '/v2/account/certificates',
      userAccountCertificatesFeather
    );

    if (error) {
      return <Error {...error} />;
    }

    if (!data) {
      return <SettingCardSkelton />;
    }

    const C = () => {
      // パスワード設定していない場合
      if (!data.password) {
        return <OtpImpossible />;
      }

      // OTP設定済み
      if (data.otp && data.otp_modified_at) {
        return <OtpEnableText modifiedAt={new Date(data.otp_modified_at)} />;
      }

      // OTP無効化
      return <OtpDisableText onOpen={onOpen} />;
    };

    return (
      <SettingCard title="二段階認証">
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
