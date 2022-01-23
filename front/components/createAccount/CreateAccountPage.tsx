import {Flex} from '@chakra-ui/react';
import UserPassword from './UserPassword';

const CreateAccountPage = () => {
  return (
    <Flex
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100%"
      height="80vh"
      px={{base: '1rem', md: '5rem'}}
    >
      <UserPassword />
    </Flex>
  );
};

export default CreateAccountPage;
