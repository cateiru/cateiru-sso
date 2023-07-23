'use client';

import {
  Badge,
  Button,
  ButtonGroup,
  Center,
  Flex,
  Heading,
  Link,
  ListItem,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Tr,
  UnorderedList,
  useClipboard,
} from '@chakra-ui/react';
import {useParams} from 'next/navigation';
import {TbCheck} from 'react-icons/tb';
import useSWR from 'swr';
import {clientFetcher} from '../../utils/swr/client';
import {ClientDetail as ClientDetailType} from '../../utils/types/client';
import {ErrorType} from '../../utils/types/error';
import {Avatar} from '../Common/Chakra/Avatar';
import {Error} from '../Common/Error/Error';
import {Margin} from '../Common/Margin';

export const ClientDetail = () => {
  const {id} = useParams();

  const {data, error} = useSWR<ClientDetailType, ErrorType>(
    `/api/client/?client_id=${id}`,
    () => clientFetcher(id, undefined)
  );

  const clientIdCopy = useClipboard(data?.client_id ?? '');
  const clientSecretCopy = useClipboard(data?.client_secret ?? '');

  if (error) {
    return <Error {...error} />;
  }

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        クライアントの詳細
      </Heading>
      <Center my="1rem">
        <Avatar src={data?.image ?? ''} size="lg" />
      </Center>
      <TableContainer mt="1rem">
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td fontWeight="bold">クライアント名</Td>
              <Td>{data?.name}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">クライアントの説明</Td>
              <Td>
                <Text as="span">{data?.description}</Text>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">スコープ</Td>
              <Td>
                {data?.scopes.map(v => {
                  return (
                    <Badge
                      key={`scope-badge-${v}`}
                      ml=".5rem"
                      colorScheme="cateiru"
                    >
                      {v}
                    </Badge>
                  );
                })}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">プロンプト</Td>
              <Td>{data?.prompt}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">特定のユーザーのみ許可</Td>
              <Td>
                {data?.is_allow ? (
                  <Flex alignItems="center">
                    <Text>はい</Text>
                    <Button size="sm" ml=".5rem">
                      ユーザーを編集
                    </Button>
                  </Flex>
                ) : (
                  'いいえ'
                )}
              </Td>
            </Tr>
            {data?.org_member_only && (
              <Tr>
                <Td fontWeight="bold">組織のユーザーのみ</Td>
                <Td>{data?.org_member_only ? 'はい' : 'いいえ'}</Td>
              </Tr>
            )}
            <Tr>
              <Td fontWeight="bold">リダイレクトURL</Td>
              <Td>
                <UnorderedList ml="0" spacing="5px">
                  {data?.redirect_urls.map((v, i) => {
                    return (
                      <ListItem key={`redirect-url-${i}`} listStyleType="none">
                        <Link href={v} isExternal>
                          {v}
                        </Link>
                      </ListItem>
                    );
                  })}
                </UnorderedList>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">リファラーホスト</Td>
              <Td>
                <UnorderedList ml="0" spacing="5px">
                  {data?.referrer_urls.map((v, i) => {
                    return (
                      <ListItem key={`referrer-url-${i}`} listStyleType="none">
                        {v}
                      </ListItem>
                    );
                  })}
                </UnorderedList>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">クライアントID</Td>
              <Td>
                <Flex alignItems="center" h="100%">
                  <Button
                    size="sm"
                    mr=".5rem"
                    onClick={clientIdCopy.onCopy}
                    w="66px"
                  >
                    {clientIdCopy.hasCopied ? (
                      <TbCheck size="25px" />
                    ) : (
                      'コピー'
                    )}
                  </Button>
                  <Text
                    maxW="200px"
                    textOverflow="ellipsis"
                    whiteSpace="nowrap"
                    overflow="hidden"
                  >
                    {data?.client_id}
                  </Text>
                </Flex>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">クライアントシークレット</Td>
              <Td>
                <Button
                  size="sm"
                  mr=".5rem"
                  colorScheme="cateiru"
                  onClick={clientSecretCopy.onCopy}
                  w="66px"
                >
                  {clientSecretCopy.hasCopied ? (
                    <TbCheck size="25px" />
                  ) : (
                    'コピー'
                  )}
                </Button>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">作成日</Td>
              <Td>{new Date(data?.created_at ?? '').toLocaleString()}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">更新日</Td>
              <Td>{new Date(data?.updated_at ?? '').toLocaleString()}</Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
      <ButtonGroup w="100%" mt="1rem">
        <Button w="100%" colorScheme="cateiru">
          編集
        </Button>
        <Button w="100%">削除</Button>
      </ButtonGroup>
    </Margin>
  );
};
