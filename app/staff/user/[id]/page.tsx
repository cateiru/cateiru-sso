import React from 'react';
import {StaffFrame} from '../../../../components/Staff/StaffFrame';
import {UserDetail} from '../../../../components/Staff/User/UserDetail';

interface Props {
  params: {
    id: string;
  };
}

const Page: React.FC<Props> = ({params}) => {
  return (
    <StaffFrame
      title="ユーザー詳細"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {href: '/staff/users', pageName: 'ユーザー一覧'},
        {pageName: 'ユーザー詳細'},
      ]}
    >
      <UserDetail id={params.id} />
    </StaffFrame>
  );
};

export default Page;
