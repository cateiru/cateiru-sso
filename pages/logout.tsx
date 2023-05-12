import {Button, Center} from '@chakra-ui/react';
import {Session} from '../components/Common/Session';
import {useLogout} from '../components/Common/useLogout';

const Logout = () => {
  const {logout} = useLogout();

  return (
    <Session>
      <Center h="100vh">
        <Button onClick={logout}>ログアウト</Button>
      </Center>
    </Session>
  );
};

export default Logout;
