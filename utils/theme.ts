import {ThemeConfig, UseToastOptions, extendTheme} from '@chakra-ui/react';
import {mode} from '@chakra-ui/theme-tools';
import {StepsTheme} from 'chakra-ui-steps';

const config: ThemeConfig = {
  useSystemColorMode: false,
};

interface ColorThemes {
  darkBackground: string;
  lightBackground: string;
  darkText: string;
  lightText: string;
}

export const colorTheme: ColorThemes = {
  darkBackground: '#242838',
  lightBackground: '#ffffff',
  darkText: '#e8e8e8',
  lightText: '#1f1f1f',
};

export const toastOptions: UseToastOptions = {
  duration: 5000,
  isClosable: true,
};

export const theme = extendTheme({
  fonts: {
    heading: "'Noto Sans JP', sans-serif",
    body: "'Noto Sans JP', sans-serif",
  },
  components: {
    CloseButton: {
      baseStyle: {
        _focus: {
          boxShadow: 'none',
        },
      },
    },
    Steps: StepsTheme,
    Input: {
      defaultProps: {
        focusBorderColor: 'my.secondary',
      },
    },
    PinInput: {
      defaultProps: {
        focusBorderColor: 'my.secondary',
      },
    },
    Select: {
      defaultProps: {
        focusBorderColor: 'my.secondary',
      },
    },
  },
  colors: {
    cateiru: {
      100: '#d8eef0',
      200: '#c9f2f5',
      300: '#a1e1e6',
      400: '#6ad6de',
      500: '#2bc4cf',
      600: '#1fb4bf',
      700: '#169ca6',
      800: '#0f818a',
      900: '#06585e',
    },
    brand: {
      200: '#E2E8F0',
      300: '#CBD5E0',
      500: '#404663',
      600: '#343952',
    },
    my: {
      primary: '#572bcf',
      secondary: '#2bc4cf',
      accent: '#cf2ba1',
    },
  },
  styles: {
    global: (props: {colorMode: string}) => ({
      // Chrome
      '&::-webkit-scrollbar': {
        width: '7px',
        height: '5px',
      },
      '&::-webkit-scrollbar-thumb': {
        backgroundColor: props.colorMode === 'dark' ? 'brand.600' : 'gray.400',
        borderRadius: '100px',
        ':hover': {
          backgroundColor:
            props.colorMode === 'dark' ? 'brand.500' : 'brand.500',
        },
      },
      '::-webkit-scrollbar-track': {
        backgroundColor: 'rgba(0,0,0,0)',
      },
      // FireFox
      html: {
        scrollbarWidth: 'thin',
        scrollbarColor: props.colorMode === 'dark' ? 'brand.600' : 'gray.400',
        overflow: 'overlay',
      },
      body: {
        // background: mode(
        //   colorTheme.lightBackground,
        //   colorTheme.darkBackground
        // )(props),
        // color:
        //   props.colorMode === 'dark'
        //     ? colorTheme.darkText
        //     : colorTheme.lightText,
      },
      ':root': {
        '--background-color':
          props.colorMode === 'dark'
            ? colorTheme.darkBackground
            : colorTheme.lightBackground,
        '--text-color':
          props.colorMode === 'dark'
            ? colorTheme.darkText
            : colorTheme.lightText,
      },
    }),
  },
  config: config,
});
