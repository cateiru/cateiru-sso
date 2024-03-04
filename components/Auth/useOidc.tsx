import {useSearchParams} from 'next/navigation';
import {OidcParams, OidcParamsSchema} from '../../utils/types/oidc';

export const useOidc = () => {
  const searchParams = useSearchParams();

  const getParams = (): OidcParams => {
    const data = OidcParamsSchema.safeParse({
      scope: searchParams.get('scope'),
      response_type: searchParams.get('response_type'),
      client_id: searchParams.get('client_id'),
      redirect_uri: searchParams.get('redirect_uri'),
      state: searchParams.get('state'),
      response_mode: searchParams.get('response_mode'),
      nonce: searchParams.get('nonce'),
      display: searchParams.get('display'),
      prompt: searchParams.get('prompt'),
      max_age: searchParams.get('max_age'),
      ui_locales: searchParams.get('ui_locales'),
      id_token_hint: searchParams.get('id_token_hint'),
      login_hint: searchParams.get('login_hint'),
      acr_values: searchParams.get('acr_values'),
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
