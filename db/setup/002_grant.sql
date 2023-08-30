GRANT ALL ON *.* TO `docker` @`%`;

# caching_sha2_password だとテスト時にエラーになるためmysql_native_passwordに変更する
# ref. https://yoku0825.blogspot.com/2018/10/mysql-80cachingsha2password-ssl.html
# なお、このデータベースは外に公開するものではないためセキュリティ的には問題ないと思われる
ALTER USER `docker` @`%` IDENTIFIED WITH mysql_native_password BY 'docker';
