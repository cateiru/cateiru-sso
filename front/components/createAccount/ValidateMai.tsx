import {Heading, Button} from '@chakra-ui/react';
import React from 'react';
import {useSetRecoilState} from 'recoil';
import {CreateNextState} from '../../utils/state/atom';
import Spinner from '../common/Spinner';

const ValidateMail: React.FC<{
  isKeep: boolean;
  loadVerify: boolean;
  verifyError: boolean;
}> = ({isKeep, loadVerify, verifyError}) => {
  const setNext = useSetRecoilState(CreateNextState);

  return (
    <>
      {verifyError ? (
        <Heading>Oops!</Heading>
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
