import {PublicAuthenticationOauthResponse} from '../../utils/types/auth';
import {useRequest} from '../Common/useRequest';
import {useGetOauthLoginSession} from '../Login/useGetOauthLoginSession';
import {useOidc} from './useOidc';

interface Returns {
  submit: () => Promise<void>;
  cancel: () => Promise<void>;
}

export const useOidcSubmit = (): Returns => {
  const {request: submitRequest} = useRequest('/oidc/login');
  const {request: cancelRequest} = useRequest('/oidc/cancel');

  const {getFormParams} = useOidc();
  const getOauthLoginSession = useGetOauthLoginSession();

  const p = (): FormData | undefined => {
    try {
      return getFormParams();
    } catch (e) {
      // すでに、previewでチェック済みなはずなのでエラーになることはない
      if (e instanceof Error) {
        console.error(e.message);
      }
      return;
    }
  };

  const submit = async () => {
    const params = p();
    if (typeof params === 'undefined') {
      return;
    }

    const response = await submitRequest({
      credentials: 'include',
      mode: 'cors',
      method: 'POST',
      body: params,
      headers: {
        Referer: document.referrer,
        ...getOauthLoginSession(),
      },
    });
    if (typeof response === 'undefined') {
      return;
    }

    const data = PublicAuthenticationOauthResponse.safeParse(
      await response?.json()
    );
    if (data.success) {
      console.info(`redirect to ${data.data.redirect_url}`);
      window.open(data.data.redirect_url, '_self');
      return;
    }

    console.error(data.error);
    return;
  };

  const cancel = async () => {
    const params = p();
    if (typeof params === 'undefined') {
      return;
    }

    const response = await cancelRequest({
      credentials: 'include',
      mode: 'cors',
      method: 'POST',
      body: params,
      headers: {
        Referer: document.referrer,
        ...getOauthLoginSession(),
      },
    });
    if (typeof response === 'undefined') {
      return;
    }

    const data = PublicAuthenticationOauthResponse.safeParse(
      await response?.json()
    );
    if (data.success) {
      console.info(`redirect to ${data.data.redirect_url}`);
      window.open(data.data.redirect_url, '_self');
      return;
    }

    console.error(data.error);
    return;
  };

  return {
    submit,
    cancel,
  };
};
