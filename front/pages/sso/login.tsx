import {GetServerSideProps, NextPage} from 'next';
import {useRouter} from 'next/router';
import React from 'react';
import {useRecoilValue} from 'recoil';
import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import LoginPage from '../../components/sso/login/LoginPage';
import LoginSkeleton from '../../components/sso/login/LoginSkeleton';
import {useGetUserInfo} from '../../hooks/useGetUserInfo';
import {Login, OIDCRequestQuery} from '../../utils/sso/login';
import {NoLoginState} from '../../utils/state/atom';

const SSOLogin = () => {
  const router = useRouter();
  const get = useGetUserInfo();
  const noLogin = useRecoilValue(NoLoginState);
  const [oidc, setOIdc] = React.useState<OIDCRequestQuery>();
  const [require, setRequire] = React.useState(false);

  React.useEffect(() => {
    if (!router.isReady) return;
    const o = new Login(router.query);
    setOIdc(o.parse());
    setRequire(o.require());
  }, [router.isReady, router.query]);

  React.useEffect(() => {
    if (noLogin) {
      get();
    }
  }, []);

  return (
    <Require
      isLogin={true}
      path={`/login?redirect=${encodeURIComponent(router.asPath)}`}
      loading={<LoginSkeleton />}
    >
      <Title title="ログイン | CateiruSSO" />
      <LoginPage oidc={oidc} require={require} />
    </Require>
  );
};

export default SSOLogin;
