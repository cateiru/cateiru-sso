import Require from '../components/common/Require';
import Title from '../components/common/Title';
import SSOPage from '../components/sso/SSOPage';

const SSO = () => {
  return (
    <Require isLogin={true} path="/" role="pro">
      <Title title="ダッシュボード | CateiruSSO" />
      <SSOPage />
    </Require>
  );
};

export default SSO;
