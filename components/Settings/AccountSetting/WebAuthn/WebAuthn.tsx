'use client';

import {SettingCard} from '../../SettingCard';
import {RegisterWebAuthn} from './RegisterWebAuthn';
import {WebAuthnDevices} from './WebAuthnDevices';

export const WebAuthn = () => {
  return (
    <SettingCard
      title="生体認証（パスキー）"
      description={
        <>
          パスキーを使用するとTouchIDやFaceIDなどの生体認証でログインすることが可能です。
        </>
      }
    >
      <RegisterWebAuthn />
      <WebAuthnDevices />
    </SettingCard>
  );
};
