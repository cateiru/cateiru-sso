import {
  Avatar as ChakraAvatar,
  type AvatarProps,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';

export const Avatar: React.FC<AvatarProps> = props => {
  const shadowColor = useColorModeValue('#242838', '#000');

  if (props.src === '') {
    return <ChakraAvatar {...props} key={`no-image-avatar-${Math.random()}`} />;
  }

  return (
    <ChakraAvatar {...props} boxShadow={`0px 0px 7px -2px ${shadowColor}`} />
  );
};
