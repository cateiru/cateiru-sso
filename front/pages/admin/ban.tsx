import BanPage from '../../components/admin/BanPage';
import Require from '../../components/common/Require';
import Title from '../../components/common/Title';

const CreateLog = () => {
  return (
    <Require isLogin={true} path="/" role="admin">
      <Title title="Banリスト | CateiruSSO" />
      <BanPage />
    </Require>
  );
};

export default CreateLog;
