import {OrgsTable} from '../../../components/Staff/Org/OrgsTable';
import {StaffFrame} from '../../../components/Staff/StaffFrame';

const Page = () => {
  return (
    <StaffFrame
      title="組織一覧"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: '組織一覧'},
      ]}
    >
      <OrgsTable />
    </StaffFrame>
  );
};

export default Page;
