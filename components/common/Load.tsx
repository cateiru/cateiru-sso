import {Box, Flex} from '@chakra-ui/react';
import {useRecoilValue} from 'recoil';
import {LoadState} from '../../utils/state/atom';
import Spinner from './Spinner';

const Load = () => {
  const load = useRecoilValue(LoadState);

  return (
    <>
      {load && (
        <>
          <Box
            backgroundColor="gray.400"
            opacity=".5"
            position="fixed"
            width="100vw"
            height="100vh"
            zIndex="9999"
            top="0"
            left="0"
          ></Box>
          <Flex
            width="100vw"
            height="100vh"
            position="fixed"
            top="0"
            left="0"
            zIndex="9999"
            flexDirection="column"
            justifyContent="center"
            alignItems="center"
          >
            <Spinner />
          </Flex>
        </>
      )}
    </>
  );
};

export default Load;
