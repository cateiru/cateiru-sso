import {
  Button,
  Center,
  Skeleton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
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

  const users = React.useMemo(() => {
    if (typeof data === 'undefined') return undefined;

    return data.filter(v => typeof v.user !== 'undefined');
  }, [data]);

  const emailDomain = React.useMemo(() => {
    if (typeof data === 'undefined') return undefined;

    return data.filter(v => v.email_domain);
  }, [data]);

  if (error) {
    return <Error {...error} />;
  }

  return (
    <>
      <Card
        title="ユーザー"
        description="許可するユーザーを追加することでそのユーザーのみ認証を許可することができます。"
      >
        {(users && users.length) === 0 ? (
          <Text>
            許可するユーザーが設定されていないため、
            すべてのユーザーが許可されます。
          </Text>
        ) : (
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
                {users
                  ? users.map(v => {
                      return (
                        <Tr key={`allow-user-${v.id}`}>
                          <Td>
                            <Center>
                              <Avatar src={v.user?.avatar ?? ''} size="sm" />
                            </Center>
                          </Td>
                          <Td>{v.user?.user_name}</Td>
                          <Td>
                            <Center justifyContent="end">
                              <Button
                                size="sm"
                                onClick={() => {
                                  setDeleteId(v.id);
                                  deleteModal.onOpen();
                                }}
                              >
                                削除
                              </Button>
                            </Center>
                          </Td>
                        </Tr>
                      );
                    })
                  : Array(2)
                      .fill(0)
                      .map((_, i) => {
                        return (
                          <Tr key={`loading-allow-user-${i}`}>
                            <Td>
                              <Center>
                                <Skeleton
                                  w="32px"
                                  h="32px"
                                  borderRadius="50%"
                                />
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
        )}
      </Card>
      <Card
        title="メールドメイン"
        description="許可するメールドメインを追加することでそのメールドメインのメールアドレスをもつユーザーのみ認証を許可することができます。"
      >
        {(emailDomain && emailDomain?.length) === 0 ? (
          <Text>
            許可するメールドメインが設定されていないため、
            すべてのメールドメインが許可されます。
          </Text>
        ) : (
          <TableContainer>
            <Table variant="simple">
              <Thead>
                <Tr>
                  <Th>メールドメイン</Th>
                  <Th></Th>
                </Tr>
              </Thead>
              <Tbody>
                {emailDomain
                  ? emailDomain.map(v => {
                      if (v.email_domain) {
                        return (
                          <Tr key={`allow-user-${v.id}`}>
                            <Td>{v.email_domain}</Td>
                            <Td>
                              <Center justifyContent="end">
                                <Button
                                  size="sm"
                                  onClick={() => {
                                    setDeleteId(v.id);
                                    deleteModal.onOpen();
                                  }}
                                >
                                  削除
                                </Button>
                              </Center>
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
        )}
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
