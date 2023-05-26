import {Button} from '@chakra-ui/react';
import {useSWRConfig} from 'swr';
import {useRegisterWebAuthn} from './useRegisterWebAuthn';

export const RegisterWebAuthn = () => {
  const {mutate} = useSWRConfig();
  const {registerWebAuthn, load} = useRegisterWebAuthn(() => {
    // SWRのキャッシュ飛ばす
    mutate(
      key => typeof key === 'string' && key.startsWith('/v2/account/webauthn'),
      undefined,
      {revalidate: true}
    );
  });

  return (
    <Button
      onClick={registerWebAuthn}
      isLoading={load}
      w="100%"
      colorScheme="cateiru"
    >
      生体認証を登録する
    </Button>
  );
};
