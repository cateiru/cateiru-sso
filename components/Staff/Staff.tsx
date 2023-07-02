'use client';

import {Button} from '@chakra-ui/react';
import {Link} from '../Common/Next/Link';
import {StaffCard} from './StaffCard';

export const Staff = () => {
  return (
    <>
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
          href="/staff/brands"
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
          href="/staff/clients"
        >
          クライアント一覧
        </Button>
      </StaffCard>
      <StaffCard title="デプロイデータ">
        <Button
          variant="outline"
          colorScheme="cateiru"
          w="100%"
          as={Link}
          href="/staff/deploy_data"
        >
          デプロイデータ
        </Button>
      </StaffCard>
    </>
  );
};
