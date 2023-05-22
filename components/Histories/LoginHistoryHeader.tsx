'use client';

import {Heading, Select} from '@chakra-ui/react';
import {usePathname, useRouter} from 'next/navigation';
import React from 'react';
import {Margin} from '../Common/Margin';
import {UserName} from '../Common/UserName';

export const LoginHistoryHeader: React.FC<{children: React.ReactNode}> = ({
  children,
}) => {
  const pathname = usePathname();
  const router = useRouter();

  const title = React.useCallback(() => {
    switch (pathname) {
      case '/histories':
        return 'ログイン履歴';
      case '/histories/try':
        return 'ログイントライ履歴';
      default:
        return 'ログイン履歴';
    }
  }, [pathname]);

  return (
    <Margin>
      <Heading mt="3rem" mb="1rem" textAlign="center">
        {title()}
      </Heading>
      <UserName />
      <Select
        w={{base: '100%', md: '300px'}}
        mb="1rem"
        size="md"
        mx="auto"
        onChange={v => router.replace(v.target.value)}
        defaultValue={pathname}
      >
        <option value="/histories">ログイン履歴</option>
        <option value="/histories/try">ログイントライ履歴</option>
      </Select>
      {children}
    </Margin>
  );
};
