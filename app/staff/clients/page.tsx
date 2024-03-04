import {ClientsTable} from '../../../components/Staff/Client/ClientsTable';
import {StaffFrame} from '../../../components/Staff/StaffFrame';

const Page = () => {
  return (
    <StaffFrame
      title="クライアント一覧"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: 'クライアント一覧'},
      ]}
    >
      <ClientsTable />
    </StaffFrame>
  );
};

export default Page;
