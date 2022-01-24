import {Text, Heading, Spinner, Button} from '@chakra-ui/react';
import React from 'react';
import useVerifySurveillance from '../../hooks/useVerifySurveillance';

const WaitMail = React.memo<{mail: string; token: string; next: () => void}>(
  ({mail, token, next}) => {
    const [surveillance, receive, close] = useVerifySurveillance();

    React.useEffect(() => {
      let unmounted = false;
      if (token && !unmounted) {
        surveillance(token);
      }

      return () => {
        unmounted = true;
      };
    }, []);

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
              <Button colorScheme="blue" mt="2rem" onClick={next}>
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
  }
);

WaitMail.displayName = 'WaitMail';

export default WaitMail;
