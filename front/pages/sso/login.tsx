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

const SSOLogin: NextPage<{
  oidc: OIDCRequestQuery;
  require: boolean;
}> = ({oidc, require}) => {
  const router = useRouter();
  const get = useGetUserInfo();
  const noLogin = useRecoilValue(NoLoginState);

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

export const getServerSideProps: GetServerSideProps<{
  oidc: OIDCRequestQuery;
  require: boolean;
}> = async context => {
  const oidc = new Login(context.query);

  return {
    props: {
      oidc: oidc.parse(),
      require: oidc.require(),
    },
  };
};

export default SSOLogin;
