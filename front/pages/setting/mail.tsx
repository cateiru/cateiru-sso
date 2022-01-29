import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import SettingPage from '../../components/setting/SettingPage';

const MailSetting = () => {
  return (
    <Require isLogin={true} path="/">
      <Title title="設定 | CateiruSSO" />
      <SettingPage index={1} />
    </Require>
  );
};

export default MailSetting;
