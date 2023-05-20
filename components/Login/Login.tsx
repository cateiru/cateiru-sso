'use client';

import {Box, Center, Heading} from '@chakra-ui/react';
import React from 'react';
import {LoginUser} from '../../utils/types/login';
import {LoginSuccess} from './LoginSuccess';
import {OtpPage} from './OtpPage';
import {UserIDEmailPage} from './UserIDEmailPage';

export enum LoginStep {
  UserIDEmail,
  OTP,
  CompleteLogin,
}

export interface DefaultPageProps {
  setStep: (step: LoginStep) => void;
  reset: () => void;
}

export const Login = () => {
  const [step, setStep] = React.useState<LoginStep>(LoginStep.UserIDEmail);
  const [loginUser, setLoginUser] = React.useState<LoginUser | null>(null);
  const [otpToken, setOtpToken] = React.useState('');

  const reset = () => setStep(LoginStep.UserIDEmail);

  const C = React.useCallback(() => {
    switch (step) {
      case LoginStep.UserIDEmail:
        return (
          <UserIDEmailPage
            setStep={setStep}
            setLoginUser={setLoginUser}
            reset={reset}
            setOTPToken={setOtpToken}
          />
        );
      case LoginStep.OTP:
        return (
          <OtpPage
            setStep={setStep}
            reset={reset}
            loginUser={loginUser}
            otpToken={otpToken}
          />
        );
      case LoginStep.CompleteLogin:
        return <LoginSuccess />;
    }
  }, [step]);

  return (
    <Box minH="80vh">
      <Heading textAlign="center" mt="3rem">
        ログイン
      </Heading>
      <Center mt="1rem" h="50%">
        <C />
      </Center>
    </Box>
  );
};
