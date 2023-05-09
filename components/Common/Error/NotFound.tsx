import {Box, Center, Heading} from '@chakra-ui/react';
import {KeyGen} from '../Animation/KeyGen';

export const NotFound = () => {
  return (
    <Center h="80vh">
      <Box>
        <Heading mb="1rem" textAlign="center">
          404 Not Found
        </Heading>
        <KeyGen />
      </Box>
    </Center>
  );
};
