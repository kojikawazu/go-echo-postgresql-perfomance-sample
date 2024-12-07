# GoLang(Echo)の構築マニュアル

## 1. プロジェクトの構築

```bash
cd backend
go mod init backend
```

## 2. パッケージのインストール

```bash
go get github.com/labstack/echo/v4
go get github.com/jackc/pgx/v4
```

## 3. サーバーの起動

```bash
go run main.go
```

## 4. パッケージの整理

```bash
go mod tidy
```
