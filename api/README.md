# Cateiru認証 API

## 環境変数

```env
# デプロイモード
# production or other
DEPLOY_MODE=
```

## テスト

1. DBを起動する

   ```bash
   docker-compose up -d
   ```

2. テスト実行

    ```bash
    make test
    ```

3. DB停止

    ```bash
    docker-compose down --rmi all
    ```
