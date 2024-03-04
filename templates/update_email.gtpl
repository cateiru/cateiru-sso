{{ .OldEmail }} から {{ .Email }} に更新しようとしています。
ブラウザで以下の{{ len .Code }}桁の確認コードを入力してください。
この確認コードの有効期限は{{ timeDiffMinutes .Expiration }}分です。

確認コード: {{.Code}}


もし、このメールに見に覚えがない場合は、
お手数ですが下記のリンクからお問い合わせください。

https://todo

---

@{{.BrandDomain}} #{{.Code}}
