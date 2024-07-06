import {Spinner as ChakraSpinner, useColorModeValue} from '@chakra-ui/react';

export const Spinner: typeof ChakraSpinner = props => {
  const spinnerColor = useColorModeValue('my.primary', 'my.secondary');

  return <ChakraSpinner color={spinnerColor} {...props} />;
};
