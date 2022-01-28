import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import OTPLoginPage from '../../components/login/OTPLoginPage';

const OTPLogin = () => {
  return (
    <Require isLogin={false} path="/hello">
      <Title title="ログイン | CateiruSSO" />
      <OTPLoginPage />
    </Require>
  );
};

export default OTPLogin;
