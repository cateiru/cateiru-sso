'use client';

import {Center, Heading, Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {Margin} from '../Common/Margin';
import {useRecaptcha} from '../Common/useRecaptcha';
import {useRequest} from '../Common/useRequest';
import {ForgetForm, type ForgetFormData} from './ForgetForm';
import {SendMainSuccess} from './SendMailSuccess';

export const Forget = () => {
  const descriptionColor = useColorModeValue('gray.500', 'gray.400');
  const {request} = useRequest('/v2/account/forget/password');
  const {getRecaptchaToken} = useRecaptcha();

  const [isSuccess, setIsSuccess] = React.useState(false);
  const [email, setEmail] = React.useState('');

  const onSubmit = async (data: ForgetFormData) => {
    const form = new FormData();
    form.append('email', data.email);

    const recaptchaToken = await getRecaptchaToken();
    if (typeof recaptchaToken === 'undefined') {
      return;
    }
    form.append('recaptcha', recaptchaToken);

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
      <Heading textAlign="center" mt="15vh">
        パスワードを再登録します
      </Heading>
      <Text mt="1rem" textAlign="center" mb="2rem" color={descriptionColor}>
        入力したメールアドレスにパスワード再登録メールを送信します。
      </Text>
      <ForgetForm onSubmit={onSubmit} />
    </Margin>
  );
};
