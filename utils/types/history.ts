import {z} from 'zod';

export const LoginDeviceScheme = z.object({
  id: z.number(),

  device: z.string().nullable(),
  os: z.string().nullable(),
  browser: z.string().nullable(),
  is_mobile: z.boolean().nullable(),
  ip: z.string(),

  is_current: z.boolean(),

  created: z.string().datetime({offset: true}),
});
export type LoginDevice = z.infer<typeof LoginDeviceScheme>;

export const LoginDeviceListScheme = z.array(LoginDeviceScheme);
export type LoginDeviceList = z.infer<typeof LoginDeviceListScheme>;

export const LoginTryHistoryScheme = z.object({
  id: z.number(),

  device: z.string().nullable(),
  os: z.string().nullable(),
  browser: z.string().nullable(),
  is_mobile: z.boolean().nullable(),
  ip: z.string(),

  // 識別子
  // 0: ログイン
  // 1: パスワード再登録
  identifier: z.number(),

  created: z.string().datetime({offset: true}),
});
export type LoginTryHistory = z.infer<typeof LoginTryHistoryScheme>;

export const LoginTryHistoryListScheme = z.array(LoginTryHistoryScheme);
export type LoginTryHistoryList = z.infer<typeof LoginTryHistoryListScheme>;

export const LOGIN_TRY_IDENTIFIER: {[key: number]: string} = {
  0: 'ログイン',
  1: 'パスワード再登録',
};
