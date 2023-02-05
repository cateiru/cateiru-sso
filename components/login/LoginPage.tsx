import {Flex} from '@chakra-ui/react';
import LoginForm from './LoginForm';

const LoginPage = () => {
  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <LoginForm />
    </Flex>
  );
};

export default LoginPage;
