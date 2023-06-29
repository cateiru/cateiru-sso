import React from 'react';
import {StaffFrame} from '../../../components/Staff/StaffFrame';

const Layout: React.FC<{children: React.ReactNode}> = ({children}) => {
  return (
    <StaffFrame
      title="デプロイデータ"
      paths={[
        {href: '/staff', pageName: 'スタッフ管理画面'},
        {pageName: 'デプロイデータ'},
      ]}
    >
      {children}
    </StaffFrame>
  );
};

export default Layout;
