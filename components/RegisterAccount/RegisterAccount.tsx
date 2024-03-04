import {Box, Center, Text} from '@chakra-ui/react';
import {useSteps} from 'chakra-ui-steps';
import React from 'react';
import {Spinner} from '../Common/Icons/Spinner';
import {useSecondaryColor} from '../Common/useColor';
import {CompleteRegisterPage} from './CompleteRegisterPage';
import {EmailInputPage} from './EmailInputPage';
import {EmailVerifyPage} from './EmailVerifyPage';
import {RegisterCertificatePage} from './RegisterCertificatePage';
import {type StepStatus, Steps} from './Steps';
import {useInvite} from './useInvite';

export enum RegisterAccountStep {
  EmailInput = 0,
  EmailVerify = 1,
  RegisterCertificate = 2,
  CompleteRegister = 3,
}

export interface DefaultPageProps {
  nextStep: () => void;
  prevStep: () => void;
  setStatus: (status: StepStatus) => void;
  reset: () => void;
}

export const RegisterAccount = () => {
  const textColor = useSecondaryColor();

  const [registerToken, setRegisterToken] = React.useState<string | null>(null);
  const {
    activeStep,
    nextStep,
    prevStep,
    setStep,
    reset: resetStep,
  } = useSteps({
    initialStep: RegisterAccountStep.EmailInput,
  });
  const [status, setStatus] = React.useState<StepStatus>(undefined);
  const [email, setEmail] = React.useState<string | null>(null);

  const {isInvite} = useInvite((email: string, token: string) => {
    setEmail(email);
    setRegisterToken(token);
    setStep(RegisterAccountStep.RegisterCertificate);
  });

  const reset = React.useCallback(() => {
    setRegisterToken(null);
    setStatus(undefined);
    setEmail(null);

    resetStep();
  }, []);

  const C = React.useCallback(() => {
    switch (activeStep) {
      case RegisterAccountStep.EmailInput:
        // invite_tokenを使用した場合はスピナーを表示させる
        if (isInvite) {
          return (
            <Center>
              <Spinner />
            </Center>
          );
        }

        return (
          <EmailInputPage
            nextStep={nextStep}
            prevStep={prevStep}
            setStatus={setStatus}
            reset={reset}
            setRegisterToken={setRegisterToken}
            setEmail={setEmail}
          />
        );
      case RegisterAccountStep.EmailVerify:
        if (!registerToken) {
          reset();
          return null;
        }
        return (
          <EmailVerifyPage
            nextStep={nextStep}
            prevStep={prevStep}
            setStatus={setStatus}
            reset={reset}
            registerToken={registerToken}
          />
        );
      case RegisterAccountStep.RegisterCertificate:
        if (!registerToken) {
          reset();
          return null;
        }
        return (
          <RegisterCertificatePage
            nextStep={nextStep}
            prevStep={prevStep}
            setStatus={setStatus}
            reset={reset}
            registerToken={registerToken}
          />
        );
      case RegisterAccountStep.CompleteRegister:
        if (!registerToken) {
          reset();
          return null;
        }
        return <CompleteRegisterPage />;
    }
    return <></>;
  }, [activeStep]);

  return (
    <Box>
      {email && (
        <Text textAlign="center" mb="1rem" color={textColor}>
          {email}
        </Text>
      )}
      <Steps activeStep={activeStep} state={status} />
      <Center mt="1rem" h="50%">
        <C />
      </Center>
    </Box>
  );
};
