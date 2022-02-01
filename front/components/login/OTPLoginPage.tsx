import {Flex} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useSetRecoilState} from 'recoil';
import {LoadState} from '../../utils/state/atom';
import Spinner from '../common/Spinner';
import OTPLoginForm from './OTPLoginForm';

const OTPLoginPage = () => {
  const [token, setToken] = React.useState('');
  const router = useRouter();
  const setLoad = useSetRecoilState(LoadState);

  // クエリパラメータからトークンを取得する
  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['t'] === 'string') {
      setToken(query['t']);
      setLoad(false);
    }
  }, [router.isReady, router.query]);

  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      {token ? <OTPLoginForm token={token} /> : <Spinner />}
    </Flex>
  );
};

export default OTPLoginPage;
