import {Heading, Button, Text} from '@chakra-ui/react';
import React from 'react';
import {useSetRecoilState} from 'recoil';
import {CreateNextState} from '../../utils/state/atom';
import Spinner from '../common/Spinner';

const ValidateMail: React.FC<{
  isKeep: boolean;
  loadVerify: boolean;
  verifyError: string;
}> = ({isKeep, loadVerify, verifyError}) => {
  const setNext = useSetRecoilState(CreateNextState);

  return (
    <>
      {verifyError ? (
        <>
          <Heading>不正なURLです</Heading>
          <Text>{verifyError}</Text>
        </>
      ) : (
        <>
          {loadVerify ? (
            <Spinner />
          ) : (
            <>
              <Heading>メールアドレスを確認しました</Heading>
              {isKeep || (
                <Button
                  colorScheme="blue"
                  mt="2rem"
                  onClick={() => {
                    setNext(true);
                  }}
                >
                  このページで続きをする
                </Button>
              )}
            </>
          )}
        </>
      )}
    </>
  );
};

export default ValidateMail;
