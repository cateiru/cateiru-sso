interface Config {
  // APIのホスト
  // undefined の場合は相対パスになる
  apiHost?: string;
  // APIのパスのプレフィックス
  apiPathPrefix?: string;
  apiCors?: boolean;
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
  apiHost: process.env.NEXT_PUBLIC_API_HOST,
  apiPathPrefix: '/api/v2',
  apiCors: false,
  loginStateCookieName: 'oreore-me-login-state',
  emailCodeLength: 6,
  mode: process.env.NODE_ENV ?? 'development',
  title: `${
    process.env.NEXT_PUBLIC_PUBLICATION_TYPE
      ? `${process.env.NEXT_PUBLIC_PUBLICATION_TYPE}.`
      : ''
  }oreore.me`,
  revision: process.env.NEXT_PUBLIC_REVISION ?? '-',
  branchName: process.env.NEXT_PUBLIC_BRANCH_NAME ?? '-',
  reCaptchaKey: process.env.NEXT_PUBLIC_RE_CAPTCHA,
  gaTracingId: process.env.NEXT_PUBLIC_GOOGLE_ANALYTICS_ID,
};
