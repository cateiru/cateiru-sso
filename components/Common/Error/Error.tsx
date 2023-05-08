import {Center, Heading} from '@chakra-ui/react';
import React from 'react';
import {ErrorType, ErrorUniqueMessage} from '../../../utils/types/error';

export const Error: React.FC<ErrorType> = props => {
  return (
    <Center>
      <Heading color="red.500">
        {props.unique_code
          ? ErrorUniqueMessage[props.unique_code] ?? props.message
          : props.message}
      </Heading>
    </Center>
  );
};
