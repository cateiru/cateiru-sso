import {Session} from '../../components/Common/Session';
import {LoginHistory} from '../../components/Histories/LoginHistory';

const Page = () => {
  return (
    <Session>
      <LoginHistory />
    </Session>
  );
};

export default Page;
