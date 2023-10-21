'use client';

import {Box, Divider, Heading} from '@chakra-ui/react';
import {useSearchParams} from 'next/navigation';
import {useBorderColor, useShadowColor} from '../Common/useColor';
import {AccountList} from './AccountList';

export const SwitchAccount = () => {
  const borderColor = useBorderColor();
  const shadowColor = useShadowColor();

  const isOauth = !!useSearchParams().get('oauth');

  return (
    <Box
      w={{base: '96%', sm: '450px'}}
      h={{base: '600px', sm: '700px'}}
      margin="auto"
      mt="3rem"
      borderRadius="10px"
      borderColor={borderColor}
      mb={{base: '0', sm: '3rem'}}
      boxShadow={{base: 'none', sm: `0px 0px 7px -2px ${shadowColor}`}}
    >
      <Box h="150px">
        <Heading
          textAlign="center"
          pt="40px"
          mx=".5rem"
          fontSize={{base: '1.5rem', sm: '1.8rem'}}
        >
          {isOauth ? '使用するアカウントを' : 'ログインするアカウントを'}
          <br />
          選択してください
        </Heading>
        <Divider mt="1.5rem" w="90%" mx="auto" />
      </Box>
      <AccountList isOauth={isOauth} />
    </Box>
  );
};
