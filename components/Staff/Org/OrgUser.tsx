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
} from '@chakra-ui/react';
import React from 'react';
import {useSWRConfig} from 'swr';
import {badgeColor} from '../../../utils/color';
import {OrganizationUser} from '../../../utils/types/organization';
import {Avatar} from '../../Common/Chakra/Avatar';
import {OrgJoinUser} from '../../Common/Form/OrgJoinUser';
import {Link as NextLink} from '../../Common/Next/Link';
import {useSecondaryColor} from '../../Common/useColor';
import {OrgDeleteUser} from './OrgDeleteUser';

interface Props {
  orgId: string;
  users: OrganizationUser[];
}

export const OrgUser: React.FC<Props> = ({users, orgId}) => {
  const textColor = useSecondaryColor();

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
      <OrgJoinUser
        orgId={orgId}
        handleSuccess={purge}
        apiEndpoint="/admin/org/member"
        defaultRole="owner"
      />
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
                    <Badge colorScheme={badgeColor(user.role)}>
                      {user.role}
                    </Badge>
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
