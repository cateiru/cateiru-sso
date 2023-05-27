import {Button, Text, useColorModeValue} from '@chakra-ui/react';
import Link from 'next/link';

export const OtpImpossible = () => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <>
      <Text color={textColor}>
        パスワードが設定されていないため二段階認証は現在使用できません。
      </Text>
      <Button
        mt="1rem"
        w="100%"
        colorScheme="cateiru"
        as={Link}
        href="/settings/password"
      >
        パスワードを設定する
      </Button>
    </>
  );
};
