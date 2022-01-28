/**
 * @param {number} code - エラーコード
 * @returns {string} - エラーメッセージ
 */
export default function decodeErrorCode(code: number): string {
  switch (code) {
    // 0: Success
    // 1: DefaultError
    // 2: Response内でエラー
    // 3: ブロックリストに入っていたエラー
    // 4: メールアドレスがすでに存在しているエラー
    // 5: Bot判定したためエラー
    // 6: 時間切れ
    // 7: すでに認証済み
    // 8: アカウントない
    // 9: ログインできなかった
    case 2:
      return '予測不可能なエラー';
    case 3:
      return 'あなたは許可されていないようです';
    case 4:
      return 'そのメールアドレスはすでに登録されているようです';
    case 5:
      return 'あなたの操作がBotと判定されてしまいました';
    case 6:
      return '有効時間が切れてしまいました…もっと素早く行動しましょう';
    case 7:
      return 'あれ？すでに認証されているようです';
    case 8:
      return 'あれ？あなたのアカウントはありませんよ？';
    case 9:
      return 'そのユーザIDはすでに存在しています';
    case 10:
      return 'ログインができませんでした。パスワードまたはメールアドレスが間違っています。';
    default:
      return 'エラー';
  }
}
