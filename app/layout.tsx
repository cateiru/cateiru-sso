import nprogress from 'nprogress';
import React from 'react';
import {Providers} from './providers';
import {usePageEvent} from './usePageEvent';

nprogress.configure({showSpinner: false, speed: 400, minimum: 0.25});

const Layout: React.FC<{children: React.ReactNode}> = ({children}) => {
  usePageEvent();

  return (
    <html lang="ja">
      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
};

export default Layout;
