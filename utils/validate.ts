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
