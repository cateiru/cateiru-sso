import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  Center,
  Avatar,
  Link,
  useToast,
  Heading,
  Box,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {getAllUsers} from '../../utils/api/admin';
import {UserInfo} from '../../utils/state/types';

const AdminPage = () => {
  const [users, setUsers] = React.useState<UserInfo[]>([]);
  const toast = useToast();

  React.useEffect(() => {
    const f = async () => {
      try {
        const users = await getAllUsers();
        setUsers(users);
      } catch (error) {
        if (error instanceof Error) {
          toast({
            title: error.message,
            status: 'error',
            isClosable: true,
            duration: 9000,
          });
        }
      }
    };

    f();
  }, []);

  const user = (user: UserInfo) => {
    return (
      <Tr key={user.user_id}>
        <Td>
          <Center>
            <Avatar src={user.avatar_url} size="sm" />
          </Center>
        </Td>
        <Td textAlign="center">
          <NextLink href={`/admin/user?id=${user.user_id}`} passHref>
            <Link>{user.user_name}</Link>
          </NextLink>
        </Td>
        <Td textAlign="center">{user.mail}</Td>
        <Td textAlign="center">{`${user.last_name} ${user.first_name}`}</Td>
      </Tr>
    );
  };

  return (
    <Center>
      <Box width={{base: '100%', lg: '1000px'}} mt="2rem">
        <Heading textAlign="center">
          すべてのユーザ（全{users.length}人）
        </Heading>
        <Box mx=".5rem" overflowX={{base: 'auto', lg: 'visible'}} mt="1rem">
          <Table
            variant="striped"
            minWidth="calc(1000px - 1rem)"
            size="lg"
            alignItems="center"
          >
            <Thead>
              <Tr>
                <Th></Th>
                <Th textAlign="center">ユーザ名</Th>
                <Th textAlign="center">メールアドレス</Th>
                <Th textAlign="center">名前</Th>
              </Tr>
            </Thead>
            <Tbody>{users.map(value => user(value))}</Tbody>
          </Table>
        </Box>
      </Box>
    </Center>
  );
};

export default AdminPage;
