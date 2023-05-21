'use client';

import {Heading, Select, Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {Margin} from '../Common/Margin';
import {LoginHistoriesTable} from './LoginHistoriesTable';
import {LoginTryTable} from './LoginTryTable';

export const LoginHistory = () => {
  const user = useRecoilValue(UserState);
  const userNameColor = useColorModeValue('gray.500', 'gray.400');

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
      <Text
        textAlign="center"
        mb="1rem"
        fontWeight="bold"
        color={userNameColor}
      >
        @{user?.user.user_name ?? '???'}
      </Text>
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
