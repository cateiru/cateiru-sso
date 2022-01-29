import {Button} from '@chakra-ui/react';
import {logout, deleteAccount} from '../../utils/api/logout';

const AccountSetting = () => {
  const logoutHandle = () => {
    const f = async () => {
      await logout();
    };
    f();
  };

  const deleteHandle = () => {
    const f = async () => {
      await deleteAccount();
    };
    f();
  };

  return (
    <>
      工事中:
      <Button colorScheme="blue" onClick={logoutHandle}>
        ログアウト
      </Button>
      <Button onClick={deleteHandle}>アカウント削除</Button>
    </>
  );
};

export default AccountSetting;
