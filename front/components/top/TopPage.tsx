import {Flex, Box, Text, Center} from '@chakra-ui/react';
import LoginButtons from '../createAccount/LoginButtons';
import AnimationLogo from './AnimationLogo';

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
        <Box width={{base: '20rem', md: '30rem', lg: '45rem'}}>
          <AnimationLogo />
        </Box>
        <Center>
          <Text fontSize={{base: '1rem', md: '1.5rem'}}>
            CateiruのSSOサービス
          </Text>
        </Center>
      </Box>
      <Box width={{base: '20rem', sm: 'auto'}}>
        <LoginButtons />
      </Box>
    </Flex>
  );
};

export default TopPage;
