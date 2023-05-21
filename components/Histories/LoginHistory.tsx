'use client';

import {Heading, Select} from '@chakra-ui/react';
import React from 'react';
import {Margin} from '../Common/Margin';
import {UserName} from '../Common/UserName';
import {LoginHistoriesTable} from './LoginHistoriesTable';
import {LoginTryTable} from './LoginTryTable';

export const LoginHistory = () => {
  const [select, setSelect] = React.useState('1');

  const C = React.useCallback(() => {
    switch (select) {
      case '1':
        return <LoginHistoriesTable />;
      case '2':
        return <LoginTryTable />;
      default:
        return null;
    }
  }, [select]);

  return (
    <Margin>
      <Heading mt="3rem" mb="1rem" textAlign="center">
        ログイン履歴
      </Heading>
      <UserName />
      <Select
        w={{base: '100%', md: '300px'}}
        mb="1rem"
        size="md"
        mx="auto"
        onChange={v => setSelect(v.target.value)}
        value={select}
      >
        <option value="1">ログイン履歴</option>
        <option value="2">ログイントライ履歴</option>
      </Select>
      <C />
    </Margin>
  );
};
