# 実行計画

## DBのパフォーマンスチューニング

### 1. 実行計画を理解し、クエリの動作を把握する

- **目的**: クエリがどのようにデータを取得しているか、スキャン手法や結合の仕方を確認。
- **手段**:
  - EXPLAIN や EXPLAIN ANALYZE を使用してクエリの実行計画を解析。
  - ボトルネックとなる箇所（Seq Scan、Nested Loopなど）を特定。
- **ポイント**:
  - 適切なインデックスが使用されているか。
  - 不必要なスキャンやフィルタが発生していないか。

### 2. Web APIやDBの計測でパフォーマンスデータを収集する

- **目的**: 実際の負荷やレスポンス時間を数値化して、現状を把握。
- **手段**:
  - Web API:
    - 各エンドポイントでの処理時間を測定し、ログに記録。
    - 各処理（DBアクセス、ビジネスロジック、ネットワーク遅延）を分解して分析。
  - DB:
    - クエリの実行時間を記録 (pg_stat_statements や slow_query_log を活用)。
    - 高頻度で実行されるクエリを特定。
- **ポイント**:
  - クエリの実行回数が多い場合、キャッシュの導入を検討。
  - 実行回数が少なくても時間がかかるクエリに注目。

### 3. 優先順位をつけて対応する

- **目的**: リソースや時間を効率的に活用し、最大の成果を得る。
- **手段**:
  - 高コストクエリ:
    - 実行計画で最もコストが高い操作（例: Seq Scan）を優先的に改善。
  - 頻度の高いクエリ:
    - アクセス頻度が多いテーブルやエンドポイントのクエリを最適化。
- **ポイント**:
  - トレードオフを考慮し、リスクや効果を比較して取り組む。

### 4. 具体的な対応方法

- **インデックスの最適化**:
  - 適切な列にインデックスを追加（B-tree、GIN、Partial Indexなど）。
- **クエリのリファクタリング**:
  - 冗長な結合や不要なカラムの取得を削除。
  - サブクエリをCTEやJOINに変更して効率化。
- **キャッシュの利用**:
  - 結果セットをメモリやRedisにキャッシュ。
- **パーティショニングやアーカイブ**:
  - データ量が膨大な場合、パーティショニングや古いデータのアーカイブを検討。

## 実行計画結果例

### Nested Loop Left Join

- **概要**:
  - projectsテーブルとtasksテーブルを結合する際に使用される手法。
  - 結合条件はprojects.id = tasks.project_idです。
- **特徴**:
  - projectsテーブルから1行ずつ取得し、それに対応するtasksを検索します。
  - 小規模なデータセットでは効果的。
  - **例**: プロジェクト数が少なく、各プロジェクトに紐づくタスクが少数の場合。
  - 大規模なデータセットでは非効率になる可能性があります。
    - 大量のデータを処理する際は、`Hash Join` や `Merge Join` を検討。

```sql
EXPLAIN ANALYZE
SELECT
    projects.id,
    projects.name,
    tasks.id AS task_id,
    tasks.name AS task_name
FROM
    projects
LEFT JOIN tasks ON tasks.project_id = projects.id;
```

```sql:結果例
Nested Loop Left Join  (cost=0.00..104.00 rows=100 width=64) (actual time=0.030..0.350 rows=100 loops=1)
   -> Seq Scan on projects  (cost=0.00..2.00 rows=10 width=32) (actual time=0.010..0.020 rows=10 loops=1)
   -> Index Scan on tasks  (cost=0.00..10.00 rows=10 width=32) (actual time=0.020..0.030 rows=10 loops=10)
```

- 結果解説:
  - Seq Scan on projects
    - projectsテーブル全体をシーケンシャルスキャンしています。
    - 10行をスキャンし、1行ずつ処理を進めています。
  - Index Scan on tasks
    - tasksテーブルに対して、project_id に基づくインデックススキャンを実施。
    - 10回のループで各プロジェクトに関連するタスクを検索。
  - Nested Loop Left Join
    - projectsの各行に対して、tasksから該当するデータを結合。
    - すべてのタスクが見つかるまで繰り返し処理を実行。

### Seq Scan on projects

- **概要**:
  - projects テーブルをシーケンシャルスキャンで走査。
  - テーブル全体を順番に確認し、フィルター条件に一致する行を探します。
- **フィルター条件**:
  - name LIKE '%example%'
