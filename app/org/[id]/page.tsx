import {Session} from '../../../components/Common/Session';
import {OrganizationDetail} from '../../../components/Organization/OrganizationDetail';

interface Props {
  params: {
    id: string;
  };
}

const Page: React.FC<Props> = ({params}) => {
  return (
    <Session>
      <OrganizationDetail id={params.id} />
    </Session>
  );
};

export default Page;
