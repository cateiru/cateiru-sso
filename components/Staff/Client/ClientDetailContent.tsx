import {
  Badge,
  Button,
  Center,
  Flex,
  Link,
  ListItem,
  Skeleton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
  UnorderedList,
} from '@chakra-ui/react';
import React from 'react';
import {ClientDetail} from '../../../utils/types/staff';
import {validatePrompt} from '../../../utils/validate';
import {Avatar} from '../../Common/Chakra/Avatar';
import {Link as NextLink} from '../../Common/Next/Link';
import {useSecondaryColor} from '../../Common/useColor';

interface Props {
  data: ClientDetail | undefined;
}

export const ClientDetailContent: React.FC<Props> = ({data}) => {
  const textColor = useSecondaryColor();

  return (
    <>
      <Center mt="3rem">
        <Avatar src={data?.client.image ?? ''} size="lg" />
      </Center>
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        クライアント情報
      </Text>
      <TableContainer>
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                クライアント名
              </Td>
              <Td>
                {data ? data?.client.name : <Skeleton w="200px" h="1rem" />}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                クライアントの説明
              </Td>
              <Td>
                <Text as="pre">{data?.client.description}</Text>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                スコープ
              </Td>
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
              <Td fontWeight="bold" color={textColor}>
                認証
              </Td>
              <Td>
                {data ? (
                  validatePrompt(data.client.prompt)
                ) : (
                  <Skeleton w="200px" h="1rem" />
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                特定のユーザーのみ許可
              </Td>
              <Td>
                {data ? (
                  data?.client.is_allow ? (
                    'はい'
                  ) : (
                    'いいえ'
                  )
                ) : (
                  <Skeleton w="100px" h="1rem" />
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織のユーザーのみ許可
              </Td>
              <Td>
                {data ? (
                  data?.client.org_member_only ? (
                    'はい'
                  ) : (
                    'いいえ'
                  )
                ) : (
                  <Skeleton w="100px" h="1rem" />
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                リダイレクトURL
              </Td>
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
              <Td fontWeight="bold" color={textColor}>
                リファラーホスト
              </Td>
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
              <Td fontWeight="bold" color={textColor}>
                クライアントID
              </Td>
              <Td>
                {data ? (
                  <Text
                    maxW="200px"
                    textOverflow="ellipsis"
                    whiteSpace="nowrap"
                    overflow="hidden"
                  >
                    {data?.client.client_id}
                  </Text>
                ) : (
                  <Skeleton w="200px" h="1rem" />
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                クライアントシークレット
              </Td>
              <Td>
                {data ? (
                  <Text
                    maxW="200px"
                    textOverflow="ellipsis"
                    whiteSpace="nowrap"
                    overflow="hidden"
                  >
                    {data?.client.client_secret}
                  </Text>
                ) : (
                  <Skeleton w="200px" h="1rem" />
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                作成者
              </Td>
              <Td>
                <Button
                  size="sm"
                  as={NextLink}
                  colorScheme="cateiru"
                  href={`/staff/user/${data?.client.owner_user_id}`}
                >
                  詳細
                </Button>
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織
              </Td>
              <Td>
                {data?.client.org_id ? (
                  <Button
                    size="sm"
                    as={NextLink}
                    colorScheme="cateiru"
                    href={`/staff/org/${data?.client.org_id}`}
                  >
                    詳細
                  </Button>
                ) : (
                  '未所属'
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                作成日
              </Td>
              <Td>
                {data ? (
                  new Date(data?.client.created_at ?? '').toLocaleString()
                ) : (
                  <Skeleton w="200px" h="1rem" />
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                更新日
              </Td>
              <Td>
                {data ? (
                  new Date(data?.client.updated_at ?? '').toLocaleString()
                ) : (
                  <Skeleton maxW="200px" h="1rem" />
                )}
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        許可ユーザー
      </Text>
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
              ? data.allow_rules
                  .filter(v => v.user)
                  .map(v => {
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
                              colorScheme="cateiru"
                              as={NextLink}
                              href={`/staff/user/${v.user?.id}`}
                            >
                              詳細
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
                            <Skeleton w="32px" h="32px" borderRadius="50%" />
                          </Center>
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
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        許可ドメイン
      </Text>
      <TableContainer>
        <Table variant="simple">
          <Thead>
            <Tr>
              <Th>メールドメイン</Th>
            </Tr>
          </Thead>
          <Tbody>
            {data
              ? data.allow_rules
                  .filter(v => v.email_domain)
                  .map(v => {
                    if (v.email_domain) {
                      return (
                        <Tr key={`allow-user-${v.id}`}>
                          <Td>{v.email_domain}</Td>
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
                          <Skeleton w="52px" h="32px" borderRadius="0.375rem" />
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
