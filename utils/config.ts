import {URL} from 'url';

interface Config {
  apiHost: string;
  loginStateCookieName: string;
}

export const config: Config = {
  apiHost: process.env.NEXT_PUBLIC_API_HOST ?? 'http://localhost:8080',
  loginStateCookieName: 'cateiru-sso-login-state',
};
