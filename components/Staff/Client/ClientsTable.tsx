'use client';

import {
  Button,
  Center,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import useSWR from 'swr';
import {staffClientsFeather} from '../../../utils/swr/staff';
import {ErrorType} from '../../../utils/types/error';
import {StaffClients} from '../../../utils/types/staff';
import {Avatar} from '../../Common/Chakra/Avatar';
import {Error} from '../../Common/Error/Error';
import {Link} from '../../Common/Next/Link';

export const ClientsTable = () => {
  const {data, error} = useSWR<StaffClients, ErrorType>(
    '/admin/clients',
    staffClientsFeather
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
            <Th>クライアント名</Th>
            <Th></Th>
          </Tr>
        </Thead>
        <Tbody>
          {data.map(v => {
            return (
              <Tr key={`staff-client-${v.client_id}`}>
                <Td>
                  <Center>
                    <Avatar src={v.image ?? ''} size="sm" />
                  </Center>
                </Td>
                <Td>{v.name}</Td>
                <Td>
                  <Button
                    size="sm"
                    colorScheme="cateiru"
                    as={Link}
                    href={`/staff/client/${v.client_id}`}
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
