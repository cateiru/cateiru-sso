#!/bin/bash

DB_USER="docker"
DB_PASSWORD="docker"

TARGET_DB="diff_target"
CURRENT_DB="diff_current"

MIGRATIONS_PATH="/migrations"

setup_and_cleanup() {
    # db をリセットする
    mysql -u $DB_USER -p$DB_PASSWORD -h db -e "DROP DATABASE IF EXISTS $TARGET_DB; CREATE DATABASE IF NOT EXISTS $TARGET_DB;";
    mysql -u $DB_USER -p$DB_PASSWORD -h db -e "DROP DATABASE IF EXISTS $CURRENT_DB; CREATE DATABASE IF NOT EXISTS $CURRENT_DB;";
}

dump() {
    DB_NAME=$1
    rm -rf /dump_data/$DB_NAME.sql
    mysqldump -u $DB_USER -p$DB_PASSWORD -h db --no-data --skip-add-drop-table --compact $DB_NAME > /dump_data/$DB_NAME.sql
}

setup_and_cleanup;

# current にマイグレーションする
MYSQL_DSN="mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp(db:3306)/$CURRENT_DB"
migrate -path $MIGRATIONS_PATH -database "$MYSQL_DSN" up
# go-migrate の schema_migrations テーブルを削除する
mysql -u $DB_USER -p$DB_PASSWORD -h db -e "DROP TABLE $CURRENT_DB.schema_migrations;"

# target に schema.sql をマイグレーションする
mysql -u $DB_USER -p$DB_PASSWORD -h db $TARGET_DB < /schema.sql

# dumpする
dump $CURRENT_DB
dump $TARGET_DB

UP=$(mysqldef -u $DB_USER -p$DB_PASSWORD -h db $CURRENT_DB --enable-drop-table --dry-run < /dump_data/$TARGET_DB.sql | grep -v '^-- dry run --$')
DOWN=$(mysqldef -u $DB_USER -p$DB_PASSWORD -h db $TARGET_DB --enable-drop-table --dry-run < /dump_data/$CURRENT_DB.sql | grep -v '^-- dry run --$')

TIMESTAMP=$(date "+%Y%m%d%H%M%S")
MIGRATION_FILE_NAME=$1

UP_FILE_PATH="$MIGRATIONS_PATH/${TIMESTAMP}_$MIGRATION_FILE_NAME.up.sql"
DOWN_FILE_PATH="$MIGRATIONS_PATH/${TIMESTAMP}_$MIGRATION_FILE_NAME.down.sql"

echo "$UP" > $UP_FILE_PATH
echo "$DOWN" > $DOWN_FILE_PATH

echo "\n\nCreated: ${TIMESTAMP}_$MIGRATION_FILE_NAME.(up|down).sql"
