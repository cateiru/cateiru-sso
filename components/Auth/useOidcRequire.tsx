import React from 'react';
import {api} from '../../utils/api';
import {
  PublicAuthenticationRequest,
  PublicAuthenticationRequestSchema,
} from '../../utils/types/auth';
import {
  ErrorSchema,
  ErrorType,
  OidcErrorSchema,
  OidcErrorType,
} from '../../utils/types/error';
import {useOidc} from './useOidc';

export const useOidcRequire = () => {
  const [oidcError, setOidcError] = React.useState<OidcErrorType | null>(null);
  const [error, setError] = React.useState<ErrorType | null>(null);
  const [data, setData] = React.useState<PublicAuthenticationRequest | null>(
    null
  );

  const {getFormParams} = useOidc();

  const require = async (): Promise<
    PublicAuthenticationRequest | undefined
  > => {
    let params: FormData;
    try {
      params = getFormParams();
    } catch (e) {
      if (e instanceof Error) {
        console.error(e.message);
        setError({
          message: 'パラメータが不正です',
        });
      }
      return;
    }

    const res = await fetch(api('/v2/oidc/require'), {
      credentials: 'include',
      mode: 'cors',
      method: 'POST',
      body: params,
    });

    const response = await res.json();

    if (!res.ok) {
      // oidcのエラー
      const oidcErrorData = OidcErrorSchema.safeParse(response);
      if (oidcErrorData.success) {
        setOidcError(oidcErrorData.data);
        return;
      }

      // 通常のエラー
      // 認証失敗などはこちら
      const errorData = ErrorSchema.safeParse(response);
      if (errorData.success) {
        setError(errorData.data);
      }
    }

    const data = PublicAuthenticationRequestSchema.safeParse(response);
    if (data.success) {
      setData(data.data);
      return data.data;
    }

    setError({
      message: '予期せぬエラーが発生しました。',
    });

    return;
  };

  return {
    data,
    oidcError,
    error,
    require,
  };
};
