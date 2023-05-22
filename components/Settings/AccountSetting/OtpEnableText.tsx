import {Box, Button, Stack, Text, useColorModeValue} from '@chakra-ui/react';
import React from 'react';

interface Props {
  modified: Date;
}

export const OtpEnableText: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <Box>
      <Text color={textColor}>二段階認証は設定されています。</Text>
      <Text color={textColor} mb=".5rem">
        設定日時:
        <Text as="span" fontWeight="bold">
          {props.modified.toLocaleString()}
        </Text>
      </Text>
      <Stack direction={{base: 'column', md: 'row'}}>
        <Button w="100%" colorScheme="cateiru">
          バックアップコードを表示する
        </Button>
        <Button w="100%">二段階認証を削除する</Button>
      </Stack>
    </Box>
  );
};
