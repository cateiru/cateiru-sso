import {Avatar as ChakraAvatar, type AvatarProps} from '@chakra-ui/react';
import React from 'react';
import {useShadowColor} from '../useColor';

export const Avatar: React.FC<AvatarProps> = props => {
  const shadowColor = useShadowColor();

  if (props.src === '') {
    return <ChakraAvatar {...props} key={`no-image-avatar-${Math.random()}`} />;
  }

  return (
    <ChakraAvatar {...props} boxShadow={`0px 0px 7px -2px ${shadowColor}`} />
  );
};
