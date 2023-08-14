{{.Data.BrandName}} へようこそ！

はじめに、 メールアドレスを認証する必要があります。このメールに含まれる {{len .Data.Data.Code}}桁のコードをブラウザで入力してください。
このコードの有効期限は{{timeDiffMinutes .Data.Data.Time}}分です。

確認コード: {{.Data.Data.Code}}


もし、このメールに見に覚えがない場合は、
お手数ですが下記のリンクからお問い合わせください。

https://todo

---

@{{.Data.BrandDomain}} #{{.Data.Data.Code}}
