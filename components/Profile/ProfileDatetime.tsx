import {Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';

export const ProfileDatetime = () => {
  const textColor = useColorModeValue('gray.500', 'gray.400');
  const user = useRecoilValue(UserState);

  const created = React.useCallback(() => {
    if (user?.user.created) {
      const d = new Date(user.user.created);
      return `作成日: ${d.toLocaleString()}`;
    }
    return '-';
  }, [user?.user.created]);

  const modified = React.useCallback(() => {
    if (user?.user.modified) {
      const d = new Date(user.user.modified);
      return `更新日: ${d.toLocaleString()}`;
    }
    return '-';
  }, [user?.user.modified]);

  return (
    <Text color={textColor} textAlign="center">
      {created()} / {modified()}
    </Text>
  );
};
