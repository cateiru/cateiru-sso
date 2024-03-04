import {Text, TextProps} from '@chakra-ui/react';
import {useAtomValue} from 'jotai';
import React from 'react';
import {UserState} from '../../utils/state/atom';
import {useSecondaryColor} from './useColor';

export const UserName = React.memo<TextProps>(props => {
  const user = useAtomValue(UserState);

  const userNameColor = useSecondaryColor();

  return (
    <Text
      textAlign="center"
      mb="1rem"
      fontWeight="bold"
      color={userNameColor}
      {...props}
    >
      @{user?.user.user_name ?? '???'}
    </Text>
  );
});

UserName.displayName = 'UserName';
