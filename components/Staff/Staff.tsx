'use client';

import {Button, Heading, Text, useColorModeValue} from '@chakra-ui/react';
import Link from 'next/link';
import {config} from '../../utils/config';
import {Margin} from '../Common/Margin';
import {StaffCard} from './StaffCard';

export const Staff = () => {
  const textColor = useColorModeValue('gray.500', 'gray.400');

  return (
    <Margin>
      <Heading textAlign="center" mt="3rem">
        スタッフ管理画面
      </Heading>
      <Text color={textColor} textAlign="center" mt=".5rem">
        Revision: {config.revision}{' '}
        {config.branchName && `/ Branch: ${config.branchName}`}
      </Text>
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
    </Margin>
  );
};
