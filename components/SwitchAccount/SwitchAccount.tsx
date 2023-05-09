import {Box, Divider, Heading, useColorModeValue} from '@chakra-ui/react';
import {AccountList} from './AccountList';

export const SwitchAccount = () => {
  const borderColor = useColorModeValue('gray.300', 'gray.600');

  return (
    <Box
      w={{base: '96%', sm: '450px'}}
      h={{base: '600px', sm: '700px'}}
      borderWidth={{base: 'none', sm: '2px'}}
      margin="auto"
      mt="3rem"
      borderRadius="10px"
      borderColor={borderColor}
      mb={{base: '0', sm: '3rem'}}
    >
      <Box h="150px">
        <Heading textAlign="center" pt="2rem" mx=".5rem">
          ログインするアカウントを選択してください
        </Heading>
        <Divider mt="1.5rem" w="90%" mx="auto" />
      </Box>
      <AccountList />
    </Box>
  );
};
