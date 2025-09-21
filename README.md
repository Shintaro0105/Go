# Go + PostgreSQL API

このプロジェクトは、**Go** 言語と **PostgreSQL** を使ったシンプルな API サンプルです。  
DevContainer (`.devcontainer/`) を使って、Docker 上で Go + PostgreSQL の開発環境を構築しています。

---

## 🚀 開発環境の起動方法

1. DevContainer を VSCode で開く  
2. `docker-compose` が PostgreSQL を起動  
3. Go コンテナ内で API を実行  

---

## 📦 セットアップ

依存をインストール:

```bash
go mod tidy
```

環境変数は .devcontainer/.env に定義されています:

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
POSTGRES_HOSTNAME=db
POSTGRES_PORT=5432
```

---

## ▶️ 実行方法

```bash
go run main.go
```

サーバーが起動したら、次のエンドポイントが利用できます:
- ```GET http://localhost:8080/users```

登録済みユーザーの一覧を JSON で返します。

---

## 🗄️ 初期データ

PostgreSQL 内に以下のテーブルを作成してください:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO users (name) VALUES ('Alice'), ('Bob');
```