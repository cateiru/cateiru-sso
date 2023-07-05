import {
  Badge,
  Button,
  ButtonGroup,
  Center,
  Link,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import {OrganizationDetail} from '../../../utils/types/staff';
import {Avatar} from '../../Common/Chakra/Avatar';
import {Link as NextLink} from '../../Common/Next/Link';

export const OrgDetailContent: React.FC<OrganizationDetail> = data => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <>
      <Center mt="3rem">
        <Avatar src={data.org.image ?? ''} size="lg" />
      </Center>
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        組織情報
      </Text>
      <TableContainer>
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織ID
              </Td>
              <Td>{data.org.id}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織名
              </Td>
              <Td>{data.org.name}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織URL
              </Td>
              <Td>
                {data.org.link ? (
                  <Link href={data.org.link} isExternal>
                    {data.org.link}
                  </Link>
                ) : (
                  ''
                )}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織作成日
              </Td>
              <Td>{new Date(data.org.created_at).toLocaleString()}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                組織更新日
              </Td>
              <Td>{new Date(data.org.updated_at).toLocaleString()}</Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
      <ButtonGroup mt="1rem" w="100%">
        <Button
          w="100%"
          colorScheme="cateiru"
          as={NextLink}
          href={`/staff/org/edit/${data.org.id}`}
        >
          組織を編集
        </Button>
        <Button w="100%">組織削除</Button>
      </ButtonGroup>
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        組織ユーザー
      </Text>
      <TableContainer>
        <Table variant="simple">
          <Thead>
            <Tr>
              <Th></Th>
              <Th>ユーザー名</Th>
              <Th textAlign="center">ロール</Th>
              <Th>参加日</Th>
              <Th></Th>
            </Tr>
          </Thead>
          <Tbody>
            {data.users.map(user => {
              const joinDate = new Date(user.created_at);

              const badgeColor = () => {
                if (user.role === 'admin') {
                  return 'red';
                }

                if (user.role === 'member') {
                  return 'blue';
                }

                return 'gray';
              };

              return (
                <Tr key={`org-user-${user.id}`}>
                  <Td>
                    <Avatar src={user.user.avatar ?? ''} size="sm" />
                  </Td>
                  <Td>{user.user.user_name}</Td>
                  <Td textAlign="center">
                    <Badge colorScheme={badgeColor()}>{user.role}</Badge>
                  </Td>
                  <Td>{joinDate.toLocaleDateString()}</Td>
                  <Td>
                    <Button
                      colorScheme="cateiru"
                      size="sm"
                      as={NextLink}
                      href={`/staff/user/${user.user.id}`}
                    >
                      ユーザー詳細
                    </Button>
                  </Td>
                </Tr>
              );
            })}
          </Tbody>
        </Table>
      </TableContainer>
      <Button colorScheme="cateiru" w="100%" mt="1rem">
        ユーザーを追加する
      </Button>
    </>
  );
};
