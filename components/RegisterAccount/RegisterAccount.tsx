import {Box, Center} from '@chakra-ui/react';
import {useSteps} from 'chakra-ui-steps';
import React from 'react';
import {EmailInputPage} from './EmailInputPage';
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
}

export const RegisterAccount = () => {
  const [registerToken, setRegisterToken] = React.useState<string | null>(null);
  const {activeStep, nextStep, prevStep} = useSteps({
    initialStep: RegisterAccountStep.EmailInput,
  });
  const [status, setStatus] = React.useState<StepStatus>(undefined);

  const C = () => {
    switch (activeStep) {
      case RegisterAccountStep.EmailInput:
        return (
          <EmailInputPage
            nextStep={nextStep}
            prevStep={prevStep}
            setStatus={setStatus}
            setRegisterToken={setRegisterToken}
          />
        );
      case RegisterAccountStep.EmailVerify:
        return <></>;
      case RegisterAccountStep.RegisterCertificate:
        return <></>;
      case RegisterAccountStep.CompleteRegister:
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
