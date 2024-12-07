# PostgreSQL Docker環境構築手順書

## 1. 事前確認と準備

```bash
# ホスト側のPostgreSQLの状態を確認
sudo service postgresql status

# もし動いていれば停止（重要！）
sudo service postgresql stop

# ポート5432が使用されていないことを確認
sudo lsof -i :5432
```

## 2. Docker環境の設定

```yml
services:
  postgres:
    image: postgres:15
    container_name: postgres_container
    environment:
      POSTGRES_USER: postgresql
      POSTGRES_PASSWORD: postgresql
      POSTGRES_DB: postgresql
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - postgres_network

networks:
  postgres_network:
    driver: bridge

volumes:
  postgres_data:
```

## 3. コンテナの起動

```bash
# コンテナの起動
docker compose up -d

# ログを確認
docker logs postgres_container
```

## 4. PostgreSQL設定の確認と更新

```bash
# コンテナに入る
docker exec -it postgres_container bash

# PostgreSQLに接続
psql -U postgresql -d postgresql

# パスワードを設定
ALTER USER postgresql WITH PASSWORD 'postgresql';

# 設定を再読み込み
SELECT pg_reload_conf();

# 終了
\q
```

## 5. 確認

```bash
# ポート確認
sudo lsof -i :5432

# PostgreSQL接続テスト
PGPASSWORD=postgresql psql -h localhost -U postgresql -d postgresql

# コンテナに入る
docker exec -it postgres_container /bin/bash

# PostgreSQL接続テスト
psql -U postgresql -d postgresql
```

