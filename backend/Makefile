.PHONY: run migrate seed setup

# アプリケーションを実行
run:
	go run main.go

# マイグレーションを実行
migrate:
	go run migration/main.go

# シードデータを挿入
seed:
	go run migration/seed/main.go

# セットアップ（マイグレーション、シード実行後にアプリケーションを起動）
setup: migrate seed run

