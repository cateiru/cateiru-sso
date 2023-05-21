import {Heading, VStack, useColorModeValue} from '@chakra-ui/react';
import {useRouter, useSearchParams} from 'next/navigation';
import React from 'react';
import {CheckMark, CheckmarkProps} from '../Common/Icons/CheckMark';

export const CompleteRegisterPage = () => {
  const router = useRouter();
  const params = useSearchParams();
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
    const t = setTimeout(() => {
      const redirectTo = params.get('redirect_to');
      if (typeof redirectTo === 'string') {
        try {
          const url = new URL(redirectTo);
          router.replace(url.pathname);
        } catch {
          router.replace(redirectTo);
        }
      } else {
        router.push('/profile');
      }
    }, 2000);

    return () => {
      clearTimeout(t);
    };
  }, []);

  return (
    <VStack mt="3rem">
      <CheckMark {...checkmarkProps} />
      <Heading textAlign="center" mt=".5rem" color={checkmarkProps.bgColor}>
        アカウントを作成しました🎉
      </Heading>
    </VStack>
  );
};
