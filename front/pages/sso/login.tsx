import {GetServerSideProps, NextPage} from 'next';
import {useRouter} from 'next/router';
import React from 'react';
import {useRecoilValue} from 'recoil';
import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import LoginPage from '../../components/sso/login/LoginPage';
import {useGetUserInfo} from '../../hooks/useGetUserInfo';
import cookieValue from '../../utils/cookie';
import {Login, OIDCRequestQuery} from '../../utils/sso/login';
import {UserState} from '../../utils/state/atom';

const SSOLogin: NextPage<{
  oidc: OIDCRequestQuery;
  require: boolean;
}> = ({oidc, require}) => {
  const router = useRouter();
  const get = useGetUserInfo();
  const user = useRecoilValue(UserState);

  React.useEffect(() => {
    if (typeof user === 'undefined' && cookieValue('refresh-token')) {
      get();
    }
  }, []);

  return (
    <Require
      isLogin={true}
      path={`/login?redirect=${encodeURIComponent(router.asPath)}`}
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
