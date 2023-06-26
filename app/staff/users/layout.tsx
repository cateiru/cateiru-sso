import React from 'react';
import {StaffFrame} from '../../../components/Staff/StaffFrame';

const Layout: React.FC<{children: React.ReactNode}> = ({children}) => {
  return (
    <StaffFrame
      title="ユーザー一覧"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: 'ユーザー一覧'},
      ]}
    >
      {children}
    </StaffFrame>
  );
};

export default Layout;
