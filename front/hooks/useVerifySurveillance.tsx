import {useToast} from '@chakra-ui/react';
import React from 'react';
import Socket from '../utils/api/socket/socket';

const useVerifySurveillance = (): [
  (token: string) => void,
  boolean,
  boolean
] => {
  const [receive, setReceive] = React.useState(false);
  const toast = useToast();
  const [close, setClose] = React.useState(false);

  const ref = React.useRef<Socket>();

  const surveillance = (token: string) => {
    const socket = new Socket(`/create/verify?cct=${token}`);

    socket.error(() => {
      setClose(true);
      toast({
        title: 'エラーによりコネクションが切断されました',
        status: 'error',
        isClosable: true,
      });
    });

    socket.end(() => {
      setClose(true);
      toast({
        title: 'コネクションが終了しました',
        status: 'info',
        isClosable: true,
      });
    });

    socket.get(data => {
      if (data === 'true') {
        setReceive(true);
      }
    });

    ref.current = socket;
  };

  // receiveが取得できたらcloseする
  React.useEffect(() => {
    if (receive && ref.current) {
      ref.current.close();
    }
  }, [receive]);

  return [surveillance, receive, close];
};

export default useVerifySurveillance;
