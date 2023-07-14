'use client';

import {Heading, Text, useColorModeValue} from '@chakra-ui/react';
import {Margin} from '../Common/Margin';
import {UserName} from '../Common/UserName';
import {OrgListTable} from './OrgListTable';

export const OrgList = () => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        組織一覧
      </Heading>
      <Text color={textColor} mt=".5rem" textAlign="center" mb=".5rem">
        あなたの所属している組織の一覧表示されます。
        <br />
        組織のオーナーの場合、組織のユーザー追加や削除ができます。
      </Text>
      <UserName />
      <OrgListTable />
    </Margin>
  );
};
