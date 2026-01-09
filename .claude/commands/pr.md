# Create Pull Request

GitHubにPull Requestを作成します。

## 手順

1. 現在のブランチの変更内容を確認
2. 適切なPRタイトルと説明を生成
3. `gh pr create`でPRを作成

## 実行内容

以下の手順でPRを作成してください：

1. `git status`で現在の状態を確認
2. `git log main..HEAD --oneline`で含まれるコミットを確認
3. `git diff main...HEAD --stat`で変更ファイルを確認
4. 変更内容を分析し、PRタイトルと説明を作成
5. 以下の形式でPRを作成：

```bash
gh pr create --title "PRタイトル" --body "$(cat <<'EOF'
## Summary
- 変更点1
- 変更点2

## Test plan
- [ ] テスト項目1
- [ ] テスト項目2

🤖 Generated with [Claude Code](https://claude.ai/code)
EOF
)"
```

PRのURLを表示してください。
