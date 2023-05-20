'use client';

import {
  Center,
  Heading,
  Text,
  useColorModeValue,
  useToast,
} from '@chakra-ui/react';
import React from 'react';
import {useGoogleReCaptcha} from 'react-google-recaptcha-v3';
import {Margin} from '../Common/Margin';
import {useRequest} from '../Common/useRequest';
import {ForgetForm, type ForgetFormData} from './ForgetForm';
import {SendMainSuccess} from './SendMailSuccess';

export const Forget = () => {
  const descriptionColor = useColorModeValue('gray.500', 'gray.400');
  const {request} = useRequest('/v2/account/forget/password');
  const {executeRecaptcha} = useGoogleReCaptcha();
  const toast = useToast();

  const [isSuccess, setIsSuccess] = React.useState(false);
  const [email, setEmail] = React.useState('');

  const onSubmit = async (data: ForgetFormData) => {
    if (!executeRecaptcha) {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const form = new FormData();
    form.append('email', data.email);

    try {
      form.append('recaptcha', await executeRecaptcha());
    } catch {
      toast({
        title: 'reCAPTCHAの読み込みに失敗しました',
        status: 'error',
      });
      return;
    }

    const res = await request({
      method: 'POST',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      setEmail(data.email);
      setIsSuccess(true);
    }
  };

  if (isSuccess) {
    return (
      <Center h="80vh">
        <SendMainSuccess email={email} />
      </Center>
    );
  }

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        パスワードを再登録します
      </Heading>
      <Text mt="1rem" textAlign="center" mb="2rem" color={descriptionColor}>
        入力したメールアドレスにパスワード再登録メールを送信します。
      </Text>
      <ForgetForm onSubmit={onSubmit} />
    </Margin>
  );
};
