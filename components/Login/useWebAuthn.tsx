import {useToast} from '@chakra-ui/react';
import {
  get,
  parseRequestOptionsFromJSON,
} from '@github/webauthn-json/browser-ponyfill';
import React from 'react';
import {LoginResponseSchema} from '../../utils/types/login';
import {User} from '../../utils/types/user';
import {useRequest} from '../Common/useRequest';

interface Returns {
  isConditionSupported: boolean;
  onClickWebAuthn: () => Promise<void>;
}

export const useWebAuthn = (loginSuccess: (user: User) => void): Returns => {
  const toast = useToast();

  const [isConditionSupported, setIsConditionSupported] = React.useState(true);
  const abortRef = React.useRef<AbortController>();

  const {request: getBeginKey} = useRequest('/v2/login/begin_webauthn');
  const {request: pushCredential} = useRequest('/v2/login/webathn');

  React.useEffect(() => {
    const abort = new AbortController();
    abortRef.current = abort;

    // ブラウザが対応していない場合は実施しない
    if (
      !PublicKeyCredential.isConditionalMediationAvailable ||
      !PublicKeyCredential.isConditionalMediationAvailable()
    ) {
      setIsConditionSupported(false);
      return;
    }

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
        // シグナルがAbortされたらエラー出さないでReturn
        if (abort.signal.aborted) return;

        // 1passwordの拡張機能は`credentials.get`をモックするが、Conditional UIに対応しておらず、
        // TypeErrorが発生する。そのため、その場合は一時的にパスキーログイン用のボタンを出現させる
        if(e instanceof TypeError) {
          setIsConditionSupported(false)
          return
        }

        console.error(e);
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
      if (data.success) {
        if (data.data.user) {
          loginSuccess(data.data.user);
        }
        return;
      }
      console.error(data.error);

      toast({
        title: 'パースエラー',
        status: 'error',
      });
    };
    f();

    return () => {
      if (abortRef.current) {
        abortRef.current.abort();
      }
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

    if (abortRef.current) {
      beginData.signal = abortRef.current.signal;
    }

    let credential: Credential;
    try {
      credential = await get(beginData);
    } catch (e) {
      // シグナルがAbortされたらエラー出さないでReturn
      if (abortRef.current && abortRef.current.signal.aborted) return;

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
    if (data.success) {
      if (data.data.user) {
        loginSuccess(data.data.user);
      }
      return;
    }
    console.error(data.error);

    toast({
      title: 'パースエラー',
      status: 'error',
    });
  };

  return {isConditionSupported, onClickWebAuthn};
};
