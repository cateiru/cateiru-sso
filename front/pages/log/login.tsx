import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import LoginLogPage from '../../components/log/LoginLogPage';

const LoginLog = () => {
  return (
    <Require isLogin={true} path="/">
      <Title title="ログインログ | CateiruSSO" />
      <LoginLogPage />
    </Require>
  );
};

export default LoginLog;
