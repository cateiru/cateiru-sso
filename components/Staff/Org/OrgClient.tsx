import {
  Avatar,
  Button,
  Center,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import Link from 'next/link';
import React from 'react';
import {StaffClient} from '../../../utils/types/staff';
import {useSecondaryColor} from '../../Common/useColor';

interface Props {
  clients: StaffClient[];
}

export const OrgClient: React.FC<Props> = props => {
  const textColor = useSecondaryColor();

  return (
    <>
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        組織で作成したクライアント
      </Text>
      <TableContainer mt="1rem">
        <Table variant="simple">
          <Thead>
            <Tr>
              <Th></Th>
              <Th>クライアント名</Th>
              <Th></Th>
            </Tr>
          </Thead>
          <Tbody>
            {props.clients.map(client => {
              return (
                <Tr key={`client-${client.client_id}`}>
                  <Td>
                    <Center>
                      <Avatar src={client.image ?? ''} size="sm" />
                    </Center>
                  </Td>
                  <Td>{client.name}</Td>
                  <Td>
                    <Button
                      size="sm"
                      colorScheme="cateiru"
                      as={Link}
                      href={`/staff/client/${client.client_id}`}
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
    </>
  );
};
