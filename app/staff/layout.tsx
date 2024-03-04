import React from 'react';
import {Session} from '../../components/Common/Session';

const Layout: React.FC<{children: React.ReactNode}> = ({children}) => {
  return <Session isStaff>{children}</Session>;
};

export default Layout;
