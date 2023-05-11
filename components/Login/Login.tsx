import {Box, Center, Heading} from '@chakra-ui/react';
import React from 'react';
import {LoginUser} from '../../utils/types/login';
import {OtpPage} from './OtpPage';
import {PasswordPage} from './PasswordPage';
import {UserIDEmailPage} from './UserIDEmailPage';

export enum LoginStep {
  UserIDEmail,
  Password,
  OTP,
  WebAuthn,
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
          />
        );
      case LoginStep.Password:
        if (!loginUser) throw new Error('loginUser is null');
        return (
          <PasswordPage
            setStep={setStep}
            reset={reset}
            loginUser={loginUser}
            setOTPToken={setOtpToken}
          />
        );
      case LoginStep.OTP:
        if (!loginUser) throw new Error('loginUser is null');
        return (
          <OtpPage
            setStep={setStep}
            reset={reset}
            loginUser={loginUser}
            otpToken={otpToken}
          />
        );
      case LoginStep.WebAuthn:
        if (!loginUser) throw new Error('loginUser is null');
        return <></>;
      case LoginStep.CompleteLogin:
        if (!loginUser) throw new Error('loginUser is null');
        return <></>;
    }
  }, [step]);

  return (
    <Box h="80vh">
      <Heading textAlign="center" mt="3rem">
        ログイン
      </Heading>
      <Center mt="1rem" h="50%">
        <C />
      </Center>
    </Box>
  );
};
