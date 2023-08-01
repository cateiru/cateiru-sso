import {useRouter, useSearchParams} from 'next/navigation';
import React from 'react';
import {CreateAccountRegisterEmailResponseSchema} from '../../utils/types/createAccount';
import {useRequest} from '../Common/useRequest';

interface Returns {
  isInvite: boolean;
}

export const useInvite = (
  handleSuccess: (email: string, token: string) => void
): Returns => {
  const prams = useSearchParams();
  const router = useRouter();

  const {request} = useRequest('/v2/register/invite_generate_session');

  const [isInvite, setIsInvite] = React.useState<boolean>(false);

  // invite_token が存在する場合、それを使用して登録用のセッションを取得する
  React.useEffect(() => {
    const token = prams.get('invite_token');
    const email = prams.get('email');

    const f = async () => {
      if (typeof token !== 'string') return;
      if (typeof email !== 'string') return;

      // 重複して送信しないようにする
      if (isInvite) return;
      // 成功、失敗に関わらずにフラグを立てる
      setIsInvite(true);

      const form = new FormData();

      form.append('invite_token', token);
      form.append('email', email);

      const res = await request({
        method: 'POST',
        mode: 'cors',
        credentials: 'include',
        body: form,
      });

      if (res) {
        const data = CreateAccountRegisterEmailResponseSchema.safeParse(
          await res.json()
        );
        if (data.success) {
          handleSuccess(email, data.data.register_token);
          return;
        }
        console.error(data.error);
      }

      router.replace('/');
    };
    f();
  }, [prams.get('invite_token'), prams.get('email')]);

  return {
    isInvite,
  };
};
