{{.Data.BrandName}} へようこそ！

はじめに、 メールアドレスを認証する必要があります。このメールに含まれる {{len .Data.Data.Code}}桁のコードを認証ページに入力してください。
このコードの有効期限は{{timeDiffMinutes .Data.Data.Time}}分です。有効期限が切れると使用できなくなるので注意してください。

確認コード: {{.Data.Data.Code}}


もし、このメールに見に覚えがない場合は、
お手数ですが下記のリンクから報告いただければ幸いです。

https://todo

---

{{.Data.BrandEmail}}
`
