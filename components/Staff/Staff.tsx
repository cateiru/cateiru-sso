'use client';

import {Button, Heading} from '@chakra-ui/react';
import {Margin} from '../Common/Margin';
import {StaffCard} from './StaffCard';

export const Staff = () => {
  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        スタッフ管理画面
      </Heading>
      <StaffCard title="全ユーザー参照">
        <Button variant="outline" colorScheme="cateiru" w="100%">
          全ユーザー参照
        </Button>
      </StaffCard>
    </Margin>
  );
};
