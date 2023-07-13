import {
  Badge,
  Button,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
  useColorModeValue,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import {useSWRConfig} from 'swr';
import {OrganizationUser} from '../../../utils/types/organization';
import {Avatar} from '../../Common/Chakra/Avatar';
import {OrgDeleteUser} from './OrgDeleteUser';
import {OrgJoinUser} from './OrgJoinUser';

interface Props {
  orgId: string;
  users: OrganizationUser[];
}

export const OrgUser: React.FC<Props> = ({users, orgId}) => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  const {mutate} = useSWRConfig();

  const purge = () => {
    mutate(
      key =>
        typeof key === 'string' &&
        key.startsWith(`/v2/admin/org?org_id=${orgId}`),
      undefined,
      {revalidate: true}
    );
  };

  return (
    <>
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        組織ユーザー
      </Text>
      <OrgJoinUser orgId={orgId} handleSuccess={purge} />
      <TableContainer mt="1rem">
        <Table variant="simple">
          <Thead>
            <Tr>
              <Td></Td>
              <Th></Th>
              <Th>ユーザー名</Th>
              <Th textAlign="center">ロール</Th>
              <Th>参加日</Th>
              <Th></Th>
            </Tr>
          </Thead>
          <Tbody>
            {users.map(user => {
              const joinDate = new Date(user.created_at);

              const badgeColor = () => {
                if (user.role === 'owner') {
                  return 'red';
                }

                if (user.role === 'member') {
                  return 'blue';
                }

                return 'gray';
              };

              return (
                <Tr key={`org-user-${user.id}`}>
                  <Td>
                    <OrgDeleteUser
                      userId={user.id}
                      userName={user.user.user_name}
                      handleSuccess={purge}
                    />
                  </Td>
                  <Td>
                    <Avatar src={user.user.avatar ?? ''} size="sm" />
                  </Td>
                  <Td>{user.user.user_name}</Td>
                  <Td textAlign="center">
                    <Badge colorScheme={badgeColor()}>{user.role}</Badge>
                  </Td>
                  <Td>{joinDate.toLocaleDateString()}</Td>
                  <Td>
                    <Button
                      colorScheme="cateiru"
                      size="sm"
                      as={NextLink}
                      href={`/staff/user/${user.user.id}`}
                    >
                      ユーザー詳細
                    </Button>
                  </Td>
                </Tr>
              );
            })}
          </Tbody>
        </Table>
      </TableContainer>
    </>
  );
};