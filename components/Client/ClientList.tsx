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
import React from 'react';
import {hawManyDaysAgo} from '../../utils/date';
import type {ClientList as ClientListType} from '../../utils/types/client';
import {Avatar} from '../Common/Chakra/Avatar';
import {Tooltip} from '../Common/Chakra/Tooltip';

interface Props {
  clients: ClientListType;
}

export const ClientList: React.FC<Props> = ({clients}) => {
  return (
    <TableContainer>
      <Table variant="simple">
        <Thead>
          <Tr>
            <Th></Th>
            <Th>クライアント名</Th>
            <Th>説明</Th>
            <Th>作成日</Th>
            <Th></Th>
          </Tr>
        </Thead>
        <Tbody>
          {clients.map(v => {
            const created = new Date(v.created_at);

            return (
              <Tr key={`client-list-${v.client_id}`}>
                <Td>
                  <Center>
                    <Avatar src={v.image ?? ''} size="sm" />
                  </Center>
                </Td>
                <Td>{v.name}</Td>
                <Td
                  maxW="200px"
                  textOverflow="ellipsis"
                  whiteSpace="nowrap"
                  overflowX="hidden"
                >
                  {v.description}
                </Td>
                <Td>
                  <Tooltip placement="top" label={created.toLocaleString()}>
                    {hawManyDaysAgo(created)}
                  </Tooltip>
                </Td>
                <Td>
                  <Button size="sm" colorScheme="cateiru">
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
