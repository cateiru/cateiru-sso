import {extendTheme} from '@chakra-ui/react';
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
});

export default theme;
