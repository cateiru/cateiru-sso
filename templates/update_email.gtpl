{{ .EmailAddress }}から{{ .Data.Data.NewEmail }}に更新しようとしています。
ブラウザで以下の{{ len .Data.Data.Code }}桁の確認コードを入力してください。
この確認コードの有効期限は{{ timeDiffMinus .Data.Data.Period }}分です。

確認コード: {{.Data.Data.Code}}


もし、このメールに見に覚えがない場合は、
お手数ですが下記のリンクからお問い合わせください。

https://todo

---

@{{.Data.BrandDomain}} #{{.Data.Data.Code}}
