'use client';

import {
  Badge,
  Button,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  useToast,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import {TbTrashX} from 'react-icons/tb';
import useSWR, {useSWRConfig} from 'swr';
import {staffRegisterSessionsFeather} from '../../../utils/swr/staff';
import {ErrorType} from '../../../utils/types/error';
import {RegisterSessions} from '../../../utils/types/staff';
import {DeleteButton} from '../../Common/DeleteButton';
import {Error} from '../../Common/Error/Error';
import {useRequest} from '../../Common/useRequest';

export const RegisterSessionTable = () => {
  const toast = useToast();
  const {mutate} = useSWRConfig();

  const {data, error} = useSWR<RegisterSessions, ErrorType>(
    '/admin/register_session',
    staffRegisterSessionsFeather
  );
  const {request} = useRequest('/admin/register_session');

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
            <Th>メールアドレス</Th>
            <Th>認証済みか</Th>
            <Th>メール送信数</Th>
            <Th>検証回数</Th>
            <Th>有効期限</Th>
            <Th>組織</Th>
          </Tr>
        </Thead>
        <Tbody>
          {data.map(v => {
            const onDelete = async () => {
              const param = new URLSearchParams();
              param.append('email', v.email);

              const res = await request(
                {
                  method: 'DELETE',
                  mode: 'cors',
                  credentials: 'include',
                },
                param
              );

              if (res) {
                toast({
                  title: '登録セッションを削除しました',
                  status: 'success',
                });

                mutate(
                  key =>
                    typeof key === 'string' &&
                    key.startsWith('/admin/register_session'),
                  undefined,
                  {revalidate: true}
                );
              }
            };

            return (
              <Tr key={`staff-client-${v.email}`}>
                <Td>
                  <DeleteButton
                    tooltipLabel={`${v.email}のセッションを削除`}
                    onSubmit={onDelete}
                    text={{
                      confirmHeader: `${v.email}のセッションを削除しますか？`,
                      confirmOkText: '削除',
                    }}
                    icon={<TbTrashX size="25px" />}
                  />
                </Td>
                <Td>{v.email}</Td>
                <Td>
                  {v.email_verified ? (
                    <Badge colorScheme="green" variant="subtle">
                      はい
                    </Badge>
                  ) : (
                    <Badge colorScheme="red" variant="subtle">
                      いいえ
                    </Badge>
                  )}
                </Td>
                <Td>{v.send_count}</Td>
                <Td>{v.retry_count}</Td>
                <Td>{new Date(v.period).toLocaleString()}</Td>
                <Td>
                  {v.org_id && (
                    <Button
                      as={NextLink}
                      size="sm"
                      href={`/staff/org/${v.org_id}`}
                    >
                      組織
                    </Button>
                  )}
                </Td>
              </Tr>
            );
          })}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
