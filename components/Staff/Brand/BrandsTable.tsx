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
} from '@chakra-ui/react';
import useSWR from 'swr';
import {hawManyDaysAgo} from '../../../utils/date';
import {brandFeather} from '../../../utils/swr/featcher';
import {ErrorType} from '../../../utils/types/error';
import {Brands} from '../../../utils/types/staff';
import {Tooltip} from '../../Common/Chakra/Tooltip';
import {Error} from '../../Common/Error/Error';
import {Link} from '../../Common/Next/Link';

export const BrandsTable = () => {
  const {data, error} = useSWR<Brands, ErrorType>('/v2/admin/brand', () =>
    brandFeather()
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
        as={Link}
        href="/staff/brand/register"
      >
        新規作成
      </Button>
      <TableContainer mt=".5rem">
        <Table variant="simple">
          <Thead>
            <Tr>
              <Th>ブランド名</Th>
              <Th>詳細</Th>
              <Th>作成日</Th>
              <Th></Th>
            </Tr>
          </Thead>
          <Tbody>
            {data.map((v, i) => {
              const created = new Date(v.created_at);

              return (
                <Tr key={`brands-${i}`}>
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
                    <Button
                      as={Link}
                      href={`/staff/brand/${v.id}`}
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
