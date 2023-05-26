import {useToast} from '@chakra-ui/react';
import {
  create,
  parseCreationOptionsFromJSON,
} from '@github/webauthn-json/browser-ponyfill';
import React from 'react';
import {useRequest} from '../../../Common/useRequest';

interface Returns {
  registerWebAuthn: () => Promise<void>;
  load: boolean;
}

export const useRegisterWebAuthn = (
  afterRegisterHandler?: () => void
): Returns => {
  const toast = useToast();
  const {request: beginWebauthnRequest} = useRequest(
    '/v2/account/begin_webauthn'
  );
  const {request: pushCredential} = useRequest('/v2/account/webauthn');

  const [load, setLoad] = React.useState(false);
  const abortRef = React.useRef<AbortController>();

  React.useEffect(() => {
    const abort = new AbortController();
    abortRef.current = abort;

    return () => {
      if (abortRef.current) {
        abortRef.current.abort();
      }
    };
  }, []);

  const registerWebAuthnHandle = async () => {
    const res = await beginWebauthnRequest({
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
    });
    if (!res) return;

    const beginData = parseCreationOptionsFromJSON(await res.json());

    if (abortRef.current) {
      beginData.signal = abortRef.current.signal;
    }

    let credential: Credential;
    try {
      credential = await create(beginData);
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

    toast({
      title: '生体認証を登録しました',
      status: 'success',
    });

    // 後処理
    afterRegisterHandler && afterRegisterHandler();
  };

  const registerWebAuthn = async () => {
    setLoad(true);
    await registerWebAuthnHandle();
    setLoad(false);
  };

  return {
    registerWebAuthn,
    load,
  };
};
