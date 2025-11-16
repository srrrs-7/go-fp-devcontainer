# Database Package

このパッケージは、PostgreSQL データベースのスキーマ、マイグレーション、および生成されたクエリコードを管理します。

## ディレクトリ構造

```
db/
├── atlas.hcl           # Atlas マイグレーション設定
├── sqlc.yaml           # sqlc コード生成設定
├── schema/             # データベーススキーマ定義（SQL）
│   └── tasks.sql
├── queries/            # SQLクエリ定義（sqlc用）
│   └── tasks.sql
├── migrations/         # バージョン管理されたマイグレーションファイル
│   └── .gitkeep
└── db/                 # sqlc生成コード
    ├── db.go
    ├── models.go
    ├── querier.go
    └── tasks.sql.go
```

## ツール

- **[Atlas](https://atlasgo.io/)**: データベーススキーママイグレーションツール
- **[sqlc](https://sqlc.dev/)**: SQLからタイプセーフなGoコードを生成

## 環境変数

データベース接続は `.devcontainer/compose.override.yaml` で定義された環境変数から構築されます:

- `DB_HOST`: データベースホスト（例: `db`）
- `DB_PORT`: データベースポート（例: `5432`）
- `DB_DBNAME`: データベース名（例: `mydb`）
- `DB_USERNAME`: データベースユーザー（例: `postgres`）
- `DB_PASSWORD`: データベースパスワード

## マイグレーションワークフロー

### 1. スキーマファイルの編集

`schema/` ディレクトリ内のSQLファイルを編集して、希望するデータベーススキーマを定義します。

```sql
-- schema/tasks.sql
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    -- ...
);
```

### 2. マイグレーションの生成

スキーマファイルと現在のデータベースの差分からマイグレーションを生成:

```bash
make atlas-diff
```

または、空のマイグレーションファイルを作成:

```bash
make atlas-new NAME=add_users_table
```

### 3. マイグレーションの確認

生成されたマイグレーションファイルを `migrations/` ディレクトリで確認します。

### 4. マイグレーションの適用

```bash
# 確認プロンプトあり
make atlas-apply

# 自動承認（注意して使用）
make atlas-apply-auto
```

### 5. マイグレーションステータスの確認

```bash
make atlas-status
```

## Atlas コマンド一覧

| コマンド | 説明 |
|---------|------|
| `make atlas-new NAME=<name>` | 新しいマイグレーションファイルを作成 |
| `make atlas-diff` | スキーマの差分からマイグレーションを生成 |
| `make atlas-apply` | 保留中のマイグレーションを適用 |
| `make atlas-apply-auto` | マイグレーションを自動承認で適用 |
| `make atlas-status` | マイグレーションステータスを表示 |
| `make atlas-hash` | マイグレーションディレクトリを再ハッシュ |
| `make atlas-lint` | マイグレーションファイルをリント |
| `make atlas-validate` | マイグレーションディレクトリの整合性を検証 |
| `make atlas-inspect` | データベーススキーマを検査 |
| `make atlas-schema-apply` | スキーマを直接適用（開発用） |
| `make atlas-clean` | 開発用データベースをクリーン/リセット |
| `make atlas-help` | Atlas コマンドのヘルプを表示 |

## sqlc コマンド

| コマンド | 説明 |
|---------|------|
| `make sqlc-gen` | クエリからGoコードを生成 |
| `make sqlc-compile` | SQLクエリをコンパイル（検証） |
| `make sqlc-verify` | 生成されたコードを検証 |

## 環境の切り替え

デフォルトは `local` 環境ですが、`ATLAS_ENV` 変数で変更可能:

```bash
make atlas-apply ATLAS_ENV=docker
make atlas-apply ATLAS_ENV=ci
```

利用可能な環境:
- `local`: ローカル開発（デフォルト）
- `docker`: Docker Compose環境
- `ci`: CI/CD環境

## 開発データベースについて

Atlas は差分生成や検証のために開発用データベース (`mydb_dev`) を使用します。
このデータベースは Atlas が自由に操作できる必要があります。

## ベストプラクティス

1. **スキーマファーストアプローチ**: `schema/` ディレクトリでスキーマを定義し、`atlas-diff` でマイグレーションを生成
2. **マイグレーションのレビュー**: 適用前に生成されたマイグレーションを必ず確認
3. **バージョン管理**: すべてのマイグレーションファイルをGitにコミット
4. **手動編集の禁止**: 適用後のマイグレーションファイルは編集しない
5. **リント実行**: `make atlas-lint` で問題を早期発見

## トラブルシューティング

### マイグレーションの整合性エラー

```bash
make atlas-hash
```

### データベース接続エラー

環境変数が正しく設定されているか確認:

```bash
echo $DB_HOST $DB_PORT $DB_DBNAME
```

### スキーマの完全なリセット（開発用）

```bash
make atlas-clean
make atlas-apply
```
