# PostgreSQL パフォーマンスサンプル

## Summary

- PostgreSQL + GoLang/Echo を使ったバックエンドWebAPI。
- マイグレーションやシードデータ挿入を自動化。
- バックエンドWebAPIを構築し、パフォーマンスチューニングを試行する。

## Tech Stack

[![GoLang](https://img.shields.io/badge/GoLang-1.22.x-blue)](https://go.dev/)
[![Echo](https://img.shields.io/badge/Echo-4.12.x-blue)](https://echo.labstack.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16.1-blue)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-26.1.x-blue)](https://www.docker.com/)
[![Docker Compose](https://img.shields.io/badge/Docker_Compose-2.22.x-blue)](https://docs.docker.com/compose/)
[![GORM](https://img.shields.io/badge/GORM-v2.1.x-blue)](https://gorm.io/)

## ディレクトリ構成

```
.
├── backend
│   ├── migration
│   │   ├── seeds
│   │   │   └── main.go(シードデータ挿入)
│   │   └── main.go(マイグレーション実行)
│   ├── src
│   │   ├── (GoLang/Echoコード)
│   │   └── ...
│   ├── Dockerfile
│   ├── .env
│   ├── .gitignore
│   └── ...
├── docker-compose.yml
├── .env
├── .gitignore
└── README.md
```

## データベース

- PostgreSQL 16.1

## DBコンテナ起動(Docker Compose)

```bash
docker compose up -d
```

## マイグレーション

```bash
cd backend
make migrate
```

## シードデータ挿入

```bash
cd backend
make seed
```

## バックエンドWebAPI実行

```bash
cd backend
make run
```