- **結果**:
  - 条件に一致する10行が取得されました。
  - フィルター条件によって他の行は除外されています。
- **注意点**:
  - テーブルのデータ量が大きい場合、シーケンシャルスキャンはコストが高くなります。
  - 条件にインデックスが利用可能であれば、Index Scan に切り替える方が効率的です。

```sql
EXPLAIN ANALYZE
SELECT * FROM projects WHERE name LIKE '%example%';
```

```sql:結果例
Seq Scan on projects  (cost=0.00..10.00 rows=10 width=32) (actual time=0.010..0.020 rows=10 loops=1)
   Filter: (name ~~ 'example%'::text)
```

- 結果解説:
  - Seq Scan on projects:
    - テーブル全体をスキャンし、name 列の値が LIKE '%example%' に一致するかを確認。
    - 条件に一致した行のみ返されます。
  - Filter:
    - LIKE '%example%' の条件はインデックスを利用できないため、すべての行を評価しています。
  - 注意:
    - フィルタリング条件を変更してインデックスを活用できる場合、Index Scan により効率化できます。
      - 例: name LIKE 'example%' の場合、B-Tree インデックスを使用可能。

### Bitmap Heap Scan on tasks

- **概要**:
  - tasks テーブルから、project_id = 'uuid' の条件に合致する行を効率的に取得するためのスキャン手法です。
  - フィルタリング条件に対応するインデックスを使用し、必要なデータブロックにアクセスします。
- **条件**:
  - project_id = 'uuid'
- **特徴**:
  - インデックス (Bitmap Index Scan) を活用して条件に合う行を特定し、その後に必要なデータを取得するためにテーブル (Heap) にアクセスします。
  - Bitmap Heap Scan は、特定の条件で対象行が多い場合に特に有効です。
- **結果**:
  - 該当する 1000行 が取得されました。
  - 実行計画に示された Heap Blocks: exact=22 は、22個のデータブロックが効率的に読み込まれたことを意味します。

```sql
EXPLAIN ANALYZE
SELECT * FROM tasks WHERE project_id = 'uuid';
```

```sql:結果例
Bitmap Heap Scan on tasks  (cost=0.00..10.00 rows=10 width=32) (actual time=0.010..0.020 rows=10 loops=1)
   Recheck Cond: (project_id = 'uuid'::uuid)
   -> Bitmap Index Scan on idx_project_status  (cost=0.00..0.00 rows=10 width=0) (actual time=0.010..0.010 rows=10 loops=1)
         Index Cond: (project_id = 'uuid'::uuid)
```

- 結果解説:
  - Bitmap Heap Scan on tasks:
    - インデックスで取得された条件に合致する行が、該当するデータブロックから効率的に取得されます。
  - Heap Blocks: exact=22 はアクセスしたデータブロック数を示しています。
  - Recheck Cond:
    - インデックスで絞り込まれた条件を再チェックして正確な行を取得します。
  - Bitmap Index Scan on idx_project_status:
    - tasks テーブルの project_id 列に設定されたインデックス (idx_project_status) を使用したスキャン。
    - Index Cond に示された条件 (project_id = 'uuid') に一致する行を検索します。

### Bitmap Index Scan on idx_project_status

- **概要**:
  - tasks テーブルの project_id 列に設定されたインデックスを使用して、条件に一致する行を特定します。
  - インデックススキャンは、条件に合う行が特定の列に集中している場合に特に効果的です。
- **インデックス条件**:
  - project_id = 'uuid'
- **結果**:
  - Index Cond で指定された条件に基づき、1000行が該当しました。
- **メリット**:
  - テーブル全体をスキャンするシーケンシャルスキャンと比較して、特定条件に基づくデータ取得が高速化します。
  - データ量が増加しても、インデックスが有効であれば検索パフォーマンスが保たれます。

```sql
EXPLAIN ANALYZE
SELECT * FROM tasks WHERE project_id = 'uuid';
```

```sql:結果例
Bitmap Index Scan on idx_project_status  (cost=0.00..0.00 rows=10 width=0) (actual time=0.010..0.010 rows=10 loops=1)
   Index Cond: (project_id = 'uuid'::uuid)
```

