import {Avatar as ChakraAvatar, useColorMode} from '@chakra-ui/react';
import type {AvatarProps} from '@chakra-ui/react';
import React from 'react';

const Avatar = React.memo<AvatarProps & {isShadow?: boolean}>(props => {
  const {colorMode} = useColorMode();

  const {isShadow, ...rest} = props;

  return (
    <ChakraAvatar
      boxShadow={
        isShadow
          ? colorMode === 'dark'
            ? '10px 10px 30px #000000CC'
            : '10px 10px 30px #A0AEC0B3'
          : undefined
      }
      bgColor={props.src && 'white'}
      {...rest}
    />
  );
});

Avatar.displayName = 'avatar';

export default Avatar;
