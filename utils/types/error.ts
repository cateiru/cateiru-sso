import z from 'zod';

export const ErrorSchema = z.object({
  message: z.string(),
  unique_code: z.number().optional(),
});
export type ErrorType = z.infer<typeof ErrorSchema>;

export const ErrorUniqueMessage: {[key: number]: string | undefined} = {
  0: undefined,
  1: 'reCAPTCHAの認証に失敗しました',
  2: 'このEmailは現在登録には使用できません',
  3: 'このアカウントは作成できません',
  4: 'リトライの上限を超えました',
  5: 'セッションの有効期限が切れました',
  6: 'メールの送信上限を超えました',
  7: 'このメールアドレスはまだ認証されていません',
  8: 'ログインに失敗しました',
  9: 'ログインに失敗しました。しかし、別のアカウントでログインできる可能性があります',
  10: 'このユーザーは存在しません',
  11: 'このユーザーはパスワードを設定していません',
  12: 'このユーザーはすでに存在しています',
  13: '認証に失敗しました',
  14: 'パスワードが設定されていないため、生体認証をすべて削除することはできません', // WebAuthnの削除でのみ使用している
  15: 'すでにログイン済みです',
  16: 'この組織に加入されていないようです',
  17: 'この組織の詳細にアクセスできる権限がありません',
};

export const OidcErrorSchema = z.object({
  error: z.string(),
  error_description: z.string().optional(),
  error_uri: z.string().optional(),
  state: z.string().optional(),
});
export type OidcErrorType = z.infer<typeof OidcErrorSchema>;