- 結果解説:
  - Bitmap Index Scan on idx_project_status:
    - インデックススキャンを実行し、条件 (project_id = 'uuid') に一致する行を特定します。
    - スキャン結果が1000行を超える場合、後続の Bitmap Heap Scan によるデータブロックアクセスが発生します。
  - Index Cond:
    - インデックスで使用された条件 (project_id = 'uuid') が表示されます。
  - コストと実行時間:
    - cost=0.00..0.00: インデックススキャンのコスト。
    - actual time=0.010..0.010: 実際のスキャン時間。

### Hash Join

- **概要**:
  - 大規模なデータセットの結合に使用される手法。
  - 片方のテーブルをハッシュ化し、もう片方と比較して結合。
- **特徴**:
  - データ量が大きい場合に効率的。
  - 小規模なデータセットではNested Loopよりもコストが高い可能性。
- **使用場面**:
  - 両方のテーブルが比較的大きい場合。
- **最適化**:
  - 結合対象にインデックスを追加することで、Index Nested Loop Join に切り替わる場合があります。

```sql
EXPLAIN ANALYZE
SELECT
    tasks.id,
    users.username
FROM
    tasks
JOIN users ON tasks.assigned_to = users.id;
```

```sql:結果例
Hash Join  (cost=10.00..120.00 rows=100 width=64) (actual time=0.020..0.080 rows=100 loops=1)
   Hash Cond: (tasks.assigned_to = users.id)
   -> Seq Scan on tasks  (cost=0.00..80.00 rows=100 width=32) (actual time=0.010..0.050 rows=100 loops=1)
   -> Hash  (cost=10.00..10.00 rows=10 width=32) (actual time=0.010..0.020 rows=10 loops=1)
         Buckets: 1024  Batches: 1  Memory Usage: 9kB
         -> Seq Scan on users  (cost=0.00..10.00 rows=10 width=32) (actual time=0.002..0.010 rows=10 loops=1)
```

### Merge Join

- **概要**:
  - Hash Join は、大規模なデータセットの結合に適した手法です。
  - 片方のテーブルをハッシュ化（メモリ上でハッシュテーブルを作成）し、もう片方のテーブルと比較して結合を行います。
- **特徴**:
  - テーブルサイズが大きい場合に効率的。
  - 結合条件が等値 (=) である場合に最適。
  - メモリに依存するため、メモリ容量が少ない場合はディスクI/Oが増える可能性があります。
- **使用場面**:
  - 両方のテーブルが比較的大きく、インデックスを使用できない場合。
  - 並列処理が可能な環境。
- **最適化**:
  - 結合対象のカラムにインデックスを設定することで、Index Nested Loop Join への切り替えを検討。
  - ハッシュテーブルのサイズを適切に設定するためにメモリを最適化。

```sql
EXPLAIN ANALYZE
SELECT
    tasks.id,
    users.username
FROM
    tasks
JOIN users ON tasks.assigned_to = users.id;
```

```sql:結果例
Merge Join  (cost=10.00..120.00 rows=100 width=64) (actual time=0.020..0.080 rows=100 loops=1)
   Merge Cond: (tasks.assigned_to = users.id)
   -> Seq Scan on tasks  (cost=0.00..80.00 rows=100 width=32) (actual time=0.010..0.050 rows=100 loops=1)
   -> Sort  (cost=10.00..10.00 rows=10 width=32) (actual time=0.010..0.020 rows=10 loops=1)
         Sort Key: users.id
         Sort Method: quicksort  Memory: 25kB
         -> Seq Scan on users  (cost=0.00..10.00 rows=10 width=32) (actual time=0.002..0.010 rows=10 loops=1)
```

- 結果解説:
  - Hash Join:
    - tasks.assigned_to = users.id の条件で結合を実行。
    - 両テーブルのデータ量に基づいて、ハッシュテーブルが作成されます。
  - Hash Cond:
    - ハッシュ結合で使用された条件 (tasks.assigned_to = users.id) が明示されます。
  - Seq Scan on tasks:
    - tasks テーブルをシーケンシャルスキャンで走査。
    - 条件がないため、すべての行をスキャンします。
  - Hash:
    - users テーブルを走査し、users.id 列を基にハッシュテーブルを作成。
    - Buckets（ハッシュバケット数）と Memory Usage（メモリ使用量）が示されています。

### Index Scan

- **概要**:
  - Index Scan は、テーブルの特定の行を効率的に取得するためにインデックスを使用する手法です。
  - インデックスがクエリのフィルタリング条件に適合している場合、高速にデータを取得できます。
