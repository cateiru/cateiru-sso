import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import LoginLogPage from '../../components/log/LoginLogPage';

const LoginLog = () => {
  return (
    <Require isLogin={true} path="/">
      <Title title="ログイン履歴 | CateiruSSO" />
      <LoginLogPage />
    </Require>
  );
};

export default LoginLog;
