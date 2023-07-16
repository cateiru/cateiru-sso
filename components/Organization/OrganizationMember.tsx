import {
  Badge,
  Box,
  Button,
  Center,
  FormControl,
  FormErrorMessage,
  IconButton,
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
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {TbEdit} from 'react-icons/tb';
import {useRecoilValue} from 'recoil';
import useSWR, {useSWRConfig} from 'swr';
import {badgeColor} from '../../utils/color';
import {UserState} from '../../utils/state/atom';
import {orgUsersFeather} from '../../utils/swr/featcher';
import {ErrorType} from '../../utils/types/error';
import {
  OrganizationUser,
  OrganizationUserList,
} from '../../utils/types/organization';
import {Avatar} from '../Common/Chakra/Avatar';
import {Error} from '../Common/Error/Error';
import {OrgJoinUser} from '../Common/Form/OrgJoinUser';
import {Spinner} from '../Common/Icons/Spinner';
import {useRequest} from '../Common/useRequest';

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
  const {request} = useRequest('/v2/org/member');

  const {mutate} = useSWRConfig();
  const {isOpen, onOpen, onClose} = useDisclosure();
  const [modalUser, setModalUser] = React.useState<OrganizationUser | null>(
    null
  );
  const [loading, setLoading] = React.useState(false);

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
      onClose();
      reload();
    }
  };

  const onDelete = async () => {
    setLoading(true);
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
      onClose();
      reload();
    }

    setLoading(false);
  };

  const reload = () => {
    mutate(
      key =>
        typeof key === 'string' &&
        key.startsWith(`/v2/org/member?org_id=${id}`),
      undefined,
      {revalidate: true}
    );
  };

  return (
    <Box>
      <OrgJoinUser
        apiEndpoint="/v2/org/member"
        orgId={id}
        handleSuccess={reload}
      />
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
                const joinDate = new Date(user.created_at);

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
                    <Td>{joinDate.toLocaleDateString()}</Td>
                    <Td>
                      <IconButton
                        size="sm"
                        colorScheme="cateiru"
                        icon={<TbEdit size="20px" />}
                        aria-label="edit user"
                        onClick={() => {
                          setModalUser(user);
                          onOpen();
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
      <Modal isOpen={isOpen} onClose={onClose} isCentered>
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
            <Button w="100%" onClick={onDelete} isLoading={loading}>
              組織から外す
            </Button>
          </ModalBody>
        </ModalContent>
      </Modal>
    </Box>
  );
};
