import {Heading, Button} from '@chakra-ui/react';
import React from 'react';

const ValidateMail: React.FC<{token: string}> = ({token}) => {
  return (
    <>
      <Heading>メールアドレスを確認しました</Heading>
      <Button colorScheme="blue" mt="2rem">
        このページで続きをする
      </Button>
    </>
  );
};

export default ValidateMail;
