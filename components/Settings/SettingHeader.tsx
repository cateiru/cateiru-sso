'use client';

import {Box, Heading, Tab, TabList, Tabs} from '@chakra-ui/react';
import Link from 'next/link';
import {usePathname} from 'next/navigation';
import React from 'react';
import {Margin} from '../Common/Margin';
import {UserName} from '../Common/UserName';

export const SettingHeader: React.FC<{children: React.ReactNode}> = ({
  children,
}) => {
  const pathname = usePathname();

  const settingTitle = React.useCallback(() => {
    switch (pathname) {
      case '/settings':
        return 'アカウント設定';
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
      <Box overflowX="auto" pb=".1rem">
        <Tabs
          isFitted
          defaultIndex={settingIndex()}
          mt="1rem"
          w="650px"
          colorScheme="cateiru"
          fontWeight="bold"
        >
          <TabList>
            <Tab as={Link} href="/settings">
              アカウント設定
            </Tab>
            <Tab as={Link} href="/settings/email">
              メールアドレス設定
            </Tab>
            <Tab as={Link} href="/settings/password">
              パスワード設定
            </Tab>
          </TabList>
        </Tabs>
      </Box>

      {children}
    </Margin>
  );
};
