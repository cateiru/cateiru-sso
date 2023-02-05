import {Box, Center} from '@chakra-ui/react';
import Spinner from '../../common/Spinner';

const LoginSkeleton = () => {
  return (
    <Center>
      <Box
        width={{base: '95%', sm: '400px'}}
        height="600px"
        mt={{base: '0', sm: '3rem'}}
        borderRadius="20px"
        borderWidth={{base: '0', sm: '2px'}}
      >
        <Center height="100%">
          <Spinner />
        </Center>
      </Box>
    </Center>
  );
};

export default LoginSkeleton;
