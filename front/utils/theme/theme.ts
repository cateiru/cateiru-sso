import {extendTheme, } from '@chakra-ui/react';
import {StepsStyleConfig as Steps} from 'chakra-ui-steps';

const theme = extendTheme({
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
    Steps,
  },
  styles: {
    global: (props: any) => ({
      // Chrome
      '&::-webkit-scrollbar': {
        width: '10px',
      },
      '&::-webkit-scrollbar-thumb': {
        backgroundColor: props.colorMode === 'dark' ? 'gray.600' : 'gray.300',
        borderRadius: '100px',
      },
      // FireFox
      html: {
        scrollbarWidth: 'thin',
        scrollbarColor: props.colorMode === 'dark' ? 'gray.600' : 'gray.300',
      }
    }),
  },
});

export default theme;
