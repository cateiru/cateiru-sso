import {Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';

export const ProfileDatetime = () => {
  const textColor = useColorModeValue('gray.500', 'gray.400');
  const user = useRecoilValue(UserState);

  const created = React.useCallback(() => {
    if (user?.user.created_at) {
      const d = new Date(user.user.created_at);
      return `作成日: ${d.toLocaleString()}`;
    }
    return '-';
  }, [user?.user.created_at]);

  const modified = React.useCallback(() => {
    if (user?.user.modified_at) {
      const d = new Date(user.user.modified_at);
      return `更新日: ${d.toLocaleString()}`;
    }
    return '-';
  }, [user?.user.modified_at]);

  return (
    <Text color={textColor} textAlign="center">
      {created()} / {modified()}
    </Text>
  );
};
