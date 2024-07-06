import React from 'react';
import {ClientsListWrapper} from '../../components/Client/ClientsListWrapper';
import {Session} from '../../components/Common/Session';

interface Props {
  children: React.ReactNode;
}

const Layout: React.FC<Props> = ({children}) => {
  return (
    <Session>
      <ClientsListWrapper>{children}</ClientsListWrapper>
    </Session>
  );
};

export default Layout;
