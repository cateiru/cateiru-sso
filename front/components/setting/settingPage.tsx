import {Box, Button} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import useDelete from '../../hooks/useDelete';
import useLogout from '../../hooks/useLogout';

const SettingPage = () => {
  const logout = useLogout();
  const deleteAccount = useDelete();
  const router = useRouter();

  const logoutHandle = () => {
    const f = async () => {
      await logout();
      router.replace('/');
    };

    f();
  };

  const deleteHandle = () => {
    const f = async () => {
      await deleteAccount();
      router.replace('/');
    };

    f();
  };

  return (
    <Box>
      <Button colorScheme="blue" onClick={logoutHandle}>
        ログアウト
      </Button>
      <Button onClick={deleteHandle}>アカウント削除</Button>
    </Box>
  );
};

export default SettingPage;
