import {
  Badge,
  Box,
  Center,
  IconButton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import {TbEdit} from 'react-icons/tb';
import useSWR, {useSWRConfig} from 'swr';
import {badgeColor} from '../../utils/color';
import {orgUsersFeather} from '../../utils/swr/featcher';
import {ErrorType} from '../../utils/types/error';
import {OrganizationUserList} from '../../utils/types/organization';
import {Avatar} from '../Common/Chakra/Avatar';
import {Error} from '../Common/Error/Error';
import {OrgJoinUser} from '../Common/Form/OrgJoinUser';
import {Spinner} from '../Common/Icons/Spinner';

interface Props {
  id: string;
}

export const OrganizationMember: React.FC<Props> = ({id}) => {
  const {data, error} = useSWR<OrganizationUserList, ErrorType>(
    `/v2/org/member?org_id=${id}`,
    () => orgUsersFeather(id)
  );
  const {mutate} = useSWRConfig();

  if (error) {
    return <Error {...error} />;
  }

  return (
    <Box>
      <OrgJoinUser
        apiEndpoint="/v2/org/member"
        orgId={id}
        handleSuccess={() => {
          mutate(
            key =>
              typeof key === 'string' &&
              key.startsWith(`/v2/org/member?org_id=${id}`),
            undefined,
            {revalidate: true}
          );
        }}
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
