import Require from '../components/common/Require';
import Title from '../components/common/Title';
import SettingPage from '../components/setting/settingPage';

const Setting = () => {
  return (
    <Require isLogin={true} path="/">
      <Title title="設定 | CateiruSSO" />
      <SettingPage />
    </Require>
  );
};

export default Setting;
