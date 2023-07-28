import {
  Badge,
  Box,
  Button,
  Center,
  FormControl,
  FormErrorMessage,
  IconButton,
  ListItem,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Select,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
  UnorderedList,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {TbEdit, TbTrash} from 'react-icons/tb';
import {useRecoilValue} from 'recoil';
import useSWR, {useSWRConfig} from 'swr';
import {badgeColor} from '../../utils/color';
import {UserState} from '../../utils/state/atom';
import {
  orgInviteMemberListFeather,
  orgUsersFeather,
} from '../../utils/swr/organization';
import {ErrorType} from '../../utils/types/error';
import {
  OrganizationInviteMember,
  OrganizationInviteMemberList,
  OrganizationUser,
  OrganizationUserList,
} from '../../utils/types/organization';
import {Card} from '../Common/Card';
import {Avatar} from '../Common/Chakra/Avatar';
import {Confirm} from '../Common/Confirm/Confirm';
import {Error} from '../Common/Error/Error';
import {Spinner} from '../Common/Icons/Spinner';
import {AgoTime} from '../Common/Time';
import {useRequest} from '../Common/useRequest';
import {JoinOrganization} from './JoinOrganization';

interface Props {
  id: string;
}

export interface EditRoleForm {
  role: string;
}

export const OrganizationMember: React.FC<Props> = ({id}) => {
  const u = useRecoilValue(UserState);

  const {data, error} = useSWR<OrganizationUserList, ErrorType>(
    `/v2/org/member?org_id=${id}`,
    () => orgUsersFeather(id)
  );
  const {data: InviteData, error: InviteError} = useSWR<
    OrganizationInviteMemberList,
    ErrorType
  >(`/v2/org/member/invite?org_id=${id}`, () => orgInviteMemberListFeather(id));

  const {request} = useRequest('/v2/org/member');
  const {request: requestJoin} = useRequest('/v2/org/member/invite');

  const {mutate} = useSWRConfig();
  const editMemberModal = useDisclosure();
  const deleteJoinModal = useDisclosure();
  const confirmModal = useDisclosure();

  const [modalUser, setModalUser] = React.useState<OrganizationUser | null>(
    null
  );
  const [joinItem, setJoinItem] =
    React.useState<OrganizationInviteMember | null>(null);

  const {
    handleSubmit,
    register,
    reset,
    setValue,
    formState: {isSubmitting, errors},
  } = useForm<EditRoleForm>();

  React.useEffect(() => {
    if (modalUser) {
      setValue('role', modalUser.role);
    }
  }, [modalUser]);

  if (error) {
    return <Error {...error} />;
  } else if (InviteError) {
    return <Error {...InviteError} />;
  }

  const onSubmit = async (data: EditRoleForm) => {
    const form = new FormData();
    form.append('org_user_id', String(modalUser?.id));
    form.append('role', data.role);

    const res = await request({
      method: 'PUT',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      reset();
      editMemberModal.onClose();
      reload();
    }
  };

  const onDelete = async () => {
    const param = new URLSearchParams();
    param.append('org_user_id', String(modalUser?.id));

    const res = await request(
      {
        method: 'DELETE',
        mode: 'cors',
        credentials: 'include',
      },
      param
    );

    if (res) {
      reset();
      editMemberModal.onClose();
      reload();
    }
  };

  const onJoinDelete = async () => {
    if (!joinItem) return;

    const param = new URLSearchParams();
    param.append('invite_id', String(joinItem?.id));

    const res = await requestJoin(
      {
        method: 'DELETE',
        mode: 'cors',
        credentials: 'include',
      },
      param
    );

    if (res) {
      deleteJoinModal.onClose();
      reload();
    }
  };

  const reload = () => {
    mutate(
      key =>
        typeof key === 'string' &&
        (key.startsWith(`/v2/org/member?org_id=${id}`) ||
          key.startsWith(`/v2/org/member/invite?org_id=${id}`)),
      undefined,
      {revalidate: true}
    );
  };

  return (
    <Box>
      <JoinOrganization orgId={id} handleSuccess={reload} />
      <Card title="組織ユーザー">
        {data ? (
          <TableContainer mt="1rem">
            <Table variant="simple">
              <Thead>
                <Tr>
                  <Th></Th>
                  <Th>ユーザー名</Th>
                  <Th textAlign="center">ロール</Th>
                  <Th>追加日</Th>
                  <Th></Th>
                </Tr>
              </Thead>
              <Tbody>
                {data.map(user => {
                  return (
                    <Tr key={`org-user-${user.id}`}>
                      <Td>
                        <Avatar src={user.user.avatar ?? ''} size="sm" />
                      </Td>
                      <Td>{user.user.user_name}</Td>
                      <Td textAlign="center">
                        <Badge colorScheme={badgeColor(user.role)}>
                          {user.role}
                        </Badge>
                      </Td>
                      <Td>
                        <AgoTime time={user.created_at} />
                      </Td>
                      <Td>
                        <IconButton
                          size="sm"
                          colorScheme="cateiru"
                          icon={<TbEdit size="20px" />}
                          aria-label="edit user"
                          onClick={() => {
                            setModalUser(user);
                            editMemberModal.onOpen();
                          }}
                          isDisabled={user.user.id === u?.user.id}
                        />
                      </Td>
                    </Tr>
                  );
                })}
              </Tbody>
            </Table>
          </TableContainer>
        ) : (
          <Center mt="2rem">
            <Spinner />
          </Center>
        )}
      </Card>
      <Card title="招待中のユーザー">
        {InviteData ? (
          <TableContainer mt="1rem">
            <Table variant="simple">
              <Thead>
                <Tr>
                  <Th>Eメール</Th>
                  <Th>追加日</Th>
                  <Th></Th>
                </Tr>
              </Thead>
              <Tbody>
                {InviteData.map(item => {
                  return (
                    <Tr key={`invite-data-${item.id}`}>
                      <Td>{item.email}</Td>
                      <Td>
                        <AgoTime time={item.created_at} />
                      </Td>
                      <Td>
                        <IconButton
                          size="sm"
                          colorScheme="cateiru"
                          icon={<TbTrash size="20px" />}
                          aria-label="edit user"
                          onClick={() => {
                            setJoinItem(item);
                            deleteJoinModal.onOpen();
                          }}
                        />
                      </Td>
                    </Tr>
                  );
                })}
              </Tbody>
            </Table>
          </TableContainer>
        ) : (
          <Center mt="2rem">
            <Spinner />
          </Center>
        )}
      </Card>
      <Modal
        isOpen={editMemberModal.isOpen}
        onClose={editMemberModal.onClose}
        isCentered
      >
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>{modalUser?.user.user_name} の編集</ModalHeader>
          <ModalCloseButton size="lg" />
          <ModalBody mb="1rem">
            <form onSubmit={handleSubmit(onSubmit)}>
              <FormControl isInvalid={!!errors.role}>
                <Select
                  {...register('role', {
                    required: 'ロールは必須です',
                  })}
                >
                  <option value="owner">管理者</option>
                  <option value="member">メンバー</option>
                  <option value="guest">ゲスト</option>
                </Select>
                <FormErrorMessage>
                  {errors.role && errors.role.message}
                </FormErrorMessage>
              </FormControl>
              <Button
                mt=".5rem"
                w="100%"
                colorScheme="cateiru"
                type="submit"
                isLoading={isSubmitting}
              >
                権限を変更する
              </Button>
            </form>
            <Text my="1rem" textAlign="center">
              もしくは、
            </Text>
            <Button
              w="100%"
              onClick={() => {
                editMemberModal.onClose();
                confirmModal.onOpen();
              }}
            >
              組織から外す
            </Button>
          </ModalBody>
        </ModalContent>
      </Modal>
      <Confirm
        isOpen={confirmModal.isOpen}
        onClose={confirmModal.onClose}
        onSubmit={onDelete}
        text={{
          confirmHeader: '組織からユーザーを削除しますか？',
          confirmOkText: '削除',
          confirmOkTextColor: 'red',
        }}
      >
        <UnorderedList spacing=".5rem">
          <ListItem>
            組織からユーザーを削除するとそのユーザーはこの組織にアクセスすることができなくなります。
          </ListItem>
          <ListItem>
            再度アクセスさせるには、組織のオーナーが招待する必要があります。
          </ListItem>
          <ListItem>
            ユーザーが作成したクライアントなどは削除されません。
          </ListItem>
        </UnorderedList>
      </Confirm>
      <Confirm
        isOpen={deleteJoinModal.isOpen}
        onClose={deleteJoinModal.onClose}
        onSubmit={onJoinDelete}
        text={{
          confirmHeader: '招待を削除しますか？',
          confirmOkText: '削除',
          confirmOkTextColor: 'red',
        }}
      >
        <UnorderedList spacing=".5rem">
          <ListItem>送信されたメールは削除されません。</ListItem>
          <ListItem>招待メールにある招待URLが無効化されます。</ListItem>
        </UnorderedList>
      </Confirm>
    </Box>
  );
};
