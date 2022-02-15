import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import ConnectedAccountPage from '../../components/connectedAccount/ConnectedAccountPage';

const ConnectedAccount = () => {
  return (
    <Require isLogin={true} path="/">
      <Title title="連携されているアカウント | CateiruSSO" />
      <ConnectedAccountPage />
    </Require>
  );
};

export default ConnectedAccount;
