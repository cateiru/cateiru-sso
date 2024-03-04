import {useRouter, useSearchParams} from 'next/navigation';
import React from 'react';
import {api, fetch} from '../../utils/api';
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
import {useGetOauthLoginSession} from '../Login/useGetOauthLoginSession';
import {useOidc} from './useOidc';

export const useOidcRequire = (submit: () => Promise<void>) => {
  const [oidcError, setOidcError] = React.useState<OidcErrorType | null>(null);
  const [error, setError] = React.useState<ErrorType | null>(null);
  const [data, setData] = React.useState<PublicAuthenticationRequest | null>(
    null
  );
  const router = useRouter();
  const searchParams = useSearchParams();
  const getOauthLoginSession = useGetOauthLoginSession();

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

    let res;

    try {
      res = await fetch(api('/oidc/require'), {
        method: 'POST',
        body: params,
        headers: {
          Referer: document.referrer,
          ...getOauthLoginSession(),
        },
      });
    } catch (e) {
      if (e instanceof Error) {
        setError({
          message: e.message,
        });
      }
      return;
    }

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

    const getUrl = (searchParams?: {[key: string]: string}) => {
      const url = new URL(window.location.href);
      url.searchParams.set('redirect_done', '1');

      if (typeof searchParams !== 'undefined') {
        Object.keys(searchParams).forEach(key => {
          url.searchParams.set(key, searchParams[key]);
        });
      }

      return url.pathname + url.search;
    };

    const redirectDone = !!searchParams.get('redirect_done');

    const data = PublicAuthenticationRequestSchema.safeParse(response);
    if (data.success) {
      // promptに`none`がある場合、同意画面は表示させずにsubmitする
      // loginやselect_accountは無視する
      if (data.data.prompts.includes('none')) {
        submit();
        return;
      }

      // login後は login_session は undefined になる
      if (
        data.data.prompts.includes('login') &&
        typeof data.data.login_session !== 'undefined'
      ) {
        // ログインページへリダイレクトする
        router.replace(
          `/login?redirect_to=${encodeURIComponent(
            getUrl({
              oauth_login_session: data.data.login_session.login_session_token,
            })
          )}&oauth=1&oauth_login_session=${
            data.data.login_session.login_session_token
          }`
        );
        return;
      }

      // promptに`select_account`がある場合、アカウント選択画面を表示させる
      if (data.data.prompts.includes('select_account') && !redirectDone) {
        router.replace(
          `/switch_account?redirect_to=${encodeURIComponent(getUrl())}&oauth=1`
        );
        return;
      }

      setData(data.data);
      return data.data;
    }

    const noLoginData =
      PublicAuthenticationLoginSessionSchema.safeParse(response);
    if (noLoginData.success) {
      // ログインページへリダイレクトする
      router.replace(
        `/login?redirect_to=${encodeURIComponent(
          getUrl({
            oauth_login_session: noLoginData.data.login_session_token,
          })
        )}&oauth_login_session=${noLoginData.data.login_session_token}`
      );
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
