{{ .Data.Data.UserName }}のパスワードを再設定しようとしています。
再設定する場合は、以下のリンクから再設定をしてください。
リンクの有効期限は{{ timeDiffMinus .Data.Data.PeriodTime }}分です。

{{ .Data.BrandUrl  }}/forget_password/reregister?token={{ .Data.Data.SessionToken }}&email={{ .Data.Data.Email }}

もし、このメールに見に覚えがない場合は、
お手数ですが下記のリンクからお問い合わせください。

https://todo

---

@{{.Data.BrandDomain}} #{{.Data.Data.Code}}
