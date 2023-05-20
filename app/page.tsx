import {Session} from '../components/Common/Session';
import {Top} from '../components/Top/Top';

const Page = () => {
  return (
    <Session isLoggedIn redirectTo="/profile">
      <Top />
    </Session>
  );
};

export default Page;
