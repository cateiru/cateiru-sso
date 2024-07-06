-- ローカルで使用するためのスタッフアカウントを作成する
-- パスワードは設定されていないので再設定する必要がある

-- ユーザーid: `01H6P61PBFHP1PZP2ZWA2XNDH0`
-- ユーザー名 `admin`
-- メールアドレス: `admin@local.test`
INSERT INTO `user` (
    id,
    user_name,
    email
) VALUES (
    '01H6P61PBFHP1PZP2ZWA2XNDH0',
    'admin',
    'admin@local.test'
);

-- ユーザーを管理者にする
INSERT INTO `staff` (
    user_id
) VALUES (
    '01H6P61PBFHP1PZP2ZWA2XNDH0'
);
