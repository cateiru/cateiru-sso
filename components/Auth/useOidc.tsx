import {useSearchParams} from 'next/navigation';
import {OidcParams, OidcParamsSchema} from '../../utils/types/oidc';

export const useOidc = () => {
  const searchParams = useSearchParams();

  const getParams = (): OidcParams => {
    const scope = searchParams.get('scope');
    const responseType = searchParams.get('response_type');
    const clientId = searchParams.get('client_id');
    const redirectUri = searchParams.get('redirect_uri');
    const state = searchParams.get('state');

    const responseMode = searchParams.get('response_mode');
    const nonce = searchParams.get('nonce');
    const display = searchParams.get('display');
    const prompt = searchParams.get('prompt');
    const maxAge = searchParams.get('max_age');
    const uiLocales = searchParams.get('ui_locales');
    const idTokenHint = searchParams.get('id_token_hint');
    const loginHint = searchParams.get('login_hint');
    const acrValues = searchParams.get('acr_values');

    const data = OidcParamsSchema.safeParse({
      scope,
      response_type: responseType,
      client_id: clientId,
      redirect_uri: redirectUri,
      state,
      response_mode: responseMode,
      nonce,
      display,
      prompt,
      max_age: maxAge,
      ui_locales: uiLocales,
      id_token_hint: idTokenHint,
      login_hint: loginHint,
      acr_values: acrValues,
    });

    if (data.success) {
      return data.data;
    }

    throw data.error;
  };

  const getFormParams = (): FormData => {
    const param = getParams();

    const form = new FormData();

    for (const key in param) {
      const value = param[key as keyof OidcParams];
      if (typeof value === 'string') {
        form.append(key, value);
      }
    }

    return form;
  };

  return {
    getParams,
    getFormParams,
  };
};
