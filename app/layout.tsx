import React from 'react';
import {PageEvents} from './PageEvents';
import {Providers} from './Providers';

const Layout: React.FC<{children: React.ReactNode}> = ({children}) => {
  return (
    <html lang="ja">
      <body>
        <PageEvents />
        <Providers>{children}</Providers>
      </body>
    </html>
  );
};

export default Layout;
