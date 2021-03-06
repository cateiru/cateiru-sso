import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  Center,
  Link,
  useToast,
  Heading,
  Box,
  Button,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {IoArrowBackOutline} from 'react-icons/io5';
import {useSetRecoilState} from 'recoil';
import {getAllUsers} from '../../utils/api/admin';
import {LoadState} from '../../utils/state/atom';
import {UserInfo} from '../../utils/state/types';
import Avatar from '../common/Avatar';

const AdminPage = () => {
  const [users, setUsers] = React.useState<UserInfo[]>([]);
  const toast = useToast();
  const setLoad = useSetRecoilState(LoadState);

  React.useEffect(() => {
    const f = async () => {
      setLoad(true);
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
      setLoad(false);
    };

    f();
  }, []);

  const user = (user: UserInfo) => {
    return (
      <Tr key={user.user_id}>
        <Td>
          <Center>
            <Avatar
              src={user.avatar_url}
              size="sm"
              borderColor={
                user.role.includes('admin') ? 'red.400' : 'yellow.400'
              }
              borderWidth={user.role.length !== 1 ? '2px' : '0'}
            />
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
        <Box mx=".5rem">
          <NextLink href="/setting/account" passHref>
            <Button
              pl=".5rem"
              variant="ghost"
              leftIcon={<IoArrowBackOutline size="25px" />}
            >
              ??????
            </Button>
          </NextLink>
        </Box>
        <Heading textAlign="center">
          ???????????????????????????{users.length}??????
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
                <Th textAlign="center">????????????</Th>
                <Th textAlign="center">?????????????????????</Th>
                <Th textAlign="center">??????</Th>
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
