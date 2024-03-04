import React from 'react';
import {Session} from '../../components/Common/Session';
import {LoginHistoryHeader} from '../../components/Histories/LoginHistoryHeader';

const Layout: React.FC<{children: React.ReactNode}> = ({children}) => {
  return (
    <Session>
      <LoginHistoryHeader>{children}</LoginHistoryHeader>
    </Session>
  );
};

export default Layout;
