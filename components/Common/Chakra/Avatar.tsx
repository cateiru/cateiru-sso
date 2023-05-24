import {Avatar as ChakraAvatar, type AvatarProps} from '@chakra-ui/react';
import React from 'react';

export const Avatar: React.FC<AvatarProps> = props => {
  if (props.src === '') {
    return <ChakraAvatar {...props} key={`no-image-avatar-${Math.random()}`} />;
  }

  return <ChakraAvatar {...props} />;
};
