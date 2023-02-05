import Require from '../../components/common/Require';
import Title from '../../components/common/Title';
import PasswordSetting from '../../components/setting/PasswordSetting';
import SettingPage from '../../components/setting/SettingPage';

const Password = () => {
  return (
    <Require isLogin={true} path="/">
      <Title title="設定 | CateiruSSO" />
      <SettingPage index={2}>
        <PasswordSetting />
      </SettingPage>
    </Require>
  );
};

export default Password;
