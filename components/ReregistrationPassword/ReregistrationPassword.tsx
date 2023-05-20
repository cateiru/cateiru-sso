'use client';

import {Center, Heading, useToast} from '@chakra-ui/react';
import {useParams, useRouter} from 'next/navigation';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {AccountReRegisterPasswordIsSessionSchema} from '../../utils/types/login';
import type {RegisterPasswordFormData} from '../Common/Form/RegisterPasswordForm';
import {Spinner} from '../Common/Icons/Spinner';
import {Margin} from '../Common/Margin';
import {useRequest} from '../Common/useRequest';
import {ReregistrationPasswordForm} from './ReregistrationPasswordForm';

export const ReregistrationPassword = () => {
  const router = useRouter();
  const params = useParams();
  const {request: availableToken} = useRequest(
    '/v2/account/reregistration/available_token'
  );
  const {request: updatePassword} = useRequest(
    '/v2/account/reregistration/password'
  );
  const {executeRecaptcha} = useGoogleReCaptcha();
  const toast = useToast();

  const [token, setToken] = React.useState<string | undefined | null>(
    undefined
  );
  const [email, setEmail] = React.useState<string>('');

  React.useEffect(() => {
    const token = params.token;
    const email = params.email;

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

  const onSubmit = async (data: RegisterPasswordFormData) => {
    if (typeof token !== 'string') return;
    if (email === '') return;
    if (!executeRecaptcha) {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const form = new FormData();
    form.append('email', email);
    form.append('reregister_token', token);
    form.append('new_password', data.new_password);

    try {
      form.append('recaptcha', await executeRecaptcha());
    } catch {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

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
        <Heading textAlign="center" mt="3rem" mb="1rem">
          新しいパスワードを設定してください
        </Heading>
        <ReregistrationPasswordForm onSubmit={onSubmit} />
      </Margin>
    );
  }

  router.replace('/');
  return null;
};
