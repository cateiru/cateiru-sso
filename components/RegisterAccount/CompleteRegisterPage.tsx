import {Heading, VStack, useColorModeValue} from '@chakra-ui/react';
import {useRouter, useSearchParams} from 'next/navigation';
import React from 'react';
import {formatRedirectUrl} from '../../utils/format';
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
    // FedCMのために、ブラウザにログイン状態を伝える
    // まだ提案段階の使用なのでanyで無理やり適用している
    // ref. https://github.com/fedidcg/login-status
    // ref2. https://developers.google.com/privacy-sandbox/blog/fedcm-chrome-120-updates?hl=ja
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const login = (navigator as any).login;
    if (typeof login !== 'undefined') {
      login.setStatus('logged-in');
    }

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
        アカウントを作成しました🎉
      </Heading>
    </VStack>
  );
};
