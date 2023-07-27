import {Text} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {useSecondaryColor} from '../Common/useColor';

export const ProfileDatetime = () => {
  const textColor = useSecondaryColor();
  const user = useRecoilValue(UserState);

  const created = React.useCallback(() => {
    if (user?.user.created_at) {
      const d = new Date(user.user.created_at);
      return `作成日: ${d.toLocaleString()}`;
    }
    return '-';
  }, [user?.user.created_at]);

  const modified = React.useCallback(() => {
    if (user?.user.updated_at) {
      const d = new Date(user.user.updated_at);
      return `更新日: ${d.toLocaleString()}`;
    }
    return '-';
  }, [user?.user.updated_at]);

  return (
    <Text color={textColor} textAlign="center">
      {created()} / {modified()}
    </Text>
  );
};
