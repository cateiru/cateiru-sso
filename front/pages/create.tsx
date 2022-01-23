import NoSSR from 'react-no-ssr';
import Title from '../components/common/Title';
import CreateAccountPage from '../components/createAccount/CreateAccountPage';

const CreateAccount = () => {
  return (
    <>
      <Title title="アカウント作成 | CateiruSSO" />
      <NoSSR>
        <CreateAccountPage />
      </NoSSR>
    </>
  );
};

export default CreateAccount;
