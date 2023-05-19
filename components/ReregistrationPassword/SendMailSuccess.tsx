import {Box, Heading, Text, VStack, useColorModeValue} from '@chakra-ui/react';
import React from 'react';
import {CheckMark, CheckmarkProps} from '../Common/Icons/CheckMark';

interface Props {
  email: string;
}

export const SendMainSuccess: React.FC<Props> = ({email}) => {
  const descriptionColor = useColorModeValue('gray.500', 'gray.400');
  const checkmarkProps = useColorModeValue<CheckmarkProps, CheckmarkProps>(
    {
      size: 100,
      bgColor: '#572bcf',
      color: '#fff',
    },
    {
      size: 100,
      bgColor: '#2bc4cf',
      color: '#fff',
    }
  );

  const hide = React.useCallback(() => {
    const matched = email.match(
      /^(?<prefix>[A-Z0-9._%+-])(?<host>[A-Z0-9._%+-]+)@(?<tdomain>[A-Z0-9.-]+)\.(?<domain>[A-Z]{2,})/i
    );
    if (!matched) return 'メールアドレス';

    const {prefix, host, tdomain, domain} = matched.groups!;

    return `${prefix}${Array(host.length).fill('*').join('')}@${Array(
      tdomain.length
    )
      .fill('*')
      .join('')}.${domain}`;
  }, [email]);

  return (
    <VStack mt="3rem">
      <CheckMark {...checkmarkProps} />
      <Heading textAlign="center" mt=".5rem" color={checkmarkProps.bgColor}>
        メールを送信しました
      </Heading>
      <Box>
        <Text mt=".5rem" textAlign="center" color={descriptionColor}>
          <Text as="span" fontWeight="bold">
            {hide()}
          </Text>{' '}
          にパスワード再設定メールを送信しました
        </Text>
        <Text textAlign="center" color={descriptionColor}>
          送られてきたメールのURLからパスワードを再設定してください
        </Text>
      </Box>
    </VStack>
  );
};
