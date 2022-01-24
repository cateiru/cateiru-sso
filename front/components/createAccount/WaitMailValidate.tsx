import {Text, Heading, Spinner, Button} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue, useRecoilState} from 'recoil';
import useVerifySurveillance from '../../hooks/useVerifySurveillance';
import {CTState, CreateNextState} from '../../utils/state/atom';

const WaitMail = React.memo<{mail: string}>(({mail}) => {
  const [surveillance, receive, close] = useVerifySurveillance();
  const ct = useRecoilValue(CTState);
  const [next, setNext] = useRecoilState(CreateNextState);

  React.useEffect(() => {
    let unmounted = false;
    if (ct.length !== 0 && !unmounted && !next) {
      surveillance(ct);
    }

    return () => {
      unmounted = true;
    };
  }, [ct]);

  return (
    <>
      {close ? (
        <>
          <Heading>Oops! サーバとの接続が切れてしまいました</Heading>
          <Text mt="1rem">
            メールアドレスに送られたリンクから続きをしてください
          </Text>
        </>
      ) : (
        <>
          <Heading>{mail} に確認メールを送信しました</Heading>
          <Text mt="1rem">
            メールアドレスが確認されるとこのページで続きをすることができます
          </Text>
          {receive ? (
            <Button
              colorScheme="blue"
              mt="2rem"
              onClick={() => {
                setNext(true);
              }}
            >
              このページで続きをする
            </Button>
          ) : (
            <Spinner
              mt="2rem"
              thickness="4px"
              speed="0.65s"
              color="blue.500"
              size="xl"
            />
          )}
        </>
      )}
    </>
  );
});

WaitMail.displayName = 'WaitMail';

export default WaitMail;
