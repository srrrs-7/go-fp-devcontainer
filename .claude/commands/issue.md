# Create GitHub Issue

GitHubにIssueを作成します。

## 引数

$ARGUMENTS - Issueの概要（タイトルや内容の説明）

## 実行内容

ユーザーの入力「$ARGUMENTS」を元に、以下の手順でIssueを作成してください：

1. 入力内容を分析し、適切なIssueタイトルと説明を生成
2. 以下の形式でIssueを作成：

```bash
gh issue create --title "Issueタイトル" --body "$(cat <<'EOF'
## 概要
問題や機能の説明

## 詳細
- 詳細1
- 詳細2

## 期待する動作
期待する結果の説明

## 関連情報
- 関連するファイルやコンポーネント
EOF
)"
```

3. 作成したIssueのURLを表示

## 例

- `/issue ログイン機能にバグがある` → ログイン関連のバグIssueを作成
- `/issue ダークモードを追加したい` → 機能リクエストIssueを作成
- `/issue テストカバレッジを改善` → 改善提案Issueを作成
