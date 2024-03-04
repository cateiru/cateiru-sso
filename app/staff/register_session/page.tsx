import {RegisterSessionTable} from '../../../components/Staff/Session/RegisterSessionTable';
import {StaffFrame} from '../../../components/Staff/StaffFrame';

const Page = () => {
  return (
    <StaffFrame
      title="登録セッション一覧"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: '登録セッション一覧'},
      ]}
    >
      <RegisterSessionTable />
    </StaffFrame>
  );
};

export default Page;
