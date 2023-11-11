'use client';

import {Heading, Select, Text} from '@chakra-ui/react';
import {usePathname, useRouter} from 'next/navigation';
import React from 'react';
import {routeChangeStart} from '../../utils/event';
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
  const router = useRouter();

  const textColor = useSecondaryColor();

  const history = React.useMemo(() => {
    return histories.find(v => v.path === pathname);
  }, [pathname]);

  return (
    <Margin>
      <Heading mt="3rem" mb="1rem" textAlign="center">
        {history?.title ?? 'ログイン履歴'}
      </Heading>
      {history?.description ? (
        <Text color={textColor} textAlign="center" mb=".5rem">
          {history?.description}
        </Text>
      ) : (
        ''
      )}
      <UserName />
      <Select
        w={{base: '100%', md: '300px'}}
        mb="1rem"
        size="md"
        mx="auto"
        onChange={v => {
          routeChangeStart();
          router.replace(v.target.value);
        }}
        defaultValue={pathname}
      >
        {histories.map(v => {
          return (
            <option value={v.path} key={`history-${v.path}`}>
              {v.title}
            </option>
          );
        })}
      </Select>
      {children}
    </Margin>
  );
};
