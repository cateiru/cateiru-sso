'use client';

import {Heading, Select, Text, useColorModeValue} from '@chakra-ui/react';
import {usePathname, useRouter} from 'next/navigation';
import React from 'react';
import {routeChangeStart} from '../../utils/event';
import {Margin} from '../Common/Margin';
import {UserName} from '../Common/UserName';

export const LoginHistoryHeader: React.FC<{children: React.ReactNode}> = ({
  children,
}) => {
  const pathname = usePathname();
  const router = useRouter();

  const textColor = useColorModeValue('gray.500', 'gray.400');

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
      <Text color={textColor} textAlign="center" mb=".5rem">
        {title() === 'ログイン履歴' ? (
          <>
            このアカウントにログインした履歴です。
            <br />
            ログインした日時、IPアドレス、ブラウザ、OSが表示されます。
          </>
        ) : (
          <>
            このアカウントにログインを試みた履歴です。
            <br />
            記録はログイン時とパスワード再登録時に成功、失敗時関わらずどちらも記録されます。
            <br />
            ログインした日時、IPアドレス、ブラウザ、OSが表示されます。
          </>
        )}
      </Text>
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
        <option value="/histories">ログイン履歴</option>
        <option value="/histories/try">ログイントライ履歴</option>
      </Select>
      {children}
    </Margin>
  );
};
