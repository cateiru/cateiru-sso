import {
  Table,
  Tbody,
  Tr,
  Td,
  Avatar,
  Text,
  Box,
  useColorModeValue,
  Tooltip,
  Skeleton,
} from '@chakra-ui/react';
import React from 'react';
import {TbCheck} from 'react-icons/tb';
import {useRecoilValue} from 'recoil';
import useSWR from 'swr';
import {UserState} from '../../utils/state/atom';
import {accountUserFeather} from '../../utils/swr/featcher';
import {AccountUserList} from '../../utils/types/account';
import {ErrorType} from '../../utils/types/error';
import {Error} from '../Common/Error/Error';
import {useSwitchAccount} from './useSwitchAccount';

export const AccountList = () => {
  const checkMarkColor = useColorModeValue('#68D391', '#38A169');

  const user = useRecoilValue(UserState);
  const {data, error} = useSWR<AccountUserList, ErrorType>(
    '/',
    accountUserFeather
  );

  const {switch: s} = useSwitchAccount();

  if (error) {
    return <Error {...error} />;
  }

  if (!data) {
    return (
      <Box overflowY="auto" maxH="calc(100% - 100px)">
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td>
                <Skeleton w="48px" h="48px" borderRadius="50%" />
              </Td>
              <Td>
                <Skeleton w={{base: '230px', sm: '250px'}} h="20px" />
              </Td>
            </Tr>
            <Tr>
              <Td>
                <Skeleton w="48px" h="48px" borderRadius="50%" />
              </Td>
              <Td>
                <Skeleton w={{base: '230px', sm: '250px'}} h="20px" />
              </Td>
            </Tr>
            <Tr>
              <Td>
                <Skeleton w="48px" h="48px" borderRadius="50%" />
              </Td>
              <Td>
                <Skeleton w={{base: '230px', sm: '250px'}} h="20px" />
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </Box>
    );
  }

  return (
    <Box overflowY="auto" maxH="calc(100% - 100px)">
      <Table variant="simple">
        <Tbody>
          {data.map(account => {
            const isCurrentUser = account.id === user?.user.id;

            return (
              <Tr
                _hover={isCurrentUser ? {} : {bgColor: 'gray.100'}}
                w="100%"
                key={account.id}
                onClick={() => {
                  if (!isCurrentUser) {
                    s(account.id, account.user_name);
                  }
                }}
              >
                <Td>
                  <Avatar src={account.avatar ?? ''} />
                </Td>
                <Td w="100%">
                  <Text
                    fontSize="1.5rem"
                    fontWeight="bold"
                    textAlign="center"
                    maxW={
                      isCurrentUser
                        ? {base: '170px', sm: '215px'}
                        : {base: '230px', sm: '250px'}
                    }
                    textOverflow="ellipsis"
                    overflow="hidden"
                    whiteSpace="nowrap"
                  >
                    {account.user_name}
                  </Text>
                </Td>
                {isCurrentUser ? (
                  <Td>
                    <Tooltip
                      label="現在ログインしているユーザーです"
                      hasArrow
                      borderRadius="7px"
                    >
                      <Box>
                        <TbCheck
                          size="30px"
                          color={checkMarkColor}
                          strokeWidth="3px"
                        />
                      </Box>
                    </Tooltip>
                  </Td>
                ) : (
                  <Td p="0"></Td>
                )}
              </Tr>
            );
          })}
        </Tbody>
      </Table>
    </Box>
  );
};
