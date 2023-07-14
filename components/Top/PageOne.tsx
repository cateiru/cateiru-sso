import {
  Box,
  Button,
  ButtonGroup,
  Center,
  Divider,
  Heading,
} from '@chakra-ui/react';
import {config} from '../../utils/config';
import {Logo} from '../Common/Icons/Logo';
import {Link} from '../Common/Next/Link';

export const PageOne = () => {
  return (
    <Center h="100vh">
      <Box>
        <Box w={{base: '150px', sm: '200px', md: '300px'}} mx="auto">
          <Logo size="100%" />
        </Box>
        <Heading
          textAlign="center"
          background="linear-gradient(124deg, #2bc4cf, #572bcf, #cf2ba1)"
          backgroundClip="text"
          fontSize={{base: '2rem', md: '3rem'}}
        >
          {config.title}
        </Heading>
        <Center mt="1rem">
          <ButtonGroup>
            <Button colorScheme="cateiru" as={Link} href="/register">
              アカウント作成
            </Button>
            <Button as={Link} href="/login">
              ログイン
            </Button>
          </ButtonGroup>
        </Center>
        <Center height="100px" mt="2.5rem">
          <Divider orientation="vertical" borderWidth="1.5px" />
        </Center>
      </Box>
    </Center>
  );
};
