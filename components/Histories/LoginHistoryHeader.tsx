'use client';

import {Box, Heading, Tab, TabList, Tabs, Text} from '@chakra-ui/react';
import Link from 'next/link';
import {usePathname} from 'next/navigation';
import React from 'react';
import {Margin} from '../Common/Margin';
import {UserName} from '../Common/UserName';
import {useSecondaryColor} from '../Common/useColor';

interface Histories {
  title: string;
  path: string;
  description?: string | React.ReactNode;
}

const histories: Histories[] = [
  {
    title: 'ログイン履歴',
    path: '/histories',
    description: (
      <>
        このアカウントにログインした履歴です。
        <br />
        ログインした日時、IPアドレス、ブラウザ、OSが表示されます。
      </>
    ),
  },
  {
    title: 'ログイントライ履歴',
    path: '/histories/try',
    description: (
      <>
        このアカウントにログインを試みた履歴です。
        <br />
        記録はログイン時とパスワード再登録時に成功、失敗時関わらずどちらも記録されます。
        <br />
        ログインした日時、IPアドレス、ブラウザ、OSが表示されます。
      </>
    ),
  },
  {
    title: '操作履歴',
    path: '/histories/operation',
    description: 'プロフィール変更などの操作に関する履歴です。',
  },
];

export const LoginHistoryHeader: React.FC<{children: React.ReactNode}> = ({
  children,
}) => {
  const pathname = usePathname();

  const textColor = useSecondaryColor();

  const history = React.useMemo(() => {
    return histories.find(v => v.path === pathname);
  }, [pathname]);

  const settingIndex = React.useMemo(() => {
    return histories.findIndex(v => v.path === pathname);
  }, [pathname]);

  return (
    <Margin>
      <Heading mt="3rem" mb="1rem" textAlign="center">
        {history?.title ?? 'ログイン履歴'}
      </Heading>

      <UserName />
      <Box
        overflowX={{base: 'auto', md: 'visible'}}
        pb=".1rem"
        px=".5rem"
        mb="1rem"
      >
        <Tabs
          isFitted
          index={settingIndex}
          mt="1rem"
          minW={{base: '650px', md: '100%'}}
          colorScheme="cateiru"
          fontWeight="bold"
        >
          <TabList>
            {histories.map(v => {
              return (
                <Tab
                  value={v.path}
                  key={`history-${v.path}`}
                  as={Link}
                  replace={true}
                  href={v.path}
                >
                  {v.title}
                </Tab>
              );
            })}
          </TabList>
        </Tabs>
      </Box>
      {history?.description ? (
        <Text color={textColor} textAlign="center" mb=".5rem">
          {history?.description}
        </Text>
      ) : (
        ''
      )}
      {children}
    </Margin>
  );
};
