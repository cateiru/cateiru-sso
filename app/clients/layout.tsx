import React from 'react';
import {ClientsListWrapper} from '../../components/Client/ClientsListWrapper';

interface Props {
  children: React.ReactNode;
}

const Layout: React.FC<Props> = ({children}) => {
  return <ClientsListWrapper>{children}</ClientsListWrapper>;
};

export default Layout;
