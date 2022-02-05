import {useRouter} from 'next/router';
import React from 'react';
import AcceptPage from './AcceptPage';
import Forget from './Forget';

const ForgetPage = () => {
  const router = useRouter();
  const [token, setToken] = React.useState('');

  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['t'] === 'string') {
      setToken(query['t']);
    }
  }, [router.isReady, router.query]);

  return <>{token ? <AcceptPage /> : <Forget />}</>;
};

export default ForgetPage;
