import {Steps as ChakraSteps, Step as ChakraStep} from 'chakra-ui-steps';
import React from 'react';

export type StepStatus = 'loading' | 'error' | undefined;

interface Props {
  state?: StepStatus;
  activeStep: number;
}

export const Steps = React.memo<Props>(({activeStep, state}) => {
  return (
    <ChakraSteps
      activeStep={activeStep}
      colorScheme="cateiru"
      variant="circles-alt"
      state={state}
      responsive={false}
    >
      <ChakraStep label="Emailを入力" />
      <ChakraStep label="Emailを認証" />
      <ChakraStep label="パスワードの設定" />
    </ChakraSteps>
  );
});

Steps.displayName = 'Steps';
