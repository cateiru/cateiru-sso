import {Box, Center, Heading} from '@chakra-ui/react';
import {config} from '../../utils/config';
import {Logo} from '../Common/Icons/Logo';

export const Top = () => {
  return (
    <Center h="100vh">
      <Box>
        <Center>
          <Logo size="50%" />
        </Center>
        <Heading
          textAlign="center"
          background="linear-gradient(124deg, #2bc4cf, #572bcf, #cf2ba1)"
          backgroundClip="text"
          fontSize={{base: '2rem', md: '3rem'}}
        >
          {config.title}
        </Heading>
      </Box>
    </Center>
  );
};
