'use client';

import {Box, Heading, Tab, TabList, Tabs} from '@chakra-ui/react';
import {usePathname} from 'next/navigation';
import React from 'react';
import {Margin} from '../Common/Margin';
import {Link} from '../Common/Next/Link';
import {UserName} from '../Common/UserName';

export const SettingHeader: React.FC<{children: React.ReactNode}> = ({
  children,
}) => {
  const pathname = usePathname();

  const settingTitle = React.useCallback(() => {
    switch (pathname) {
      case '/settings':
        return '設定';
      case '/settings/email':
        return 'メールアドレス設定';
      case '/settings/password':
        return 'パスワード設定';
      default:
        return '設定';
    }
  }, [pathname]);

  const settingIndex = React.useCallback(() => {
    switch (pathname) {
      case '/settings':
        return 0;
      case '/settings/email':
        return 1;
      case '/settings/password':
        return 2;
      default:
        return 0;
    }
  }, [pathname]);

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem" mb="1rem">
        {settingTitle()}
      </Heading>
      <UserName />
      <Box overflowX={{base: 'auto', md: 'visible'}} pb=".1rem" px=".5rem">
        <Tabs
          isFitted
          index={settingIndex()}
          mt="1rem"
          minW={{base: '650px', md: '100%'}}
          colorScheme="cateiru"
          fontWeight="bold"
        >
          <TabList>
            <Tab as={Link} href="/settings" replace={true}>
              設定
            </Tab>
            <Tab as={Link} href="/settings/email" replace={true}>
              メールアドレス設定
            </Tab>
            <Tab as={Link} href="/settings/password" replace={true}>
              パスワード設定
            </Tab>
          </TabList>
        </Tabs>
      </Box>

      {children}
    </Margin>
  );
};
