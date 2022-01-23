import {Flex, Button, Box, Text, Center, Stack} from '@chakra-ui/react';
import NextLink from 'next/link';
import Logo from '../common/Logo';

const TopPage = () => {
  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <Box marginBottom="1rem">
        <Box width={{base: '20rem', sm: '30rem', md: '50rem', lg: '60rem'}}>
          <Logo />
        </Box>
        <Center>
          <Text fontSize={{base: '1rem', md: '1.5rem'}}>
            CateiruのSSOサービス
          </Text>
        </Center>
      </Box>
      <Stack direction={['column', 'row']} spacing="10px">
        <Button colorScheme="blue" variant="solid">
          ログイン
        </Button>
        <NextLink href="/create" passHref>
          <Button colorScheme="blue" variant="outline">
            新規登録
          </Button>
        </NextLink>
      </Stack>
    </Flex>
  );
};

export default TopPage;
