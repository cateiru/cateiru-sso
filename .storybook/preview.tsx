import {ChakraProvider, useColorMode} from '@chakra-ui/react';
import React from 'react';
import {theme} from '../utils/theme';
import {GoogleReCaptchaProvider} from 'react-google-recaptcha-v3';
import {RecoilRoot} from 'recoil';

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
      <RecoilRoot>
        <ChakraProvider theme={theme}>
          <GoogleReCaptchaProvider reCaptchaKey="empty_recaptcha_key">
            <ColorMode colorMode={context.globals.colorMode}>
              <Story />
            </ColorMode>
          </GoogleReCaptchaProvider>
        </ChakraProvider>
      </RecoilRoot>
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
