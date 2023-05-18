interface Config {
  apiHost: string;
  loginStateCookieName: string;
  registerAccountEmailCodeLength: number;
  mode: 'development' | 'production' | 'test';
  title: string;
}

export const config: Config = {
  apiHost: process.env.NEXT_PUBLIC_API_HOST ?? 'http://localhost:8080',
  loginStateCookieName: 'cateiru-sso-login-state',
  registerAccountEmailCodeLength: 6,
  mode: process.env.NODE_ENV ?? 'development',
  title: `CateiruSSO${
    process.env.NEXT_PUBLIC_PUBLICATION_TYPE
      ? ` - ${process.env.NEXT_PUBLIC_PUBLICATION_TYPE}`
      : ''
  }`,
};
