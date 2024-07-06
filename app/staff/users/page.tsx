import {StaffFrame} from '../../../components/Staff/StaffFrame';
import {UsersTable} from '../../../components/Staff/User/UsersTable';

const Page = () => {
  return (
    <StaffFrame
      title="ユーザー一覧"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: 'ユーザー一覧'},
      ]}
    >
      <UsersTable />
    </StaffFrame>
  );
};

export default Page;
