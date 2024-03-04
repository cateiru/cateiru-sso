import {RegisterOrg} from '../../../../components/Staff/Org/RegisterOrg';
import {StaffFrame} from '../../../../components/Staff/StaffFrame';

const Page = () => {
  return (
    <StaffFrame
      title="組織新規作成"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {href: '/staff/orgs', pageName: '組織一覧'},
        {pageName: '組織新規作成'},
      ]}
    >
      <RegisterOrg />
    </StaffFrame>
  );
};

export default Page;
