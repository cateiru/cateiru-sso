# テスト用データベースに同じテーブルをマイグレーションする
mysql -uroot -proot cateiru-sso-test < "/docker-entrypoint-initdb.d/001_schema.sql"
