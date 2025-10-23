# CLAUDE.md

このファイルは、このリポジトリで作業する際に Claude Code (claude.ai/code) にガイダンスを提供します。

## プロジェクト概要

このリポジトリは Claude Code Skills と自動化された PR ワークフローのデモンストレーション用です。シンプルな Go 計算機パッケージ (`pkg/calc`) を含んでおり、自動化された Issue 解決と PR 作成ワークフローをデモンストレーションするために意図的なバグが含まれています。

## 開発コマンド

### テスト実行
```bash
# すべてのテストを実行
go test ./...

# 特定のパッケージのテストを実行
go test ./pkg/calc

# 特定のテストを実行
go test ./pkg/calc -run TestSumBasic

# 詳細出力でテストを実行
go test -v ./...
```

### ビルド
```bash
# すべてのパッケージをビルド
go build ./...

# 特定のパッケージをビルド
go build ./pkg/calc
```

## アーキテクチャ

### パッケージ構成
- `pkg/calc/` - デモンストレーション用の意図的なバグを含むコア計算ユーティリティ
  - `sum.go` - `Sum()` 関数を含む（現在既知のバグあり：最後の要素を除外）
  - `sum_test.go` - Sum 関数のテストケース

### Claude Code Skills

このリポジトリには `.claude/skills/` にカスタムスキルが含まれています。
これらのスキルは、ユーザーの要求内容に基づいて Claude Code が自律的に選択・実行します。

利用可能なスキル:
- **issue-fix**: GitHub Issue の自動解決と PR 作成
- **test-generator**: テストケースの自動生成
- **pr-description**: PR 本文の自動生成
- **code-review**: コード品質チェックとレビュー
- **refactor**: リファクタリング
- **doc-generator**: GoDoc コメントとドキュメント生成
- **benchmark-analyzer**: パフォーマンス分析

詳細は各スキルの SKILL.md を参照してください。
**重要**: タスクを実行する前に、必ず適切なスキルが利用可能かどうかを判断してください。

### 自動化ワークフロー

#### Issue Auto-Fix (.github/workflows/claude-issue-auto-fix.yml)

Issue に `bug` ラベルが付けられると自動的にトリガーされます：
- リポジトリをチェックアウト
- Claude Code を実行（プロンプト: "Issueを解決してください。Issue: #番号"）

これは Claude Code を使用した完全自動化された Issue-to-PR パイプラインのデモンストレーションです。

ワークフロー設定:
- 最大ターン数: 30
- パーミッションモード: `bypassPermissions`
- 必要なシークレット: `CLAUDE_CODE_OAUTH_TOKEN`
- 使用アクション: `anthropics/claude-code-action@fc4013af386ecc44b387ef2931c8d5f7c268b44e`（skills サポート版）

#### PR Comment Trigger (.github/workflows/claude-pr-comment.yml)

PR コメントで `@claude` とメンションすると Claude Code を実行できます：

使用例:
```
@claude このPRをレビューしてください
@claude テストを追加してください
@claude パフォーマンスを改善してください
```

セキュリティ対策:
- **権限チェック**: リポジトリの write または admin 権限を持つユーザーのみ実行可能
- **PR のみ**: Issue コメントでは実行されません（PR コメントのみ）
- **プロンプト制限**: コメントから抽出されるプロンプトは最大500文字に制限
- **改行削除**: プロンプトインジェクション対策として改行は削除されます
- **スコープ制限**: プロンプト内で `pkg/calc` パッケージのみを対象とするよう指示

ワークフロー設定:
- 最大ターン数: 20
- パーミッションモード: `bypassPermissions`
- 必要なシークレット: `CLAUDE_CODE_OAUTH_TOKEN` (または `/install-github-app` で自動設定)
- 必要な permissions: `contents: write`, `pull-requests: write`, `issues: write`, `id-token: write`
- 使用アクション: `anthropics/claude-code-action@fc4013af386ecc44b387ef2931c8d5f7c268b44e` (skills サポート版)
- トリガー: `issue_comment` (PR コメント)、`pull_request_review_comment` (PR レビューコメント)

セットアップ方法:
```bash
# ターミナルで実行（GitHub App と必要なシークレットを自動設定）
/install-github-app
```

注意: `.claude/workflows/claude-pr-fix.yml` にも同様のワークフロー定義がありますが、これは異なる実装例です。

## 既知の問題

`pkg/calc/sum.go` の `Sum()` 関数には、計算から最後の要素を除外する意図的なバグがあります（6行目を参照）。
