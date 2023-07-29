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
  Skeleton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Tr,
  UnorderedList,
  useClipboard,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import {useParams} from 'next/navigation';
import {TbCheck} from 'react-icons/tb';
import useSWR from 'swr';
import {clientFetcher} from '../../utils/swr/client';
import {ClientDetail as ClientDetailType} from '../../utils/types/client';
import {ErrorType} from '../../utils/types/error';
import {validatePrompt} from '../../utils/validate';
import {Avatar} from '../Common/Chakra/Avatar';
import {Error} from '../Common/Error/Error';
import {Margin} from '../Common/Margin';

export const ClientDetail = () => {
  const {id} = useParams();

  const {data, error} = useSWR<ClientDetailType, ErrorType>(
    `/v2/client/?client_id=${id}`,
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
              <Td>{data ? data?.name : <Skeleton w="200px" h="1rem" />}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">クライアントの説明</Td>
              <Td>
                <Text as="pre">{data?.description}</Text>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">スコープ</Td>
              <Td>
                {data ? (
                  data?.scopes.map(v => {
                    return (
                      <Badge
                        key={`scope-badge-${v}`}
                        ml=".5rem"
                        colorScheme="cateiru"
                      >
                        {v}
                      </Badge>
                    );
                  })
                ) : (
                  <Flex>
                    <Skeleton w="40px" h="20px" ml=".5rem" />
                    <Skeleton w="40px" h="20px" ml=".5rem" />
                    <Skeleton w="40px" h="20px" ml=".5rem" />
                  </Flex>
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">認証</Td>
              <Td>
                {data ? (
                  validatePrompt(data.prompt)
                ) : (
                  <Skeleton w="200px" h="1rem" />
                )}
              </Td>
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
                <Td fontWeight="bold">組織のユーザーのみ許可</Td>
                <Td>{data?.org_member_only ? 'はい' : 'いいえ'}</Td>
              </Tr>
            )}
            <Tr>
              <Td fontWeight="bold">リダイレクトURL</Td>
              <Td>
                <UnorderedList ml="0" spacing="5px">
                  {data ? (
                    data?.redirect_urls.map((v, i) => {
                      return (
                        <ListItem
                          key={`redirect-url-${i}`}
                          listStyleType="none"
                        >
                          <Link href={v} isExternal>
                            {v}
                          </Link>
                        </ListItem>
                      );
                    })
                  ) : (
                    <Skeleton w="200px" h="1rem" />
                  )}
                </UnorderedList>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">リファラーホスト</Td>
              <Td>
                <UnorderedList ml="0" spacing="5px">
                  {data ? (
                    data?.referrer_urls.map((v, i) => {
                      return (
                        <ListItem
                          key={`referrer-url-${i}`}
                          listStyleType="none"
                        >
                          {v}
                        </ListItem>
                      );
                    })
                  ) : (
                    <Skeleton w="200px" h="1rem" />
                  )}
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
                  {data ? (
                    <Text
                      maxW="200px"
                      textOverflow="ellipsis"
                      whiteSpace="nowrap"
                      overflow="hidden"
                    >
                      {data?.client_id}
                    </Text>
                  ) : (
                    <Skeleton w="200px" h="1rem" />
                  )}
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
              <Td>
                {data ? (
                  new Date(data?.created_at ?? '').toLocaleString()
                ) : (
                  <Skeleton w="200px" h="1rem" />
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold">更新日</Td>
              <Td>
                {data ? (
                  new Date(data?.updated_at ?? '').toLocaleString()
                ) : (
                  <Skeleton maxW="200px" h="1rem" />
                )}
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
      <ButtonGroup w="100%" mt="1rem">
        <Button
          w="100%"
          colorScheme="cateiru"
          as={NextLink}
          href={`/client/edit/${id}`}
        >
          編集
        </Button>
        <Button w="100%">削除</Button>
      </ButtonGroup>
    </Margin>
  );
};
