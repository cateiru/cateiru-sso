'use client';

import {Heading} from '@chakra-ui/react';
import {Margin} from '../../Common/Margin';
import {UsersTable} from './UsersTable';

export const Users = () => {
  return (
    <Margin>
      <Heading textAlign="center" mt="3rem" mb="2rem">
        ユーザー一覧
      </Heading>
      <UsersTable />
    </Margin>
  );
};
