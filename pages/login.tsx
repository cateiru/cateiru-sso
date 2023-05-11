import {Session} from '../components/Common/Session';
import {Login} from '../components/Login/Login';

const LoginPage = () => {
  return (
    <Session isLoggedIn>
      <Login />
    </Session>
  );
};

export default LoginPage;
