import React from 'react';
import {config} from '../utils/config';
import {PageEvents} from './PageEvents';
import {Providers} from './Providers';

const Layout: React.FC<{children: React.ReactNode}> = ({children}) => {
  return (
    <html
      lang="ja"
      data-revision={config.revision}
      data-branch={config.branchName ?? undefined}
    >
      <body>
        <PageEvents />
        <Providers>{children}</Providers>
      </body>
    </html>
  );
};

export default Layout;
