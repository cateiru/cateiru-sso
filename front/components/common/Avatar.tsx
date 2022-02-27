import {Avatar as ChakraAvatar, useColorMode} from '@chakra-ui/react';
import type {AvatarProps} from '@chakra-ui/react';

const Avatar: React.FC<AvatarProps & {isShadow?: boolean}> = props => {
  const {colorMode} = useColorMode();

  return (
    <ChakraAvatar
      {...props}
      boxShadow={
        props.isShadow
          ? colorMode === 'dark'
            ? '10px 10px 30px #000000CC'
            : '10px 10px 30px #A0AEC0B3'
          : undefined
      }
      bgColor={props.src && 'white'}
    />
  );
};

export default Avatar;
