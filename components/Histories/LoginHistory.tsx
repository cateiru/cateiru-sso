'use client';

import {Heading, Text, useColorModeValue} from '@chakra-ui/react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {Margin} from '../Common/Margin';
import {LoginHistoriesTable} from './LoginHistoriesTable';

export const LoginHistory = () => {
  const user = useRecoilValue(UserState);

  const userNameColor = useColorModeValue('gray.500', 'gray.400');

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
      <LoginHistoriesTable />
    </Margin>
  );
};
