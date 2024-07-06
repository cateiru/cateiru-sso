import {ClientDetail} from '../../../../components/Staff/Client/ClientDetail';
import {StaffFrame} from '../../../../components/Staff/StaffFrame';

interface Props {
  params: {
    id: string;
  };
}

const Page: React.FC<Props> = ({params}) => {
  return (
    <StaffFrame
      title="クライアント詳細"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {href: '/staff/clients', pageName: 'クライアント一覧'},
        {pageName: 'クライアント詳細'},
      ]}
    >
      <ClientDetail id={params.id} />
    </StaffFrame>
  );
};

export default Page;
