@{{ .Data.Data.InvitationUserName }} により組織{{ .Data.Data.OrganizationName }}に招待されています。
組織に参加する場合は、以下のリンクからアカウント作成を行ってください。
このリンクの有効期限は{{ timeDiffMinutes .Data.Data.Period }}分です。
リンクは一度アクセスすると再度アクセスすることはできなくなります。

{{ .Data.BrandUrl  }}/register?invite_token={{ .Data.Data.Token }}&email={{ .Data.Data.Email }}

もし、このメールに見に覚えがない場合は、
お手数ですが下記のリンクからお問い合わせください。

https://todo
