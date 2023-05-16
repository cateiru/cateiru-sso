import {
  Table,
  Tbody,
  Tr,
  Td,
  Text,
  Box,
  useColorModeValue,
  Tooltip,
  Skeleton,
  Button,
  Center,
} from '@chakra-ui/react';
import Link from 'next/link';
import React from 'react';
import {TbCheck, TbUserPlus} from 'react-icons/tb';
import {useRecoilValue} from 'recoil';
import useSWR, {useSWRConfig} from 'swr';
import {UserState} from '../../utils/state/atom';
import {accountUserFeather} from '../../utils/swr/featcher';
import {AccountUserList} from '../../utils/types/account';
import {ErrorType} from '../../utils/types/error';
import {Avatar} from '../Common/Avatar';
import {Error} from '../Common/Error/Error';
import {useSwitchAccount} from './useSwitchAccount';

export const AccountList = () => {
  const checkMarkColor = useColorModeValue('#68D391', '#38A169');
  const hoverColor = useColorModeValue('gray.100', 'gray.600');

  const user = useRecoilValue(UserState);
  const [userState, setUserState] = React.useState(false);
  const {mutate} = useSWRConfig();
  const {data, error} = useSWR<AccountUserList, ErrorType>(
    '/v2/account/list',
    accountUserFeather
  );

  React.useEffect(() => {
    // ログイン状態 -> ログアウトしたときのみデータをパージする
    if (user === null) {
      if (userState) {
        mutate(
          key => typeof key === 'string' && key.startsWith('/v2/account/list'),
          undefined,
          {revalidate: true}
        );
      }
    } else if (user) {
      setUserState(true);
    }
  }, [user]);

  const {switch: s} = useSwitchAccount();

  if (error) {
    return <Error {...error} />;
  }

  if (!data) {
    return (
      <Box overflowY="auto" maxH="calc(100% - 150px)">
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
    <Box overflowY="auto" maxH="calc(100% - 150px)">
      <Table variant="simple">
        <Tbody>
          {data.map(account => {
            const isCurrentUser = account.id === user?.user.id;

            return (
              <Tr
                _hover={isCurrentUser ? {} : {bgColor: hoverColor}}
                w="100%"
                key={account.id}
                onClick={() => {
                  if (!isCurrentUser) {
                    s(account.id, account.user_name);
                  }
                }}
                cursor="pointer"
              >
                <Td>
                  <Avatar src={account.avatar ?? ''} />
                </Td>
                <Td w="100%">
                  <Text
                    fontSize="1.5rem"
                    fontWeight="bold"
                    textAlign="center"
                    maxW={{base: '200px', sm: '240px'}}
                    textOverflow="ellipsis"
                    overflow="hidden"
                    whiteSpace="nowrap"
                    h="2rem"
                  >
                    @{account.user_name}
                  </Text>
                </Td>
                {isCurrentUser ? (
                  <Td p="0" pr="1rem">
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
      <Center mt="1rem" mb="2rem">
        <Button
          variant="link"
          as={Link}
          href="/login"
          leftIcon={<TbUserPlus size="24px" />}
        >
          アカウントを追加
        </Button>
      </Center>
    </Box>
  );
};
