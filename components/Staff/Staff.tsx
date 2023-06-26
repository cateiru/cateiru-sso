'use client';

import {Button} from '@chakra-ui/react';
import Link from 'next/link';
import {StaffCard} from './StaffCard';
import {StaffFrame} from './StaffFrame';

export const Staff = () => {
  return (
    <StaffFrame
      title="スタッフ管理画面"
      paths={[{pageName: 'スタッフ管理画面'}]}
    >
      <StaffCard title="ユーザー一覧">
        <Button
          variant="outline"
          colorScheme="cateiru"
          w="100%"
          as={Link}
          href="/staff/users"
        >
          ユーザー一覧
        </Button>
      </StaffCard>
      <StaffCard title="ブランド一覧">
        <Button
          variant="outline"
          colorScheme="cateiru"
          w="100%"
          as={Link}
          href="/staff"
        >
          ブランド一覧
        </Button>
      </StaffCard>
      <StaffCard title="クライアント一覧">
        <Button
          variant="outline"
          colorScheme="cateiru"
          w="100%"
          as={Link}
          href="/staff"
        >
          クライアント一覧
        </Button>
      </StaffCard>
    </StaffFrame>
  );
};
