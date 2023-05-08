import {Box, Divider, Link, Text} from '@chakra-ui/react';
import NextLink from 'next/link';

export const Footer = () => {
  return (
    <Box marginTop="auto" as="footer">
      <Box width="95%" margin="1rem auto 1rem auto">
        <Divider />
      </Box>
      <Text textAlign="center" mb="1.5rem">
        &copy; {new Date().getFullYear()}{' '}
        <NextLink href="/">
          <Text as="span" _hover={{borderBottom: '1px'}}>
            cateiru
          </Text>
        </NextLink>{' '}
        -{' '}
        <Link href="https://github.com/cateiru/cateiru-sso" isExternal>
          GitHub
        </Link>
      </Text>
    </Box>
  );
};
