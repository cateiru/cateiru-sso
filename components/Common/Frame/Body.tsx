'use client';

import {Box, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {colorTheme} from '../../../utils/theme';

/**
 * 本当はChakraUIのthremeでやりたいが、Next13でのバグでうまくcssが追加できないので諦めている
 * TODO: 治ったら消す
 */
export const Body: React.FC<{children: React.ReactNode}> = ({children}) => {
  const bgColor = useColorModeValue(
    colorTheme.lightBackground,
    colorTheme.darkBackground
  );
  const color = useColorModeValue(colorTheme.lightText, colorTheme.darkText);

  return (
    <Box w="100%" h="100%" bgColor={bgColor} color={color}>
      {children}
    </Box>
  );
};
