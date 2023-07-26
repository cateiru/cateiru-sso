interface Config {
  apiHost: string;
  loginStateCookieName: string;
  emailCodeLength: number;
  mode: 'development' | 'production' | 'test';
  title: string;
  revision: string;
  branchName?: string;
  reCaptchaKey?: string;
  gaTracingId?: string;
}

export const config: Config = {
  apiHost: process.env.NEXT_PUBLIC_API_HOST ?? 'http://localhost:8080',
  loginStateCookieName: 'oreore-me-login-state',
  emailCodeLength: 6,
  mode: process.env.NODE_ENV ?? 'development',
  title: `${
    process.env.NEXT_PUBLIC_PUBLICATION_TYPE
      ? `${process.env.NEXT_PUBLIC_PUBLICATION_TYPE}.`
      : ''
  }oreore.me`,
  revision: process.env.NEXT_PUBLIC_REVISION ?? 'unknown',
  branchName: process.env.NEXT_PUBLIC_BRANCH_NAME,
  reCaptchaKey: process.env.NEXT_PUBLIC_RE_CAPTCHA,
  gaTracingId: process.env.NEXT_PUBLIC_GOOGLE_ANALYTICS_ID,
};
