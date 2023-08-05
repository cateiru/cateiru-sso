import {Heading, VStack, useColorModeValue} from '@chakra-ui/react';
import {useRouter, useSearchParams} from 'next/navigation';
import React from 'react';
import {useRecoilValue} from 'recoil';
import {formatRedirectUrl} from '../../utils/format';
import {UserState} from '../../utils/state/atom';
import {CheckMark, CheckmarkProps} from '../Common/Icons/CheckMark';

export const LoginSuccess: React.FC = () => {
  const user = useRecoilValue(UserState);
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
        router.replace(formatRedirectUrl(redirectTo));
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
        こんにちは、{user?.user.user_name ?? ' ??? '}さん
      </Heading>
    </VStack>
  );
};
