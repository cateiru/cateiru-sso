import {Text, TextProps} from '@chakra-ui/react';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {useSecondaryColor} from './useColor';

export const UserName = React.memo<TextProps>(props => {
  const user = useRecoilValue(UserState);

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
