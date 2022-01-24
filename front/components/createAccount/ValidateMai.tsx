import {Heading, Button, Spinner} from '@chakra-ui/react';
import React from 'react';
import useVerify from '../../hooks/useVerify';

const ValidateMail: React.FC<{
  token: string;
  setCT: (t: string) => void;
  next: () => void;
}> = ({token, setCT, next}) => {
  const [verify, isKeep, ct, error] = useVerify();
  const [load, setLoad] = React.useState(true);

  React.useEffect(() => {
    console.log(token);
  }, [token]);

  // API叩く
  React.useEffect(() => {
    if (token.length !== 0) {
      verify(token);
    }
  }, []);

  React.useEffect(() => {
    if (ct.length !== 0) {
      setLoad(false);

      // CTをセットする
      setCT(ct);

      // isKeepがtrueの場合は強制的に次へ進む
      if (isKeep) {
        next();
      }
    }
  }, [ct]);

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
                <Button colorScheme="blue" mt="2rem" onClick={next}>
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
