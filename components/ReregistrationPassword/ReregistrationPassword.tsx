'use client';

import {Center, Heading} from '@chakra-ui/react';
import {useRouter, useSearchParams} from 'next/navigation';
import React from 'react';
import {AccountReRegisterPasswordIsSessionSchema} from '../../utils/types/login';
import {RegisterPasswordForm} from '../Common/Form/RegisterPasswordForm';
import type {RegisterPasswordFormContextData} from '../Common/Form/RegisterPasswordFormContext';
import {Spinner} from '../Common/Icons/Spinner';
import {Margin} from '../Common/Margin';
import {useRecaptcha} from '../Common/useRecaptcha';
import {useRequest} from '../Common/useRequest';

export const ReregistrationPassword = () => {
  const router = useRouter();
  const params = useSearchParams();
  const {request: availableToken} = useRequest(
    '/v2/account/reregistration/available_token'
  );
  const {request: updatePassword} = useRequest(
    '/v2/account/reregistration/password'
  );
  const {getRecaptchaToken} = useRecaptcha();

  const [token, setToken] = React.useState<string | undefined | null>(
    undefined
  );
  const [email, setEmail] = React.useState<string>('');

  React.useEffect(() => {
    const token = params.get('token');
    const email = params.get('email');

    if (typeof token === 'string' && typeof email === 'string') {
      // 一度tokenを送信して、正しいか検証する
      const f = async () => {
        const form = new FormData();
        form.append('email', email);
        form.append('reregister_token', token);

        const res = await availableToken({
          method: 'POST',
          body: form,
          mode: 'cors',
          credentials: 'include',
        });

        if (res) {
          const data = AccountReRegisterPasswordIsSessionSchema.safeParse(
            await res.json()
          );

          if (data.success) {
            if (data.data.active) {
              // トークンの有効性を確認できたら埋める
              setToken(token);
              setEmail(email);
              return;
            }
          } else {
            console.error(data.error);
          }
        }

        // トークンの有効性を確認できなかったら無効化
        setToken(null);
      };
      f();
    } else {
      // クエリが正しくないときは無効化
      setToken(null);
    }
  }, []);

  const onSubmit = async (data: RegisterPasswordFormContextData) => {
    if (typeof token !== 'string') return;
    if (email === '') return;

    const form = new FormData();
    form.append('email', email);
    form.append('reregister_token', token);
    form.append('new_password', data.new_password);

    const recaptchaToken = await getRecaptchaToken();
    if (typeof recaptchaToken === 'undefined') {
      return;
    }
    form.append('recaptcha', recaptchaToken);

    const res = await updatePassword({
      method: 'POST',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      // パスワード再設定に成功したらログインページにリダイレクトする
      router.replace('/login');
    }
  };

  // ロード中
  if (typeof token === 'undefined') {
    return (
      <Center h="80vh">
        <Spinner size="xl" />
      </Center>
    );
  }

  // OK
  if (token) {
    return (
      <Margin>
        <Heading textAlign="center" mt="15vh" mb="1rem">
          新しいパスワードを設定してください
        </Heading>
        <RegisterPasswordForm
          onSubmit={onSubmit}
          buttonText="パスワード再設定"
        />
      </Margin>
    );
  }

  router.replace('/');
  return null;
};
