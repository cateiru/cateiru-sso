import {Avatar, Button, Center, Text} from '@chakra-ui/react';
import React from 'react';
import type {LoginUser} from '../../utils/types/login';
import {Margin} from '../Common/Margin';
import {type DefaultPageProps, LoginStep} from './Login';

interface Props extends DefaultPageProps {
  loginUser: LoginUser;
}

export const SelectCertification: React.FC<Props> = props => {
  const onChange = (type: 'password' | 'passkey') => {
    switch (type) {
      case 'password':
        props.setStep(LoginStep.Password);
        return;
      case 'passkey':
        props.setStep(LoginStep.WebAuthn);
        return;
    }
  };

  return (
    <Margin>
      <Center my="1rem">
        <Avatar src={props.loginUser.avatar ?? ''} size="lg" />
      </Center>
      <Text fontSize="1.5rem" fontWeight="bold" mb="1rem" textAlign="center">
        ログインに使用する認証方式を選択してください
      </Text>
      <Button
        w="100%"
        mt="1rem"
        colorScheme="cateiru"
        onClick={() => onChange('passkey')}
      >
        生体認証
      </Button>
      <Button
        w="100%"
        mt="1rem"
        colorScheme="gray"
        onClick={() => onChange('password')}
      >
        パスワード
      </Button>
    </Margin>
  );
};
