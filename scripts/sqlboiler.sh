#!/bin/bash

# sqlboilerコマンド
# 詳しくは https://github.com/volatiletech/sqlboiler#initial-generation 参照
# ついでに https://zenn.dev/gami/articles/0fb2cf8b36aa09 も参照

sqlboiler mysql -c db/sqlboiler.toml -o src/models -p models --wipe
