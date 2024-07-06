export function validateGender(gender: string): string {
  switch (gender) {
    case '0':
      return '不明';
    case '1':
      return '男性';
    case '2':
      return '女性';
    case '9':
      return 'その他';
    default:
      return '不明';
  }
}

export function validatePrompt(prompt: string | null): string {
  if (prompt === null) {
    return '認証しない';
  }

  switch (prompt) {
    case 'login':
      return '認証を求める';
    case '2fa_login':
      return '二段階認証のみを求める';
    default:
      return '認証しない';
  }
}

export function validateScope(scope: string): string {
  switch (scope) {
    case 'openid':
      return 'openid';
    case 'profile':
      return 'プロフィール';
    case 'email':
      return 'メールアドレス';
    default:
      return '不明';
  }
}
