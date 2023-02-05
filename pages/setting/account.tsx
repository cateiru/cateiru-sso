import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import AccountSetting from '../../components/setting/AccountSetting';
import SettingPage from '../../components/setting/SettingPage';

const Account = () => {
  return (
    <Require isLogin={true} path="/">
      <Title title="設定 | CateiruSSO" />
      <SettingPage index={3}>
        <AccountSetting />
      </SettingPage>
    </Require>
  );
};

export default Account;
