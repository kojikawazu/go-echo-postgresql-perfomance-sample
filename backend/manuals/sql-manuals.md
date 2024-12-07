# SQLマニュアル

## 1. データベースに接続

```bash
PGPASSWORD=postgresql psql -h localhost -U postgresql -d postgresql
```

## 2. データベースに接続

```bash
psql -U postgresql -d postgresql
```

## 3. テーブル作成

```sql
CREATE TABLE samples (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
```

## 4. テーブル削除

```sql
DROP TABLE IF EXISTS samples;
```

## 5. データ全削除

```sql
DELETE FROM samples IF EXISTS samples;
```

## 6. 選択

```sql
SELECT * FROM samples;
```

## 7. 追加

```sql
INSERT INTO samples (name) VALUES ('sample');
```

## 8. 更新

```sql
UPDATE samples SET name = 'sample' WHERE id = 1;
```

## 9. 削除

```sql
DELETE FROM samples WHERE id = 1;
```

## 10. インデックス作成

```sql
CREATE INDEX idx_samples_name ON samples (name);
```

## 11. インデックス削除

```sql
DROP INDEX IF EXISTS idx_samples_name;
```

## 12. UUID拡張機能を有効にする

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

## 13. UUIDを生成

```sql
SELECT uuid_generate_v4();
```

## 14. UUIDをカラムに追加

```sql
ALTER TABLE samples ADD COLUMN id UUID PRIMARY KEY DEFAULT uuid_generate_v4();
```

## 15. UUIDをカラムから削除

```sql
ALTER TABLE samples DROP COLUMN id;
```

## 16. テーブルのカラムを変更

```sql
ALTER TABLE samples ALTER COLUMN name TYPE VARCHAR(100);
```

## 17. テーブルのカラムを削除

```sql
ALTER TABLE samples DROP COLUMN name;
```

## 18. テーブルのカラムを追加

```sql
ALTER TABLE samples ADD COLUMN name VARCHAR(255) NOT NULL;
```

## 19. テーブルのカラムを変更

```sql
ALTER TABLE samples ALTER COLUMN name TYPE VARCHAR(100);
```

## 20. 汎用計測関数の作成

```sql
CREATE OR REPLACE FUNCTION measure_execution_time(sql_query text)
RETURNS TABLE(duration_ms double precision, result jsonb) AS $$
DECLARE
    start_time timestamp;
    end_time timestamp;
    query_result jsonb;
BEGIN
    -- 開始時刻を記録
    start_time := clock_timestamp();

    -- 渡されたクエリを実行
    EXECUTE sql_query INTO query_result;

    -- 終了時刻を記録
    end_time := clock_timestamp();

    -- 実行時間をミリ秒で計算
    RETURN QUERY SELECT EXTRACT(MILLISECOND FROM end_time - start_time) AS duration_ms, query_result AS result;
END;
$$ LANGUAGE plpgsql;
```

### 使用例

```sql
SELECT * FROM measure_execution_time('SELECT pg_sleep(10);');
```

```sql
# 計測結果
duration_ms | result
------------|---------
10000       | null
```

### 特定の操作の実行時間を計測

```sql
SELECT * FROM measure_execution_time('INSERT INTO users (username, email) VALUES (\'sampleuser\', \'sample@example.com\');');
```
