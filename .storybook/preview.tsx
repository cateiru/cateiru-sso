import React from 'react';
import {theme} from '../utils/theme';
import {ChakraProvider, layout, useColorMode} from '@chakra-ui/react';

interface ColorModeProps {
  colorMode: 'light' | 'dark';
  children: JSX.Element;
}

function ColorMode(props: ColorModeProps) {
  const {setColorMode} = useColorMode();

  React.useEffect(() => {
    setColorMode(props.colorMode);
  }, [props.colorMode]);

  return props.children;
}

export const decorators = [
  (Story, context) => {
    return (
      <ChakraProvider theme={theme}>
        <ColorMode colorMode={context.globals.colorMode}>
          <Story />
        </ColorMode>
      </ChakraProvider>
    );
  },
];

export const globalTypes = {
  colorMode: {
    name: 'Color Mode',
    defaultValue: 'light',
    toolbar: {
      items: [
        {title: 'Light', value: 'light'},
        {title: 'Dark', value: 'dark'},
      ],
      dynamicTitle: true,
    },
  },
};
