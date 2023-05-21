import React from 'react';
import {Session} from '../../components/Common/Session';
import {SettingHeader} from '../../components/Settings/SettingHeader';

const Layout: React.FC<{children: React.ReactNode}> = ({children}) => {
  return (
    <Session>
      <SettingHeader>{children}</SettingHeader>
    </Session>
  );
};

export default Layout;
