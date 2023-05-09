import {
  Table,
  Tbody,
  Tr,
  Td,
  Avatar,
  Text,
  Button,
  Box,
} from '@chakra-ui/react';
import React from 'react';
import {AccountUserList} from '../../utils/types/account';

export const AccountList: React.FC<{data: AccountUserList}> = ({data}) => {
  return (
    <Box overflowY="auto" maxH="calc(100% - 100px)">
      <Table variant="simple">
        <Tbody>
          {data.map(account => {
            return (
              <Button key={account.id} h="100%" variant="unstyled" w="100%">
                <Tr _hover={{bgColor: 'gray.100'}}>
                  <Td>
                    <Avatar src={account.avatar ?? ''} />
                  </Td>
                  <Td w="100%">
                    <Text
                      fontSize="1.5rem"
                      fontWeight="bold"
                      textAlign="center"
                    >
                      {account.user_name}
                    </Text>
                  </Td>
                </Tr>
              </Button>
            );
          })}
        </Tbody>
      </Table>
    </Box>
  );
};
