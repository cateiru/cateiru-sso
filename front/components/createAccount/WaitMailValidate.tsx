import {Text, Heading, Spinner} from '@chakra-ui/react';
import React from 'react';

const WaitMail: React.FC<{mail: string}> = ({mail}) => {
  return (
    <>
      <Heading>{mail} に確認メールを送信しました</Heading>
      <Text mt="1rem">
        メールアドレスが確認されるとこのページで続きをすることができます
      </Text>
      <Spinner
        mt="2rem"
        thickness="4px"
        speed="0.65s"
        color="blue.500"
        size="xl"
      />
    </>
  );
};

export default WaitMail;
