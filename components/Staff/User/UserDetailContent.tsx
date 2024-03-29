import {
  Button,
  Center,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import {UserDetail} from '../../../utils/types/staff';
import {validateGender} from '../../../utils/validate';
import {Avatar} from '../../Common/Chakra/Avatar';
import {Link} from '../../Common/Next/Link';
import {useSecondaryColor} from '../../Common/useColor';
import {UserDetailBrand} from './UserDetailBrand';
import {UserDetailStaff} from './UserDetailStaff';

interface Props {
  data?: UserDetail;
}

export const UserDetailContent: React.FC<Props> = ({data}) => {
  const textColor = useSecondaryColor();

  return (
    <>
      <Center mt="3rem">
        <Avatar src={data?.user.avatar ?? ''} size="lg" />
      </Center>
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        ユーザー情報
      </Text>
      <TableContainer>
        <Table variant="simple">
          <Tbody>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                ID
              </Td>
              <Td>{data?.user.id}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                ユーザー名
              </Td>
              <Td>{data?.user.user_name}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                メールアドレス
              </Td>
              <Td>{data?.user.email}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                名字
              </Td>
              <Td>{data?.user.family_name}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                ミドルネーム
              </Td>
              <Td>{data?.user.middle_name}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                名前
              </Td>
              <Td>{data?.user.given_name}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                性別
              </Td>
              <Td>
                {data?.user.gender} (
                {data?.user.gender ? validateGender(data?.user.gender) : ''})
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                誕生日
              </Td>
              <Td>
                {data?.user.birthdate
                  ? new Date(data?.user.birthdate).toLocaleDateString()
                  : ''}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                ロケール
              </Td>
              <Td>{data?.user.locale_id}</Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                アカウント作成日
              </Td>
              <Td>
                {data?.user.created_at
                  ? new Date(data?.user.created_at).toLocaleString()
                  : ''}
              </Td>
            </Tr>
            <Tr>
              <Td fontWeight="bold" color={textColor}>
                アカウント更新日
              </Td>
              <Td>
                {data?.user.updated_at
                  ? new Date(data?.user.updated_at).toLocaleString()
                  : ''}
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
      <UserDetailStaff staff={data?.staff} userId={data?.user.id} />
      <UserDetailBrand
        brands={data?.user_brands ?? []}
        userId={data?.user.id ?? ''}
      />
      <Text
        mt="2rem"
        mb="1rem"
        fontSize="1.5rem"
        color={textColor}
        fontWeight="bold"
      >
        作成したクライアント情報
      </Text>
      <TableContainer>
        <Table variant="simple">
          <Thead>
            <Tr>
              <Th></Th>
              <Th>クライアント名</Th>
              <Th></Th>
            </Tr>
          </Thead>
          <Tbody>
            {data?.clients.map(client => {
              return (
                <Tr key={`client-${client.client_id}`}>
                  <Td>
                    <Center>
                      <Avatar src={client.image ?? ''} size="sm" />
                    </Center>
                  </Td>
                  <Td>{client.name}</Td>
                  <Td>
                    <Button
                      size="sm"
                      colorScheme="cateiru"
                      as={Link}
                      href={`/staff/client/${client.client_id}`}
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
