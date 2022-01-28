import Require from '../components/common/Require';
import Title from '../components/common/Title';
import LoginPage from '../components/login/LoginPage';

const Login = () => {
  return (
    <Require isLogin={false} path="/hello">
      <Title title="ログイン | CateiruSSO" />
      <LoginPage />
    </Require>
  );
};

export default Login;
