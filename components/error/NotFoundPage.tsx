import {Text, Flex} from '@chakra-ui/react';
import React from 'react';
import KeyAA from './KeyAA';

const NotFoundPage = () => {
  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <Text fontSize="2.5rem" fontWeight="light" marginBottom="1.5rem">
        404 | Not Found.
      </Text>
      <KeyAA />
    </Flex>
  );
};

export default NotFoundPage;
