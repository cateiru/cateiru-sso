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
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import useSWR, {useSWRConfig} from 'swr';
import {allowUserFetcher} from '../../utils/swr/client';
import {Card} from '../Common/Card';
import {Avatar} from '../Common/Chakra/Avatar';
import {Confirm} from '../Common/Confirm/Confirm';
import {Error} from '../Common/Error/Error';
import {useRequest} from '../Common/useRequest';

export const ClientAllowUserTable: React.FC<{id: string | string[]}> = ({
  id,
}) => {
  const {mutate} = useSWRConfig();
  const {data, error} = useSWR(`/v2/client/allow_user?client_id=${id}`, () =>
    allowUserFetcher(id)
  );
  const {request} = useRequest('/v2/client/allow_user');

  const deleteModal = useDisclosure();
  const [deleteId, setDeleteId] = React.useState<number | undefined>(undefined);

  const onDelete = async () => {
    if (!deleteId) return;

    const param = new URLSearchParams();
    param.append('id', deleteId.toString());

    const res = await request(
      {
        method: 'DELETE',
        mode: 'cors',
        credentials: 'include',
      },
      param
    );

    if (res) {
      // パージする
      mutate(
        key =>
          typeof key === 'string' &&
          key.startsWith(`/v2/client/allow_user?client_id=${id}`),
        undefined,
        {revalidate: true}
      );
    }
  };

  if (error) {
    return <Error {...error} />;
  }

  return (
    <>
      <Card title="ユーザー">
        <TableContainer>
          <Table variant="simple">
            <Thead>
              <Tr>
                <Th></Th>
                <Th>ユーザーID</Th>
                <Th></Th>
              </Tr>
            </Thead>
            <Tbody>
              {data
                ? data.map(v => {
                    if (v.user) {
                      return (
                        <Tr key={`allow-user-${v.id}`}>
                          <Td>
                            <Center>
                              <Avatar src={v.user.avatar ?? ''} size="sm" />
                            </Center>
                          </Td>
                          <Td>{v.user.user_name}</Td>
                          <Td>
                            <Button
                              size="sm"
                              colorScheme="cateiru"
                              onClick={() => {
                                setDeleteId(v.id);
                                deleteModal.onOpen();
                              }}
                            >
                              削除
                            </Button>
                          </Td>
                        </Tr>
                      );
                    }
                    return undefined;
                  })
                : Array(2)
                    .fill(0)
                    .map((_, i) => {
                      return (
                        <Tr key={`loading-allow-user-${i}`}>
                          <Td>
                            <Center>
                              <Skeleton w="32px" h="32px" borderRadius="50%" />
                            </Center>
                          </Td>
                          <Td>
                            <Skeleton w="100px" h="16px" />
                          </Td>
                          <Td>
                            <Skeleton
                              w="52px"
                              h="32px"
                              borderRadius="0.375rem"
                            />
                          </Td>
                        </Tr>
                      );
                    })}
            </Tbody>
          </Table>
        </TableContainer>
      </Card>
      <Card title="メールドメイン">
        <TableContainer>
          <Table variant="simple">
            <Thead>
              <Tr>
                <Th>メールドメイン</Th>
                <Th></Th>
              </Tr>
            </Thead>
            <Tbody>
              {data
                ? data.map(v => {
                    if (v.email_domain) {
                      return (
                        <Tr key={`allow-user-${v.id}`}>
                          <Td>{v.email_domain}</Td>
                          <Td>
                            <Button
                              size="sm"
                              colorScheme="cateiru"
                              onClick={() => {
                                setDeleteId(v.id);
                                deleteModal.onOpen();
                              }}
                            >
                              削除
                            </Button>
                          </Td>
                        </Tr>
                      );
                    }
                    return undefined;
                  })
                : Array(2)
                    .fill(0)
                    .map((_, i) => {
                      return (
                        <Tr key={`loading-allow-user-${i}`}>
                          <Td>
                            <Skeleton w="100px" h="16px" />
                          </Td>
                          <Td>
                            <Skeleton
                              w="52px"
                              h="32px"
                              borderRadius="0.375rem"
                            />
                          </Td>
                        </Tr>
                      );
                    })}
            </Tbody>
          </Table>
        </TableContainer>
      </Card>
      <Confirm
        onSubmit={onDelete}
        isOpen={deleteModal.isOpen}
        onClose={deleteModal.onClose}
        text={{
          confirmHeader: 'ルールを削除しますか？',
          confirmOkText: '削除',
          confirmOkTextColor: 'red',
        }}
      />
    </>
  );
};
