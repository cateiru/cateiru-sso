interface Config {
  apiHost: string;
  loginStateCookieName: string;
  registerAccountEmailCodeLength: number;
}

export const config: Config = {
  apiHost: process.env.NEXT_PUBLIC_API_HOST ?? 'http://localhost:8080',
  loginStateCookieName: 'cateiru-sso-login-state',
  registerAccountEmailCodeLength: 6,
};
