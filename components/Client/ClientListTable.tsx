import {
  Button,
  Center,
  Skeleton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import Link from 'next/link';
import React from 'react';
import type {ClientList as ClientListType} from '../../utils/types/client';
import {Avatar} from '../Common/Chakra/Avatar';
import {AgoTime} from '../Common/Time';

interface Props {
  clients?: ClientListType;
}

export const ClientListTable: React.FC<Props> = ({clients}) => {
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
          {clients
            ? clients.map(v => {
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
                      <AgoTime time={v.created_at} />
                    </Td>
                    <Td>
                      <Button
                        size="sm"
                        colorScheme="cateiru"
                        as={Link}
                        href={`/client/${v.client_id}`}
                      >
                        詳細
                      </Button>
                    </Td>
                  </Tr>
                );
              })
            : Array(5)
                .fill(0)
                .map((_, i) => {
                  return (
                    <Tr key={`loading-client-${i}`}>
                      <Td>
                        <Center>
                          <Skeleton w="32px" h="32px" borderRadius="50%" />
                        </Center>
                      </Td>
                      <Td>
                        <Skeleton w="100px" h="16px" />
                      </Td>
                      <Td>
                        <Skeleton w="200px" h="16px" />
                      </Td>
                      <Td>
                        <Skeleton w="100px" h="16px" />
                      </Td>
                      <Td>
                        <Skeleton w="52px" h="32px" borderRadius="0.375rem" />
                      </Td>
                    </Tr>
                  );
                })}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
