import {Box, Center} from '@chakra-ui/react';
import {useSteps} from 'chakra-ui-steps';
import React from 'react';
import {EmailInputPage} from './EmailInputPage';
import {EmailVerifyPage} from './EmailVerifyPage';
import {type StepStatus, Steps} from './Steps';

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
  const [registerToken, setRegisterToken] = React.useState<string | null>(null);
  const {
    activeStep,
    nextStep,
    prevStep,
    reset: resetStep,
  } = useSteps({
    initialStep: RegisterAccountStep.EmailInput,
  });
  const [status, setStatus] = React.useState<StepStatus>(undefined);

  const reset = React.useCallback(() => {
    setRegisterToken(null);
    setStatus(undefined);

    resetStep();
  }, []);

  const C = () => {
    switch (activeStep) {
      case RegisterAccountStep.EmailInput:
        return (
          <EmailInputPage
            nextStep={nextStep}
            prevStep={prevStep}
            setStatus={setStatus}
            reset={reset}
            setRegisterToken={setRegisterToken}
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
        return <></>;
      case RegisterAccountStep.CompleteRegister:
        if (!registerToken) {
          reset();
          return null;
        }
        return <></>;
    }
    return <></>;
  };

  return (
    <Box h="90vh">
      <Steps activeStep={activeStep} state={status} />
      <Center mt="1rem" h="50%">
        {C()}
      </Center>
    </Box>
  );
};
