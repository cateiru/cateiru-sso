import {Heading, Button, Spinner} from '@chakra-ui/react';
import React from 'react';
import {useSetRecoilState, useRecoilState} from 'recoil';
import useVerify from '../../hooks/useVerify';
import {CTState, CreateNextState} from '../../utils/state/atom';

const ValidateMail: React.FC<{
  token: string;
}> = ({token}) => {
  const [verify, isKeep, ct, error] = useVerify();
  const [load, setLoad] = React.useState(true);
  const setCT = useSetRecoilState(CTState);
  const [next, setNext] = useRecoilState(CreateNextState);

  React.useEffect(() => {
    console.log(token);
  }, [token]);

  // API叩く
  React.useEffect(() => {
    if (token.length !== 0 && !next) {
      verify(token);
    }
  }, []);

  React.useEffect(() => {
    if (ct.length !== 0) {
      setLoad(false);

      // CTをセットする
      setCT(ct);
    }

    // isKeepがtrueの場合は強制的に次へ進む
    if (isKeep) {
      setNext(true);
    }
  }, [ct, isKeep]);

  return (
    <>
      {error ? (
        <Heading>Oops!</Heading>
      ) : (
        <>
          {load ? (
            <Spinner thickness="4px" speed="0.65s" color="blue.500" size="xl" />
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
