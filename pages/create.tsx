import Require from '../components/common/Require';
import Title from '../components/common/Title';
import CreateAccountPage from '../components/createAccount/CreateAccountPage';

const CreateAccount = () => {
  return (
    <Require isLogin={false} path="/hello">
      <Title title="アカウント作成 | CateiruSSO" />
      <CreateAccountPage />
    </Require>
  );
};

export default CreateAccount;
