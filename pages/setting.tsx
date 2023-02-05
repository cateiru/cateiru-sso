import Require from '../components/common/Require';
import Title from '../components/common/Title';
import SettingPage from '../components/setting/SettingPage';
import UserSetting from '../components/setting/UserSetting';

const Setting = () => {
  return (
    <Require isLogin={true} path="/">
      <Title title="設定 | CateiruSSO" />
      <SettingPage index={0}>
        <UserSetting />
      </SettingPage>
    </Require>
  );
};

export default Setting;
