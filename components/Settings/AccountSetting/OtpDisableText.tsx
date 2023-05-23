import {
  Button,
  Spacer,
  Stack,
  Text,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import {OtpRegister} from './OtpRegister';

export const OtpDisableText = () => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  const {isOpen, onOpen, onClose} = useDisclosure();

  return (
    <>
      <Stack
        direction={{base: 'column', md: 'row'}}
        alignItems={{base: undefined, md: 'center'}}
      >
        <Text color={textColor}>二段階認証は現在設定されていません。</Text>
        <Spacer />
        <Button colorScheme="cateiru" onClick={onOpen}>
          二段階認証を設定する
        </Button>
      </Stack>
      <OtpRegister isOpen={isOpen} onClose={onClose} />
    </>
  );
};
