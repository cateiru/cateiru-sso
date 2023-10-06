import {Tooltip as ChakraTooltip} from '@chakra-ui/react';
import React from 'react';

export const Tooltip: typeof ChakraTooltip = props => {
  return <ChakraTooltip hasArrow {...props} />;
};
