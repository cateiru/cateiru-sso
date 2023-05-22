import {Button, Spacer, Stack, Text, useColorModeValue} from '@chakra-ui/react';

export const OtpDisableText = () => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <Stack
      direction={{base: 'column', md: 'row'}}
      alignItems={{base: undefined, md: 'center'}}
    >
      <Text color={textColor}>二段階認証は現在設定されていません。</Text>
      <Spacer />
      <Button colorScheme="cateiru">二段階認証を設定する</Button>
    </Stack>
  );
};
