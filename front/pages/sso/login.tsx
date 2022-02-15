import {GetServerSideProps, NextPage} from 'next';
import React from 'react';
import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import LoginPage from '../../components/sso/login/LoginPage';
import {Login, OIDCRequestQuery} from '../../utils/sso/login';

const SSOLogin: NextPage<{
  oidc: OIDCRequestQuery;
  require: boolean;
}> = ({oidc, require}) => {
  return (
    <Require isLogin={true} path="/hello">
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
