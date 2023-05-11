import {
  get,
  parseRequestOptionsFromJSON,
} from '@github/webauthn-json/browser-ponyfill';
import React from 'react';
import {useRequest} from '../Common/useRequest';
import {useToast} from '@chakra-ui/react';
import {LoginResponseSchema} from '../../utils/types/login';
import {useSetRecoilState} from 'recoil';
import {UserState} from '../../utils/state/atom';
import {useRouter} from 'next/router';

interface Returns {
  isConditionSupported: boolean;
  onClickWebAuthn: () => Promise<void>;
}

export const useWebAuthn = (): Returns => {
  const setUser = useSetRecoilState(UserState);
  const router = useRouter();
  const toast = useToast();

  const [isConditionSupported, setIsConditionSupported] = React.useState(true);

  const {request: getBeginKey} = useRequest('/v2/login/begin_webauthn');
  const {request: pushCredential} = useRequest('/v2/login/webathn');

  React.useEffect(() => {
    // ブラウザが対応していない場合は実施しない
    if (
      !PublicKeyCredential.isConditionalMediationAvailable ||
      !PublicKeyCredential.isConditionalMediationAvailable()
    ) {
      setIsConditionSupported(false);
      return;
    }

    const abort = new AbortController();

    const f = async () => {
      const res = await getBeginKey({
        mode: 'cors',
        credentials: 'include',
        method: 'POST',
      });
      if (!res) return;

      const beginData = parseRequestOptionsFromJSON(await res.json());
      beginData.signal = abort.signal;

      // See also: https://github.com/w3c/webauthn/wiki/Explainer:-WebAuthn-Conditional-UI
      // W3C上ではConditional UIはdraftなので一旦anyで型を無視している
      (beginData as any).mediation = 'conditional';

      let credential: Credential;
      try {
        credential = await get(beginData);
      } catch (e) {
        toast({
          title: 'WebAuthnエラー',
          status: 'error',
        });
        return;
      }

      const resCredential = await pushCredential({
        mode: 'cors',
        credentials: 'include',
        body: JSON.stringify(credential),
        headers: {
          'Content-Type': 'application/json',
        },
        method: 'POST',
      });
      if (!resCredential) return;

      const data = LoginResponseSchema.safeParse(await resCredential.json());
      if (data.success && data.data.user) {
        setUser({
          user: data.data.user,
        });
        await router.push('/');
        return;
      }

      toast({
        title: 'パースエラー',
        status: 'error',
      });
    };
    f();

    return () => {
      abort.abort();
    };
  }, []);

  const onClickWebAuthn = async () => {
    const res = await getBeginKey({
      mode: 'cors',
      credentials: 'include',
      method: 'POST',
    });
    if (!res) return;

    const beginData = parseRequestOptionsFromJSON(await res.json());

    let credential: Credential;
    try {
      credential = await get(beginData);
    } catch (e) {
      toast({
        title: 'WebAuthnエラー',
        status: 'error',
      });
      return;
    }

    const resCredential = await pushCredential({
      mode: 'cors',
      credentials: 'include',
      body: JSON.stringify(credential),
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'POST',
    });
    if (!resCredential) return;

    const data = LoginResponseSchema.safeParse(await resCredential.json());
    if (data.success && data.data.user) {
      setUser({
        user: data.data.user,
      });
      await router.push('/');
      return;
    }

    toast({
      title: 'パースエラー',
      status: 'error',
    });
  };

  return {isConditionSupported, onClickWebAuthn};
};
