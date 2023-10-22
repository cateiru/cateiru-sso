import {useRouter, useSearchParams} from 'next/navigation';
import React from 'react';
import {useRecoilState} from 'recoil';
import {api} from '../../utils/api';
import {OAuthLoginSessionState} from '../../utils/state/atom';
import {
  PublicAuthenticationLoginSessionSchema,
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

export const useOidcRequire = (submit: () => Promise<void>) => {
  const [oidcError, setOidcError] = React.useState<OidcErrorType | null>(null);
  const [error, setError] = React.useState<ErrorType | null>(null);
  const [data, setData] = React.useState<PublicAuthenticationRequest | null>(
    null
  );
  const [oauthLoginSession, setOAuthLoginSession] = useRecoilState(
    OAuthLoginSessionState
  );
  const router = useRouter();
  const searchParams = useSearchParams();

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

    // TODO: サーバに送る
    console.log(oauthLoginSession);

    const res = await fetch(api('/v2/oidc/require'), {
      credentials: 'include',
      mode: 'cors',
      method: 'POST',
      body: params,
      headers: {
        Referer: document.referrer,
      },
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

    const url = new URL(window.location.href);
    url.searchParams.set('redirect_done', '1');
    const relativeUrl = url.pathname + url.search;
    const redirectDone = !!searchParams.get('redirect_done');

    const data = PublicAuthenticationRequestSchema.safeParse(response);
    if (data.success) {
      // promptに`none`がある場合、同意画面は表示させずにsubmitする
      // loginやselect_accountは無視する
      if (data.data.prompts.includes('none')) {
        submit();
        return;
      }

      if (data.data.prompts.includes('login') && !redirectDone) {
        // prompt = login 時にセッションが来ないことはないが
        // 念のためチェックする
        if (data.data.login_session === null) {
          setError({
            message: 'error',
          });
          return;
        }

        setOAuthLoginSession(data.data.login_session);

        // ログインページへリダイレクトする
        router.replace(`/login?redirect_to=${encodeURIComponent(relativeUrl)}`);
        return;
      }

      // promptに`select_account`がある場合、アカウント選択画面を表示させる
      if (data.data.prompts.includes('select_account') && !redirectDone) {
        router.replace(
          `/switch_account?redirect_to=${encodeURIComponent(
            relativeUrl
          )}&oauth=1`
        );
        return;
      }

      setData(data.data);
      setOAuthLoginSession(undefined);
      return data.data;
    }

    const noLoginData =
      PublicAuthenticationLoginSessionSchema.safeParse(response);
    if (noLoginData.success) {
      setOAuthLoginSession(noLoginData.data);

      // ログインページへリダイレクトする
      router.replace(`/login?redirect_to=${encodeURIComponent(relativeUrl)}`);
      return;
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