- **特徴**:
  - 主キーやユニークインデックスに基づいた検索では非常に効率的。
  - 必要な行のみを取得するため、テーブル全体のスキャンを回避します。
- **使用場面**:
  - 主キー検索 (id = 'uuid')。
  - 特定の条件でデータを絞り込む場合。
- **最適化**:
  - クエリで頻繁に使用される列にインデックスを設定。
  - フィルタリング条件を具体的に指定することで、インデックスの恩恵を最大化。

```sql
EXPLAIN ANALYZE
SELECT * FROM users WHERE id = 'uuid';
```

```sql:結果例
Index Scan using users_pkey on users  (cost=0.00..1.00 rows=1 width=64) (actual time=0.010..0.020 rows=1 loops=1)
   Index Cond: (id = 'uuid')
``` 

- 結果解説:
  - Index Scan using users_pkey:
    - users テーブルの主キーインデックス (users_pkey) を利用。
    - 主キー id を条件に、効率的に1行を取得しています。
  - Index Cond:
    - 条件 id = 'uuid' が適用されていることを示します。
  - 実行時間:
    - actual time=0.010..0.020 により、インデックスを利用して高速に行が取得されていることが分かります。
  - コスト:
    - cost=0.00..1.00 は、非常に低コストでクエリが実行されることを示します。

### Index Only Scan

- **概要**:
  - Index Only Scan は、必要なデータがすべてインデックスに含まれる場合に使用されるスキャン手法です。
  - テーブル本体へのアクセスを回避し、インデックスのみからデータを取得します。
- **特徴**:
  - 実際のテーブルデータにアクセスしないため、高速。
  - 主に読み取り専用のクエリに適しています。
- **使用場面**:
  - クエリで指定される列がインデックスに含まれる場合。
  - テーブルサイズが大きいが、少数の列のみを取得したい場合。
- **最適化**:
  - インデックスに必要最小限の列を含める。
  - 必要に応じてカバリングインデックス（すべての参照列を含むインデックス）を設計。

```sql
EXPLAIN ANALYZE
SELECT id FROM users WHERE created_at > '2024-01-01';
```

```sql:結果例
Index Only Scan using idx_users_created_at on users  (cost=0.00..10.00 rows=50 width=16) (actual time=0.010..0.030 rows=50 loops=1)
   Index Cond: (created_at > '2024-01-01')
   Heap Fetches: 0
```

- 結果解説:
  - Index Only Scan using idx_users_created_at:
    - users テーブルの created_at 列に設定されたインデックス idx_users_created_at を使用。
    - 条件 created_at > '2024-01-01' に基づいてインデックスのみからデータを取得しています。
  - Index Cond:
    - インデックス条件として (created_at > '2024-01-01') が適用されていることを示します。
  - Heap Fetches:
    - Heap Fetches: 0 は、テーブルデータへのアクセスが一切行われなかったことを意味します。
    - これにより、I/O コストが大幅に削減されています。
  - 実行時間:
    - actual time=0.010..0.030 により、非常に短時間でクエリが実行されたことが分かります。

### Seq Scan (シーケンシャルスキャン)

- **概要**:
  - テーブル全体を一行ずつスキャンし、指定された条件に一致する行を探します。
  - インデックスを使用しないため、すべての行を確認します。
- **特徴**:
  - テーブルサイズが小さい場合や、条件がインデックスを活用できない場合に適しています。
  - データ量が増えると、クエリの実行時間が長くなります。
- **最適化**:
  - インデックスを作成して、フィルタ条件に基づいてインデックススキャンを使用するよう調整。
  - フィルタリングの条件を最適化してスキャン範囲を狭める。

```sql
EXPLAIN ANALYZE
SELECT * FROM users WHERE email LIKE '%example.com';
```

```sql:結果例
Seq Scan on users  (cost=0.00..100.00 rows=100 width=128) (actual time=0.020..0.080 rows=50 loops=1)
   Filter: (email ~~ '%example.com%'::text)
   Rows Removed by Filter: 950
```

- 結果解説:
  - Seq Scan on users:
    - users テーブル全体をシーケンシャルスキャンしています。
    - インデックスが利用されていないため、すべての行をチェックしています。
  - Filter:
    - フィルタ条件 email ~~ '%example.com%'::text（LIKE句）が適用されています。
  - Rows Removed by Filter:
    - 条件に一致しなかった行（950行）がフィルタによって除外されました。
  - 実行時間:
    - 実行時間は actual time=0.020..0.080 と短いですが、これはテーブルサイズが小さいためです。

