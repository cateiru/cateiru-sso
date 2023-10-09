import {Box, Center, Heading, Text} from '@chakra-ui/react';
import React from 'react';
import {TbMoodSadSquint} from 'react-icons/tb';
import {
  ErrorType,
  ErrorUniqueMessage,
  OidcErrorType,
} from '../../../utils/types/error';
import {useSecondaryColor} from '../useColor';

export const Error: React.FC<ErrorType> = props => {
  return (
    <Center h="80vh">
      <Box>
        <Center color="red.500">
          <TbMoodSadSquint size="50px" />
        </Center>
        <Heading color="red.500" textAlign="center">
          {props.unique_code
            ? ErrorUniqueMessage[props.unique_code] ?? props.message
            : props.message}
        </Heading>
      </Box>
    </Center>
  );
};

export const OidcError: React.FC<OidcErrorType> = props => {
  const secondaryColor = useSecondaryColor();

  return (
    <Center h="80vh">
      <Box>
        <Center color="red.500">
          <TbMoodSadSquint size="50px" />
        </Center>
        <Heading color="red.500" textAlign="center">
          {props.error}
        </Heading>
        <Text color={secondaryColor} textAlign="center" mt=".5rem">
          {props.error_description}
        </Text>
      </Box>
    </Center>
  );
};
