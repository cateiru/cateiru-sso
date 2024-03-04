import {RegisterBrand} from '../../../../components/Staff/Brand/RegisterBrand';
import {StaffFrame} from '../../../../components/Staff/StaffFrame';

const Page = () => {
  return (
    <StaffFrame
      title="ブランド新規作成"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {href: '/staff/brands', pageName: 'ブランド一覧'},
        {pageName: 'ブランド新規作成'},
      ]}
    >
      <RegisterBrand />
    </StaffFrame>
  );
};

export default Page;
