import {Staff} from '../../components/Staff/Staff';
import {StaffFrame} from '../../components/Staff/StaffFrame';

const Page = () => {
  return (
    <StaffFrame
      title="スタッフ管理画面"
      paths={[{pageName: 'スタッフ管理画面'}]}
    >
      <Staff />
    </StaffFrame>
  );
};

export default Page;