### CTE Scan (共通テーブル式スキャン)

- **概要**:
  - CTE（Common Table Expression）の結果をスキャンし、必要な行を取得します。
  - CTEは一時的なビューのように扱われ、クエリの可読性と再利用性が向上します。
- **特徴**:
  - クエリの構造を整理し、理解しやすくなります。
  - ただし、CTEが複数回利用される場合は、毎回再評価されるためパフォーマンスに影響を及ぼす可能性があります。
- **最適化**:
  - 頻繁に使用される場合は、サブクエリや一時テーブルを検討。
  - CTEの内容が一度だけ評価されるよう、クエリの構造を見直す。

```sql
EXPLAIN ANALYZE
WITH recent_tasks AS (
    SELECT * FROM tasks WHERE status = 'completed'
)
SELECT * FROM recent_tasks;
```

```sql:結果例
CTE Scan on recent_tasks  (cost=0.00..50.00 rows=10 width=64) (actual time=0.020..0.050 rows=10 loops=1)
   -> Seq Scan on tasks  (cost=0.00..50.00 rows=10 width=64) (actual time=0.010..0.030 rows=10 loops=1)
```

- 結果解説
  - CTE Scan on recent_tasks:
    - recent_tasks というCTEをスキャンして結果を取得しています。
    - コストは cost=0.00..50.00、実行時間は actual time=0.020..0.050 です。
  - Seq Scan on tasks:
    - CTEの元になった tasks テーブルをシーケンシャルスキャンしています。
    - 条件として status = 'completed' が適用されています。
    - 10行が条件に一致し、スキャンされました。

### Materialize

- **概要**:
  - サブクエリやCTEの結果を一時的にメモリに保存して再利用。
  - サブクエリが何度も利用される場合にキャッシュされることで計算コストを削減します。
- **特徴**:
  - 結果がメモリに保持され、繰り返し計算のコストを削減。
  - 再利用の必要がない場合には、オーバーヘッドになる可能性がある。
- **使用場面**:
  - 同じサブクエリが複数回利用される場合に有効。
  - CTEやサブクエリの結果を再計算せずに利用したい場合。
- **最適化**:
  - 再利用が不要な場合はMaterializeを避ける設計を行う。
  - サブクエリを簡略化し、計算負荷を低減する。

```sql
EXPLAIN ANALYZE
SELECT
    p.id,
    p.name,
    (
        SELECT COUNT(*) FROM tasks t WHERE t.project_id = p.id
    ) AS task_count,
    (
        SELECT COUNT(*) FROM tasks t WHERE t.project_id = p.id AND t.status = 'completed'
    ) AS completed_task_count
FROM projects p;

```

```sql:結果例
Seq Scan on projects p  (cost=0.00..100.00 rows=100 width=64) (actual time=0.050..5.000 rows=100 loops=1)
   SubPlan 1
     -> Aggregate  (cost=5.00..5.01 rows=1 width=8) (actual time=0.020..0.030 rows=1 loops=100)
           -> Seq Scan on tasks t  (cost=0.00..5.00 rows=1 width=0) (actual time=0.010..0.020 rows=10 loops=100)
                 Filter: (t.project_id = p.id)
   SubPlan 2
     -> Aggregate  (cost=5.00..5.01 rows=1 width=8) (actual time=0.020..0.030 rows=1 loops=100)
           -> Seq Scan on tasks t  (cost=0.00..5.00 rows=1 width=0) (actual time=0.010..0.020 rows=10 loops=100)
                 Filter: (t.project_id = p.id AND t.status = 'completed')
Planning Time: 0.250 ms
Execution Time: 7.000 ms
```

- 結果解説
  - Seq Scan on projects p:
    - projects テーブルをシーケンシャルスキャンしています。
    - 計100行のプロジェクトがスキャンされ、対応する tasks の集計を実行。
  - SubPlan 1:
    - tasks テーブルをスキャンして、プロジェクトに関連付けられたタスクの総数を集計。
    - 計算は100回繰り返され、条件は t.project_id = p.id。
  - SubPlan 2:
    - tasks テーブルをスキャンして、プロジェクトごとにステータスが 'completed' のタスクを集計。
    - 条件は t.project_id = p.id AND t.status = 'completed'。
  - Execution Time:
    - クエリ全体の実行時間は約7ミリ秒。

