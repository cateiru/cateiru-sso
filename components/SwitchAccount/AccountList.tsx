import {
  Table,
  Tbody,
  Tr,
  Td,
  Text,
  Box,
  useColorModeValue,
  Skeleton,
  Button,
  Center,
} from '@chakra-ui/react';
import {useAtomValue} from 'jotai';
import React from 'react';
import {TbCheck, TbUserPlus} from 'react-icons/tb';
import useSWR, {useSWRConfig} from 'swr';
import {UserState} from '../../utils/state/atom';
import {accountUserFeather} from '../../utils/swr/account';
import {AccountUserList} from '../../utils/types/account';
import {ErrorType} from '../../utils/types/error';
import {Avatar} from '../Common/Chakra/Avatar';
import {Tooltip} from '../Common/Chakra/Tooltip';
import {Error} from '../Common/Error/Error';
import {Spinner} from '../Common/Icons/Spinner';
import {Link} from '../Common/Next/Link';
import {useSwitchAccount} from './useSwitchAccount';

interface Props {
  isOauth: boolean;
}

export const AccountList: React.FC<Props> = ({isOauth}) => {
  const checkMarkColor = useColorModeValue('#68D391', '#38A169');
  const hoverColor = useColorModeValue('gray.100', 'gray.600');

  const user = useAtomValue(UserState);
  const [userState, setUserState] = React.useState(false);
  const {mutate} = useSWRConfig();
  const {data, error} = useSWR<AccountUserList, ErrorType>(
    '/account/list',
    accountUserFeather
  );
  const {switch: s, loading, redirect} = useSwitchAccount();

  React.useEffect(() => {
    // ログイン状態 -> ログアウトしたときのみデータをパージする
    if (user === null) {
      if (userState) {
        mutate(
          key => typeof key === 'string' && key.startsWith('/account/list'),
          undefined,
          {revalidate: true}
        );
      }
    } else if (user) {
      setUserState(true);
    }
  }, [user]);

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
                _hover={isCurrentUser && !isOauth ? {} : {bgColor: hoverColor}}
                w="100%"
                key={account.id}
                onClick={() => {
                  if (!isCurrentUser) {
                    s(account.id, account.user_name);
                  } else if (isOauth) {
                    redirect(account.id);
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
                <Td p="0" pr="1rem">
                  {loading !== null && loading === account.id ? (
                    <Spinner />
                  ) : (
                    isCurrentUser && (
                      <Tooltip label="現在ログインしているユーザーです">
                        <Box>
                          <TbCheck
                            size="30px"
                            color={checkMarkColor}
                            strokeWidth="3px"
                          />
                        </Box>
                      </Tooltip>
                    )
                  )}
                </Td>
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
