import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import MailSetting from '../../components/setting/MailSetting';
import SettingPage from '../../components/setting/SettingPage';

const MailSettingPage = () => {
  return (
    <Require isLogin={true} path="/hello">
      <Title title="設定 | CateiruSSO" />
      <SettingPage index={1}>
        <MailSetting />
      </SettingPage>
    </Require>
  );
};

export default MailSettingPage;
