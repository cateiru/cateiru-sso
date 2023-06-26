'use client';

import {
  Button,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  useColorModeValue,
} from '@chakra-ui/react';
import Link from 'next/link';
import useSWR from 'swr';
import {staffUsersFeather} from '../../../utils/swr/featcher';
import {ErrorType} from '../../../utils/types/error';
import {StaffUsers} from '../../../utils/types/staff';
import {Avatar} from '../../Common/Chakra/Avatar';
import {Error} from '../../Common/Error/Error';

export const UsersTable = () => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  const {data, error} = useSWR<StaffUsers, ErrorType>(
    '/v2/admin/users',
    staffUsersFeather
  );

  if (error) {
    return <Error {...error} />;
  }

  if (!data) {
    return <></>;
  }

  return (
    <TableContainer mt="2rem">
      <Table variant="simple">
        <Thead>
          <Tr>
            <Th></Th>
            <Th>ユーザー名</Th>
            <Th>Email</Th>
            <Th></Th>
          </Tr>
        </Thead>
        <Tbody>
          {data.map((v, i) => {
            return (
              <Tr key={`staff-users-${i}`}>
                <Td>
                  <Avatar src={v.avatar ?? ''} size="sm" />
                </Td>
                <Td fontWeight="bold" color={textColor}>
                  @{v.user_name}
                </Td>
                <Td>{v.email}</Td>
                <Td>
                  <Button
                    as={Link}
                    href={`/staff/user/${v.id}`}
                    size="sm"
                    colorScheme="cateiru"
                  >
                    詳細
                  </Button>
                </Td>
              </Tr>
            );
          })}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
