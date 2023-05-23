'use client';

import {Box, useToast} from '@chakra-ui/react';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {useRecoilState} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {ErrorUniqueMessage} from '../../utils/types/error';
import {
  UserUpdateEmailRegisterScheme,
  UserUpdateEmailScheme,
} from '../../utils/types/settings';
import {EmailVerifyForm} from '../Common/Form/EmailVerifyForm';
import {useRequest} from '../Common/useRequest';
import {EmailFormData, EmailSettingForm} from './EmailSettingForm';
import {EmailSettingVerifyForm} from './EmailSettingVerifyForm';

export const EmailSetting = () => {
  const [user, setUser] = useRecoilState(UserState);
  const {executeRecaptcha} = useGoogleReCaptcha();
  const toast = useToast();

  const [disabled, setDisabled] = React.useState(false);
  const [token, setToken] = React.useState('');

  const {request: updateEmail} = useRequest('/v2/user/email');
  const {request: verifyEmail} = useRequest('/v2/user/email/register', {
    customError: error => {
      const message = error.unique_code
        ? ErrorUniqueMessage[error.unique_code] ?? error.message
        : error.message;
      toast({
        title: message,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });

      if (error.unique_code !== 13) {
        // 認証失敗以外のエラーはリセット
        setDisabled(false);
        setToken('');
      }
    },
  });

  const onSubmit = async (data: EmailFormData) => {
    if (!executeRecaptcha) {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const form = new FormData();
    form.append('new_email', data.new_email);

    try {
      form.append('recaptcha', await executeRecaptcha());
    } catch {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const res = await updateEmail({
      method: 'POST',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      const data = UserUpdateEmailScheme.safeParse(await res.json());
      if (data.success) {
        setToken(data.data.session);
        setDisabled(true);
      } else {
        console.error(data.error);
      }
    }
  };

  const onSubmitVerify = async (data: EmailVerifyForm) => {
    const form = new FormData();
    form.append('update_token', token);
    form.append('code', data.code);

    const res = await verifyEmail({
      method: 'POST',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      const data = UserUpdateEmailRegisterScheme.safeParse(await res.json());
      if (data.success) {
        toast({
          title: 'メールアドレスを変更しました',
          status: 'success',
        });

        setDisabled(false);
        setToken('');

        // Email更新
        setUser(v => {
          if (v) {
            return {
              ...v,
              user: {
                ...v.user,
                email: data.data.email,
              },
            };
          }
          return v;
        });
      } else {
        console.error(data.error);
      }
    }
  };

  return (
    <Box mt="2rem" minH="80vh">
      <EmailSettingForm
        onSubmit={onSubmit}
        disabled={disabled}
        email={user?.user.email ?? ''}
      />
      {disabled && <EmailSettingVerifyForm onSubmit={onSubmitVerify} />}
    </Box>
  );
};