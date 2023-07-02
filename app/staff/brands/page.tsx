import {BrandsTable} from '../../../components/Staff/Brand/BrandsTable';
import {StaffFrame} from '../../../components/Staff/StaffFrame';

const Page = () => {
  return (
    <StaffFrame
      title="ブランド一覧"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: 'ブランド一覧'},
      ]}
    >
      <BrandsTable />
    </StaffFrame>
  );
};

export default Page;
