import {Box, Divider, Heading} from '@chakra-ui/react';
import {AccountFeather} from './AccountFeather';

export const SwitchAccount = () => {
  return (
    <Box
      w={{base: '96%', sm: '450px'}}
      h={{base: '600px', sm: '700px'}}
      borderWidth={{base: 'none', sm: '2px'}}
      margin="auto"
      mt="3rem"
      borderRadius="10px"
      borderColor="gray.300"
      mb={{base: '0', sm: '3rem'}}
    >
      <Box h="100px">
        <Heading textAlign="center" pt="2rem">
          アカウントを選択
        </Heading>
        <Divider mt=".5rem" w="90%" mx="auto" />
      </Box>
      <AccountFeather />
    </Box>
  );
};
