import {Heading, VStack, useColorModeValue} from '@chakra-ui/react';
import {useRouter} from 'next/router';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {CheckMark, CheckmarkProps} from '../Common/Icons/CheckMark';

export const LoginSuccess: React.FC = () => {
  const user = useRecoilValue(UserState);
  const router = useRouter();
  const checkmarkProps = useColorModeValue<CheckmarkProps, CheckmarkProps>(
    {
      size: 100,
      bgColor: '#572bcf',
      color: '#fff',
    },
    {
      size: 100,
      bgColor: '#2bc4cf',
      color: '#fff',
    }
  );

  React.useEffect(() => {
    if (!router.isReady) return;

    const t = setTimeout(() => {
      if (typeof router.query.redirect_to === 'string') {
        try {
          const url = new URL(router.query.redirect_to);
          router.replace(url.pathname);
        } catch {
          router.replace(router.query.redirect_to);
        }
      } else {
        router.replace('/');
      }
    }, 2000);

    return () => {
      clearTimeout(t);
    };
  }, [router.isReady]);

  return (
    <VStack mt="3rem">
      <CheckMark {...checkmarkProps} />
      <Heading textAlign="center" mt=".5rem" color={checkmarkProps.bgColor}>
        こんにちは、{user?.user.user_name ?? ' ??? '}さん
      </Heading>
    </VStack>
  );
};
