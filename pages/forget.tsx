import Require from '../components/common/Require';
import Title from '../components/common/Title';
import ForgetPage from '../components/forget/ForgetPage';

const ForgetPassword = () => {
  return (
    <Require isLogin={false} path="/hello">
      <Title title="パスワード再設定 | CateiruSSO" />
      <ForgetPage />
    </Require>
  );
};

export default ForgetPassword;
