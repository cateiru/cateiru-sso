import {OIDCRequestQuery} from '../sso/login';
import {API} from './api';

export interface ServicePreview {
  name: string;
  service_icon: string;
}

export interface ServiceLogin {
  access_token: string;
}

export const preview = async (oidc: OIDCRequestQuery, fromURI: string) => {
  const api = new API();

  api.post(
    JSON.stringify({
      scope: oidc.scopes,
      response_type: oidc.responseType.join(' '),
      client_id: oidc.clientID,
      redirect_uri: oidc.redirectURL,
      state: oidc.state,
      prompt: oidc.prompt,
      from_url: fromURI,
    })
  );

  const resp = await api.connect('/oauth/preview');

  return (await resp.json()) as ServicePreview;
};

export const login = async (
  oidc: OIDCRequestQuery,
  fromURI: string
): Promise<ServiceLogin> => {
  const api = new API();

  api.post(
    JSON.stringify({
      scope: oidc.scopes,
      response_type: oidc.responseType.join(' '),
      client_id: oidc.clientID,
      redirect_uri: oidc.redirectURL,
      state: oidc.state,
      prompt: oidc.prompt,
      from_url: fromURI,
    })
  );

  const resp = await api.connect('/oauth/login');

  return (await resp.json()) as ServiceLogin;
};
