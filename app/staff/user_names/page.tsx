import {StaffFrame} from '../../../components/Staff/StaffFrame';
import {UserNameTable} from '../../../components/Staff/User/UserNameTable';

const Page = () => {
  return (
    <StaffFrame
      title="予約されているユーザー名の一覧"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: '予約されているユーザー名の一覧'},
      ]}
    >
      <UserNameTable />
    </StaffFrame>
  );
};

export default Page;
