import {Spinner as ChakraSpinner, useColorMode} from '@chakra-ui/react';

const Spinner = () => {
  const {colorMode} = useColorMode();

  return (
    <ChakraSpinner
      mt="2rem"
      thickness="5px"
      speed=".6s"
      color={colorMode === 'dark' ? 'blue.200' : 'blue.700'}
      size="xl"
    />
  );
};

export default Spinner;
