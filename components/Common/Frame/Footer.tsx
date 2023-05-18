import {Box, Divider, Link, Text} from '@chakra-ui/react';

export const Footer = () => {
  return (
    <Box marginTop="auto" as="footer">
      <Box width="95%" margin="1rem auto 1rem auto">
        <Divider />
      </Box>
      <Text textAlign="center" mb="1.5rem">
        &copy; {new Date().getFullYear()}{' '}
        <Link href="https://cateiru.com" isExternal>
          cateiru
        </Link>{' '}
        -{' '}
        <Link href="https://github.com/cateiru/cateiru-sso" isExternal>
          GitHub
        </Link>
      </Text>
    </Box>
  );
};
