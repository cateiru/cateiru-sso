import {OidcRequirePage} from '../../components/Auth/OidcRequirePage';
import {Session} from '../../components/Common/Session';

const Page = () => {
  return (
    <Session noRedirect>
      <OidcRequirePage />
    </Session>
  );
};

export default Page;
