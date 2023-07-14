import {
  Badge,
  Box,
  Button,
  Center,
  Flex,
  FormControl,
  FormErrorMessage,
  IconButton,
  Input,
  Select,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import {useForm} from 'react-hook-form';
import {TbEdit} from 'react-icons/tb';
import useSWR, {useSWRConfig} from 'swr';
import {badgeColor} from '../../utils/color';
import {orgUsersFeather} from '../../utils/swr/featcher';
import {ErrorType} from '../../utils/types/error';
import {OrganizationUserList} from '../../utils/types/organization';
import {Avatar} from '../Common/Chakra/Avatar';
import {Error} from '../Common/Error/Error';
import {Spinner} from '../Common/Icons/Spinner';
import {useRequest} from '../Common/useRequest';

interface Props {
  id: string;
}

interface JoinFormData {
  user_name_or_email: string;
  role: string;
}

export const OrganizationMember: React.FC<Props> = ({id}) => {
  const {data, error} = useSWR<OrganizationUserList, ErrorType>(
    `/v2/org/member?org_id=${id}`,
    () => orgUsersFeather(id)
  );
  const {request} = useRequest('/v2/org/member');
  const {mutate} = useSWRConfig();

  const {
    handleSubmit,
    register,
    reset,
    formState: {isSubmitting, errors},
  } = useForm<JoinFormData>({
    defaultValues: {
      role: 'guest',
    },
  });

  const onSubmit = async (data: JoinFormData) => {
    const form = new FormData();

    form.append('org_id', id);
    form.append('user_name_or_email', data.user_name_or_email);
    form.append('role', data.role);

    const res = await request({
      method: 'POST',
      body: form,
      mode: 'cors',
      credentials: 'include',
    });

    if (res) {
      reset();

      mutate(
        key =>
          typeof key === 'string' &&
          key.startsWith(`/v2/org/member?org_id=${id}`),
        undefined,
        {revalidate: true}
      );
    }
  };

  if (error) {
    return <Error {...error} />;
  }

  return (
    <Box>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Flex>
          <FormControl isInvalid={!!errors.user_name_or_email}>
            <Input
              type="text"
              placeholder="ユーザー名またはメールアドレス"
              {...register('user_name_or_email', {
                required: 'ユーザーは必須です',
              })}
            />
            <FormErrorMessage>
              {errors.user_name_or_email && errors.user_name_or_email.message}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={!!errors.role} ml=".5rem">
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
        </Flex>

        <Button
          mt=".5rem"
          w="100%"
          colorScheme="cateiru"
          type="submit"
          isLoading={isSubmitting}
        >
          ユーザーを追加
        </Button>
      </form>
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
    </Box>
  );
};
