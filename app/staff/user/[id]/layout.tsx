import React from 'react';
import {StaffFrame} from '../../../../components/Staff/StaffFrame';

const Layout: React.FC<{children: React.ReactNode}> = ({children}) => {
  return (
    <StaffFrame
      title="ユーザー詳細"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {href: '/staff/users', pageName: 'ユーザー一覧'},
        {pageName: 'ユーザー詳細'},
      ]}
    >
      {children}
    </StaffFrame>
  );
};

export default Layout;
