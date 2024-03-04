'use client';

import {
  Button,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Link,
  Tr,
} from '@chakra-ui/react';
import useSWR from 'swr';
import {orgsFeather} from '../../../utils/swr/staff';
import {ErrorType} from '../../../utils/types/error';
import {Organizations} from '../../../utils/types/staff';
import {Avatar} from '../../Common/Chakra/Avatar';
import {Error} from '../../Common/Error/Error';
import {Link as NextLink} from '../../Common/Next/Link';
import {AgoTime} from '../../Common/Time';

export const OrgsTable = () => {
  const {data, error} = useSWR<Organizations, ErrorType>(
    '/admin/orgs',
    orgsFeather
  );

  if (error) {
    return <Error {...error} />;
  }

  if (!data) {
    return <></>;
  }

  return (
    <>
      <Button
        colorScheme="cateiru"
        mt=".5rem"
        as={NextLink}
        href="/staff/org/register"
        size="sm"
      >
        新規作成
      </Button>
      <TableContainer mt=".5rem">
        <Table variant="simple">
          <Thead>
            <Tr>
              <Th></Th>
              <Th>組織名</Th>
              <Th>作成日</Th>
              <Th></Th>
            </Tr>
          </Thead>
          <Tbody>
            {data.map((v, i) => {
              return (
                <Tr key={`brands-${i}`}>
                  <Td>
                    <Avatar src={v.image ?? ''} size="sm" />
                  </Td>
                  <Td>
                    {v.link ? (
                      <Link isExternal href={v.link}>
                        {v.name}
                      </Link>
                    ) : (
                      v.name
                    )}
                  </Td>
                  <Td>
                    <AgoTime time={v.created_at} />
                  </Td>
                  <Td>
                    <Button
                      as={NextLink}
                      href={`/staff/org/${v.id}`}
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
    </>
  );
};
