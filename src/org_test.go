package src_test

import "testing"

func TestOrgGetHandler(t *testing.T) {
	t.Run("成功", func(t *testing.T) {})
	t.Run("成功: どこにも所属していない場合は空の配列が返る", func(t *testing.T) {})
}

func TestOrgGetMemberHandler(t *testing.T) {
	t.Run("成功: orgのメンバー一覧が返る", func(t *testing.T) {})

	t.Run("失敗: org_idが空", func(t *testing.T) {})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})
}

func TestOrgPostMemberHandler(t *testing.T) {
	t.Run("成功: orgのメンバーを招待できる", func(t *testing.T) {})

	t.Run("成功: roleが空の場合はguestになる", func(t *testing.T) {})

	t.Run("失敗: org_idが空", func(t *testing.T) {})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})

	t.Run("失敗: user_idが空", func(t *testing.T) {})

	t.Run("失敗: user_idが不正", func(t *testing.T) {})

	t.Run("失敗: roleが不正", func(t *testing.T) {})

	t.Run("失敗: ユーザーはすでにメンバーになっている", func(t *testing.T) {})
}

func TestOrgUpdateMemberHandler(t *testing.T) {
	t.Run("成功: roleを変更できる", func(t *testing.T) {})

	t.Run("失敗: org_user_idが空", func(t *testing.T) {})

	t.Run("失敗: org_user_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})

	t.Run("失敗: roleが不正", func(t *testing.T) {})

	t.Run("失敗: 自分自身のroleは変更できない", func(t *testing.T) {})
}

func TestOrgDeleteMemberHandler(t *testing.T) {
	t.Run("成功: orgからユーザーを削除できる", func(t *testing.T) {})

	t.Run("失敗: org_user_idが空", func(t *testing.T) {})

	t.Run("失敗: org_user_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})

	t.Run("失敗: 自分自身は削除できない", func(t *testing.T) {})
}

func TestOrgInvitedMemberHandler(t *testing.T) {
	t.Run("成功: 招待中の一覧を取得できる", func(t *testing.T) {})

	t.Run("失敗: org_idが空", func(t *testing.T) {})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})
}

func TestOrgInviteNewMemberHandler(t *testing.T) {
	t.Run("成功: 対象のEmailに対して招待できる", func(t *testing.T) {})

	t.Run("失敗: org_idが空", func(t *testing.T) {})

	t.Run("失敗: org_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: emailが空", func(t *testing.T) {})

	t.Run("失敗: emailの値が不正", func(t *testing.T) {})

	t.Run("失敗: そのemailのユーザーはすでに存在している", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})
}

func TestOrgInviteMemberDeleteHandler(t *testing.T) {
	t.Run("成功: 招待をキャンセルできる", func(t *testing.T) {})

	t.Run("失敗:invite_idが空", func(t *testing.T) {})

	t.Run("失敗:invite_idの値が不正", func(t *testing.T) {})

	t.Run("失敗: orgに所属していない", func(t *testing.T) {})

	t.Run("失敗: orgに所属しているけどownerではない", func(t *testing.T) {})
}
