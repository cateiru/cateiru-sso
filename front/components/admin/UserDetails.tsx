import {
  useToast,
  Box,
  Stack,
  Center,
  Heading,
  Table,
  Thead,
  Tbody,
  Flex,
  Tr,
  Th,
  Td,
  Badge,
  Input,
  Button,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
} from '@chakra-ui/react';
import Link from 'next/link';
import {useRouter} from 'next/router';
import React from 'react';
import {IoArrowBackOutline} from 'react-icons/io5';
import {useRecoilValue, useSetRecoilState} from 'recoil';
import {deleteUser, getUsers, role} from '../../utils/api/admin';
import {UserState, LoadState} from '../../utils/state/atom';
import {UserInfo} from '../../utils/state/types';
import Avatar from '../common/Avatar';

const selectRoleColor = (v: string) => {
  switch (v) {
    case 'user':
      return 'blue';
    case 'pro':
      return 'yellow';
    case 'admin':
      return 'red';
    default:
      return 'gray';
  }
};

const UserDetails = () => {
  const router = useRouter();
  const [user, setUser] = React.useState<UserInfo>();
  const toast = useToast();
  const [editRole, setEditRole] = React.useState('');
  const thisUser = useRecoilValue(UserState);
  const {isOpen, onOpen, onClose} = useDisclosure();
  const setLoad = useSetRecoilState(LoadState);

  React.useEffect(() => {
    if (!router.isReady) return;
    const query = router.query;

    if (typeof query['id'] === 'string') {
      const id = query['id'];
      const f = async () => {
        setLoad(true);
        try {
          const users = await getUsers(id);
          setUser(users[0]);
          setLoad(false);
        } catch (error) {
          setLoad(false);
          if (error instanceof Error) {
            toast({
              title: error.message,
              status: 'error',
              isClosable: true,
              duration: 9000,
            });
          }
          router.replace('/admin');
        }
      };
      f();
    }
  }, [router.isReady, router.query]);

  const submitRole = (enable: boolean) => {
    const f = async () => {
      try {
        await role(enable, editRole, user?.user_id || '');

        setEditRole('');
        setUser(v => {
          if (typeof v === 'undefined') {
            return undefined;
          }

          const roles = (v.role && v.role.concat()) || [];
          if (enable) {
            roles.push(editRole);
          } else {
            const index = roles.indexOf(editRole);
            if (index !== -1) {
              roles.splice(index, 1);
            }
          }

          return {
            ...v,
            role: roles,
          };
        });

        toast({
          title: enable ? '追加しました' : '削除しました',
          status: 'info',
          isClosable: true,
          duration: 9000,
        });
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

    if (editRole) {
      f();
    }
  };

  const submitDelete = () => {
    const f = async () => {
      try {
        await deleteUser(user?.user_id || '');
        onClose();

        router.replace('/admin');
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
  };

  return (
    <Box>
      <Box mx={{base: '5%', md: '10%'}} mt="1rem">
        <Link href="/admin" passHref>
          <Button
            pl=".5rem"
            variant="ghost"
            leftIcon={<IoArrowBackOutline size="25px" />}
          >
            戻る
          </Button>
        </Link>
      </Box>
      <Heading mb="1rem" textAlign="center">
        ユーザ詳細
      </Heading>
      <Stack direction={{base: 'column', lg: 'row'}} spacing="20px">
        <Center width={{base: '100%', lg: '80%'}} mt="2.3rem" mb="1rem">
          <Box
            width={{base: '100px', lg: '140px'}}
            backgroundColor="white"
            borderRadius="256"
          >
            <Avatar src={user?.avatar_url} size="full" isShadow />
          </Box>
        </Center>
        <Box width="100%">
          <Box overflow="auto">
            <Table variant="striped" width="600px" minWidth="600px">
              <Thead>
                <Tr>
                  <Th></Th>
                  <Th></Th>
                </Tr>
              </Thead>
              <Tbody>
                <Tr>
                  <Td fontWeight="bold" whiteSpace="nowrap">
                    ユーザ名
                  </Td>
                  <Td>
                    {user?.user_name} ({user?.user_name_formatted})
                  </Td>
                </Tr>
                <Tr>
                  <Td fontWeight="bold" whiteSpace="nowrap">
                    メールアドレス
                  </Td>
                  <Td>{user?.mail}</Td>
                </Tr>
                <Tr>
                  <Td fontWeight="bold" whiteSpace="nowrap">
                    ユーザID
                  </Td>
                  <Td>{user?.user_id}</Td>
                </Tr>
                <Tr>
                  <Td fontWeight="bold" whiteSpace="nowrap">
                    名前
                  </Td>
                  <Td>
                    {user?.last_name} {user?.first_name}
                  </Td>
                </Tr>
                <Tr>
                  <Td fontWeight="bold" whiteSpace="nowrap">
                    テーマ
                  </Td>
                  <Td>{user?.theme || 'null'}</Td>
                </Tr>
                <Tr>
                  <Td fontWeight="bold" whiteSpace="nowrap">
                    アバター
                  </Td>
                  <Td>{user?.avatar_url || 'null'}</Td>
                </Tr>
                <Tr>
                  <Td fontWeight="bold" whiteSpace="nowrap">
                    権限
                  </Td>
                  <Td>
                    {user?.role
                      ? user?.role.map(v => (
                          <Badge
                            colorScheme={selectRoleColor(v)}
                            key={v}
                            mr=".5rem"
                            variant="solid"
                          >
                            {v}
                          </Badge>
                        ))
                      : 'null'}
                  </Td>
                </Tr>
              </Tbody>
            </Table>
          </Box>

          {thisUser?.user_id === user?.user_id || (
            <Box mt="1rem">
              <Flex width={{base: '100%', sm: '300px'}}>
                <Input
                  placeholder="権限追加"
                  mr=".5rem"
                  onChange={v => setEditRole(v.target.value)}
                  value={editRole}
                />
                <Button onClick={() => submitRole(true)}>適用</Button>
                <Button
                  onClick={() => submitRole(false)}
                  colorScheme="red"
                  variant="ghost"
                >
                  削除
                </Button>
              </Flex>
              <Button
                mt="1rem"
                colorScheme="red"
                variant="ghost"
                onClick={onOpen}
              >
                アカウント削除
              </Button>
            </Box>
          )}
        </Box>
      </Stack>
      <Modal isOpen={isOpen} onClose={onClose} isCentered>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>削除？</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody>削除？</ModalBody>

          <ModalFooter>
            <Button colorScheme="red" mr={3} onClick={submitDelete}>
              削除
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Box>
  );
};

export default UserDetails;
