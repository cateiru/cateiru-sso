'use client';

import {
  IconButton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import Link from 'next/link';
import {TbUser} from 'react-icons/tb';
import useSWR from 'swr';
import {formatDate} from '../../../utils/date';
import {staffUserNameFetcher} from '../../../utils/swr/staff';
import {ErrorType} from '../../../utils/types/error';
import {UserNames} from '../../../utils/types/staff';
import {Error} from '../../Common/Error/Error';
import {useSecondaryColor} from '../../Common/useColor';

export const UserNameTable = () => {
  const textColor = useSecondaryColor();

  const {data, error} = useSWR<UserNames, ErrorType>(
    '/v2/admin/user_name',
    staffUserNameFetcher
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
            <Th>ユーザー名</Th>
            <Th>有効期限</Th>
            <Th>リンクされているユーザー</Th>
          </Tr>
        </Thead>
        <Tbody>
          {data.map(v => {
            return (
              <Tr key={`user-name-${v.id}`}>
                <Td fontWeight="bold">{v.user_name}</Td>
                <Td color={textColor}>{formatDate(new Date(v.period))}</Td>
                <Td>
                  <IconButton
                    as={Link}
                    href={`/staff/user/${v.user_id}`}
                    size="sm"
                    colorScheme="cateiru"
                    icon={<TbUser size="23px" />}
                    aria-label="リンクされているユーザー"
                  />
                </Td>
              </Tr>
            );
          })}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