### Parallel Seq Scan

- **概要**:
  - テーブル全体をスキャンする際に、複数のプロセスで作業を分担して並列化。
  - 大規模なデータセットに対して、複数のCPUを活用して効率化。
- **特徴**:
  - データ量が多い場合や、インデックスが利用できない状況で効果を発揮。
  - 並列化によってスキャン時間を短縮。
- **使用場面**:
  - テーブルサイズが非常に大きく、全行をスキャンする必要がある場合。
  - CPUリソースに余裕がある環境。
- **最適化**:
  - 並列度（parallel workers）の設定をチューニング。
  - 過剰な並列化を避け、システムの負荷を適切に管理。

```sql
EXPLAIN ANALYZE
SELECT * FROM tasks;
```

```sql:結果例
Parallel Seq Scan on tasks  (cost=0.00..100.00 rows=10000 width=64) (actual time=0.010..0.050 rows=10000 loops=4)
   Filter: (status = 'completed')
   Workers Planned: 4
   Workers Launched: 4
   Actual Rows: 2500 per worker
```

- 結果解説
  - Parallel Seq Scan on tasks:
    - tasks テーブル全体をスキャンする操作が、並列で実行されています。
  - Workers Planned/Launched:
    - 計画されたワーカー数と、実際に起動されたワーカー数が表示されています。
    - 例では、4つのワーカープロセスが並列スキャンに参加しています。
  - Filter:
    - status = 'completed' の条件を適用。
  - Actual Rows:
    - 各ワーカーが処理した行数（例: 2500行）を示しています。
    - 合計で10000行がスキャンされています。
  - Execution Time:
    - 並列処理により、実行時間が大幅に短縮されています。

### Sort

- **概要**:
  - クエリ結果を並べ替える操作。
  - ORDER BY句を含むクエリで実行される。
- **特徴**:
  - ソート列にインデックスがない場合、クエリのコストが高くなる。
  - データ量が増えるほどパフォーマンスへの影響が大きくなる。
- **最適化**:
  - 並べ替え対象列にインデックスを作成して効率化する。

```sql
EXPLAIN ANALYZE
SELECT * FROM tasks ORDER BY created_at DESC;
```

```sql:結果例
Sort  (cost=50.00..60.00 rows=100 width=64) (actual time=0.020..0.050 rows=100 loops=1)
   Sort Key: created_at DESC
   Sort Method: quicksort Memory: 100kB
   -> Seq Scan on tasks  (cost=0.00..50.00 rows=100 width=64) (actual time=0.010..0.030 rows=100 loops=1)
```

- 実行計画の解説:
Sort: クエリの並べ替え部分を担当。created_at DESCの条件に従って並べ替え。
  - Sort Method: クイックソート（quicksort）を使用し、メモリ効率は良好。
  - Seq Scan: 並べ替え前にテーブル全体をスキャンしてデータを取得。

### Aggregate

- **概要**:
  - 集約関数（COUNT, SUM, AVG など）を使用する際に実行される操作。
  - データセット全体を計算対象にするため、行数に応じて処理コストが変動する。
- **特徴**:
  - 集約する対象が大きいほど処理コストが高くなる。
  - フィルタリングが適切であれば、不要な行を除外して効率化可能。
- **最適化**:
  - 必要最小限の行に絞り込んだ後で集約処理を実行する。

```sql
EXPLAIN ANALYZE
SELECT COUNT(*) FROM tasks WHERE status = 'completed';
```

```sql:結果例
Aggregate  (cost=50.00..50.10 rows=1 width=8) (actual time=0.020..0.020 rows=1 loops=1)
   -> Seq Scan on tasks  (cost=0.00..50.00 rows=10 width=0) (actual time=0.010..0.010 rows=10 loops=1)
```

- 実行計画の解説:
  - Aggregate: 集約操作を担当。ここでは COUNT(*) による行数の計算。
  - Seq Scan: tasks テーブルをシーケンシャルスキャンし、条件 status = 'completed' を適用。
  - 最適化のヒント:
    - 条件にインデックスを設定するとスキャンを効率化可能。
    - 使用頻度が高い場合、集約済みデータをキャッシュする方法も検討。
