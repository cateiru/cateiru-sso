import {useToast} from '@chakra-ui/react';
import React from 'react';
import {LoginLogResponse, getLoginLog} from '../utils/api/log';

const useLoginLog = (): [
  LoginLogResponse[],
  boolean,
  (limit: number | undefined) => void
] => {
  const [log, setLog] = React.useState<LoginLogResponse[]>([]);
  const [load, setLoad] = React.useState(false);
  const toast = useToast();

  const getLog = (limit: number | undefined) => {
    const f = async () => {
      setLoad(true);
      try {
        const resp = await getLoginLog(limit);
        setLog(resp);
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }

      setLoad(false);
    };

    f();
  };

  return [log, load, getLog];
};

export default useLoginLog;
