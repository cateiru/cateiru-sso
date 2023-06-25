'use client';

import {
  Box,
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
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
  useColorModeValue,
} from '@chakra-ui/react';
import Link from 'next/link';
import React from 'react';
import useSWR from 'swr';
import {staffUserDetailFeather} from '../../../utils/swr/featcher';
import {ErrorType} from '../../../utils/types/error';
import type {UserDetail as UserDetailType} from '../../../utils/types/staff';
import {validateGender} from '../../../utils/validate';
import {Avatar} from '../../Common/Chakra/Avatar';
import {Error} from '../../Common/Error/Error';
import {Margin} from '../../Common/Margin';

interface Props {
  id: string;
}

export const UserDetail: React.FC<Props> = props => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  const {data, error} = useSWR<UserDetailType, ErrorType>(
    `/v2/admin/user_detail?user_id=${props.id}`,
    () => staffUserDetailFeather(props.id)
  );

  if (error) {
    return <Error {...error} />;
  }

  console.log(data);

  return (
    <Margin>
      <Breadcrumb>
        <BreadcrumbItem>
          <BreadcrumbLink as={Link} href="/staff">
            スタッフ管理画面
          </BreadcrumbLink>
        </BreadcrumbItem>

        <BreadcrumbItem>
          <BreadcrumbLink as={Link} href="/staff/users">
            ユーザー一覧
          </BreadcrumbLink>
        </BreadcrumbItem>

        <BreadcrumbItem isCurrentPage>
          <BreadcrumbLink>{data?.user.user_name}</BreadcrumbLink>
        </BreadcrumbItem>
      </Breadcrumb>
      <Box>
        <Center mt="3rem">
          <Avatar src={data?.user.avatar ?? ''} size="lg" />
        </Center>
        <Text textAlign="center" mt=".5rem" color={textColor} fontWeight="bold">
          {data?.user.user_name}
        </Text>
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
        <Text
          mt="2rem"
          mb="1rem"
          fontSize="1.5rem"
          color={textColor}
          fontWeight="bold"
        >
          スタッフ情報
        </Text>
        <TableContainer>
          <Table variant="simple">
            <Tbody>
              <Tr>
                <Td fontWeight="bold" color={textColor}>
                  スタッフ
                </Td>
                <Td>{data?.staff ? 'はい' : 'いいえ'}</Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold" color={textColor}>
                  メモ
                </Td>
                <Td>{data?.staff?.memo}</Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold" color={textColor}>
                  作成日
                </Td>
                <Td>
                  {data?.staff?.created_at
                    ? new Date(data?.staff?.created_at).toLocaleString()
                    : ''}
                </Td>
              </Tr>
              <Tr>
                <Td fontWeight="bold" color={textColor}>
                  更新日
                </Td>
                <Td>
                  {data?.staff?.updated_at
                    ? new Date(data?.staff?.updated_at).toLocaleString()
                    : ''}
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
          ブランド情報
        </Text>
        <TableContainer>
          <Table variant="simple">
            <Thead>
              <Tr>
                <Th>ブランドID</Th>
                <Th>加入日</Th>
                <Th></Th>
              </Tr>
            </Thead>
            <Tbody>
              {data?.user_brands.map(brand => {
                return (
                  <Tr key={`brand-${brand.id}`}>
                    <Td>{brand.brand_id}</Td>
                    <Td>{brand.created_at}</Td>
                    <Td>
                      <Button size="sm" colorScheme="cateiru">
                        詳細
                      </Button>
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
                      <Avatar src={client.image ?? ''} size="sm" />
                    </Td>
                    <Td>{client.name}</Td>
                    <Td>
                      <Button size="sm" colorScheme="cateiru">
                        詳細
                      </Button>
                    </Td>
                  </Tr>
                );
              })}
            </Tbody>
          </Table>
        </TableContainer>
      </Box>
    </Margin>
  );
};
