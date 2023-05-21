import {Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';

export const UserName = React.memo(() => {
  const user = useRecoilValue(UserState);

  const userNameColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <Text textAlign="center" mb="1rem" fontWeight="bold" color={userNameColor}>
      @{user?.user.user_name ?? '???'}
    </Text>
  );
});

UserName.displayName = 'UserName';
