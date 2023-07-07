'use client';

import {Box, Button, ButtonGroup, Center, Heading} from '@chakra-ui/react';
import Link from 'next/link';
import {config} from '../../utils/config';
import {Logo} from '../Common/Icons/Logo';

export const Top = () => {
  return (
    <Center h="100vh">
      <Box>
        <Center>
          <Logo size="50%" />
        </Center>
        <Heading
          textAlign="center"
          background="linear-gradient(124deg, #2bc4cf, #572bcf, #cf2ba1)"
          backgroundClip="text"
          fontSize={{base: '2rem', md: '3rem'}}
        >
          {config.title}
        </Heading>
        <Center mt="1rem">
          <ButtonGroup>
            <Button colorScheme="cateiru" as={Link} href="/register">
              アカウント作成
            </Button>
            <Button as={Link} href="/login">
              ログイン
            </Button>
          </ButtonGroup>
        </Center>
      </Box>
    </Center>
  );
};
