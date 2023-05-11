import {Avatar, Center, Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import type {LoginUser} from '../../utils/types/login';
import {Margin} from '../Common/Margin';
import {type DefaultPageProps} from './Login';
import {OtpForm, OtpFormData} from './OtpForm';

interface Props extends DefaultPageProps {
  loginUser: LoginUser;
  otpToken: string;
}

export const OtpPage: React.FC<Props> = props => {
  const descriptionTextColor = useColorModeValue('gray.500', 'gray.400');
  const accentColor = useColorModeValue('my.primary', 'my.secondary');
  const onSubmit = async (data: OtpFormData) => {};

  return (
    <Margin>
      <Center mt="1rem">
        <Avatar src={props.loginUser.avatar ?? ''} size="lg" />
      </Center>
      <Text fontSize="1.5rem" fontWeight="bold" textAlign="center">
        <Text as="span" color={accentColor}>
          ワンタイムパスワード
        </Text>{' '}
        を入力
      </Text>
      <Text mb="1rem" textAlign="center" color={descriptionTextColor}>
        Authenticatorアプリで表示された6桁の数字を入力してください。
      </Text>
      <OtpForm onSubmit={onSubmit} />
    </Margin>
  );
};
