'use client';

import {
  Breadcrumb,
  Heading,
  BreadcrumbItem,
  BreadcrumbLink,
} from '@chakra-ui/react';
import Link from 'next/link';
import {Margin} from '../../Common/Margin';
import {UsersTable} from './UsersTable';

export const Users = () => {
  return (
    <Margin>
      <Breadcrumb>
        <BreadcrumbItem>
          <BreadcrumbLink as={Link} href="/staff">
            スタッフ管理画面
          </BreadcrumbLink>
        </BreadcrumbItem>
        <BreadcrumbItem isCurrentPage>
          <BreadcrumbLink>ユーザー一覧</BreadcrumbLink>
        </BreadcrumbItem>
      </Breadcrumb>
      <Heading textAlign="center" mt="3rem" mb="2rem">
        ユーザー一覧
      </Heading>
      <UsersTable />
    </Margin>
  );
};
