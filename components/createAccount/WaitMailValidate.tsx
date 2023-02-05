import {Text, Heading, Button} from '@chakra-ui/react';
import React from 'react';
import {useSetRecoilState} from 'recoil';
import {CreateNextState} from '../../utils/state/atom';
import Spinner from '../common/Spinner';

const WaitMail = React.memo<{mail: string; receive: boolean; close: boolean}>(
  ({mail, receive, close}) => {
    const setNext = useSetRecoilState(CreateNextState);

    return (
      <>
        {close && !receive ? (
          <>
            <Heading>Oops! サーバとの接続が切れてしまいました</Heading>
            <Text mt="1rem">
              メールアドレスに送られたリンクから続きをしてください
            </Text>
          </>
        ) : (
          <>
            <Heading
              fontSize={{base: '1.2rem', md: '1.5rem'}}
              textAlign="center"
            >
              {mail} に確認メールを送信しました
            </Heading>
            <Text mt="1rem" textAlign="center">
              メールアドレスが確認されるとこのページで続きをすることができます。
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
              <Spinner />
            )}
          </>
        )}
      </>
    );
  }
);

WaitMail.displayName = 'WaitMail';

export default WaitMail;
