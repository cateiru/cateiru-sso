import {menuAnatomy} from '@chakra-ui/anatomy';
import {
  ThemeConfig,
  UseToastOptions,
  extendTheme,
  defineStyleConfig,
  createMultiStyleConfigHelpers,
} from '@chakra-ui/react';
import {mode} from '@chakra-ui/theme-tools';
import {StepsTheme} from 'chakra-ui-steps';

const multiStyleConfigHelpers = createMultiStyleConfigHelpers(menuAnatomy.keys);

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
    Textarea: {
      defaultProps: {
        focusBorderColor: 'my.secondary',
      },
    },
    Switch: {
      defaultProps: {
        colorScheme: 'cateiru',
      },
    },
    Heading: {
      baseStyle: {
        background: 'linear-gradient(132deg, #17C9C9 0%, #336CFF 100%);',
        backgroundClip: 'text',
      },
    },
    Modal: defineStyleConfig({
      baseStyle: props => ({
        dialog: {
          bg:
            props.colorMode === 'dark'
              ? colorTheme.darkBackground
              : colorTheme.lightBackground,
        },
      }),
    }),
    Tooltip: defineStyleConfig({
      baseStyle: props => ({
        color:
          props.colorMode === 'dark'
            ? colorTheme.lightText
            : colorTheme.darkText,
        background: props.colorMode === 'dark' ? 'my.secondary' : 'my.primary',
        borderRadius: '7px',
        px: '.7rem',
        fontWeight: 'bold',
        hasArrow: true,
        div: {
          div: {
            background:
              props.colorMode === 'dark'
                ? 'my.secondary !important'
                : 'my.primary !important',
          },
        },
      }),
    }),
    Menu: multiStyleConfigHelpers.defineMultiStyleConfig({
      baseStyle: multiStyleConfigHelpers.definePartsStyle({
        list: {
          _dark: {
            '--menu-bg': colorTheme.darkBackground,
          },
        },
        item: {
          _dark: {
            '--menu-bg': colorTheme.darkBackground,
          },
        },
      }),
    }),
  },
  colors: {
    cateiru: {
      100: '#b7ecf0',
      200: '#93e3e9',
      300: '#6fdae1',
      400: '#4cd0da',
      500: '#2bc4cf',
      600: '#24a3ad',
      700: '#1d838a',
      800: '#166268',
      900: '#0e4145',
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
      html: {
        scrollbarWidth: 'thin',
        scrollbarColor: props.colorMode === 'dark' ? 'brand.600' : 'gray.400',
        scrollbarGutter: 'stable',
        fontFamily: "'Noto Sans JP', sans-serif",
      },
      pre: {
        fontFamily: "'Noto Sans JP', sans-serif",
      },
      body: {
        background: mode(
          colorTheme.lightBackground,
          colorTheme.darkBackground
        )(props),
        color:
          props.colorMode === 'dark'
            ? colorTheme.darkText
            : colorTheme.lightText,
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
