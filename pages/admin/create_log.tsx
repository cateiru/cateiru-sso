import MailCertLog from '../../components/admin/MailCertLog';
import Require from '../../components/common/Require';
import Title from '../../components/common/Title';

const CreateLog = () => {
  return (
    <Require isLogin={true} path="/" role="admin">
      <Title title="メール認証ログ | CateiruSSO" />
      <MailCertLog />
    </Require>
  );
};

export default CreateLog;
