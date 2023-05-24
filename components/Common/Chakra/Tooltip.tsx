import {Tooltip as ChakraTooltip, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {colorTheme} from '../../../utils/theme';

export const Tooltip: typeof ChakraTooltip = props => {
  const bgColor = useColorModeValue('my.primary', 'my.secondary');
  const textColor = useColorModeValue(
    colorTheme.darkText,
    colorTheme.lightText
  );

  return (
    <ChakraTooltip
      hasArrow
      borderRadius="7px"
      px=".7rem"
      bgColor={bgColor}
      color={textColor}
      {...props}
    />
  );
};
