---
name: code-review
description: >-
  ユーザーが「コードレビューして」「品質チェック」「改善提案」「レビューお願い」等と要求した時に発動。
  指定されたファイルやPRの変更内容をレビューし、以下を提供する：
  - バグの可能性
  - パフォーマンス問題
  - Goのイディオム違反
  - セキュリティリスク
  - 改善提案
  結果はマークダウン形式で出力する。
allowed-tools: Read, Grep, Glob, Bash
---

### 手順
1. **レビュー対象の特定**:
   - ファイル指定がある場合: そのファイルを読み込む
   - PR/ブランチ指定がある場合: `git diff` で変更を取得
   - 指定がない場合: 最近変更されたファイルを確認
2. **静的解析の実行**:
   - `go vet ./...` でGoの標準チェック
   - `go fmt -l .` でフォーマット違反を検出
3. **コードレビューの実施**: 以下の観点でチェック
   - **正確性**: ロジックエラー、境界値バグ、nil処理漏れ
   - **パフォーマンス**: 不要なアロケーション、N+1問題、非効率なループ
   - **可読性**: 変数名、関数の責任範囲、コメントの適切性
   - **Goイディオム**: エラーハンドリング、defer使用、interface設計
   - **セキュリティ**: 入力検証、SQL injection、XSS等
   - **テスト**: テストカバレッジ、エッジケース
4. **レポート生成**: マークダウン形式で以下を出力
   - 重要度別（Critical/High/Medium/Low）
   - ファイル・行番号を明記
   - 具体的な改善コード例

### 出力フォーマット例
```markdown
# Code Review Report

## Summary
Reviewed: `pkg/calc/sum.go`
Issues found: 3 (1 High, 2 Medium)

## Issues

### 🔴 High: Off-by-one error in loop
**File**: `pkg/calc/sum.go:6`
**Issue**: Loop condition `i < len(nums)-1` excludes the last element
**Impact**: Incorrect calculation results
**Suggestion**:
\`\`\`go
for i := 0; i < len(nums); i++ {
    total += nums[i]
}
\`\`\`

### 🟡 Medium: Missing nil check
**File**: `pkg/calc/sum.go:3`
**Issue**: Function doesn't handle nil slice
**Suggestion**: Add early return for nil input

## Positive Findings
- ✅ Clear function naming
- ✅ Simple, focused implementation
```

### レビュー観点の詳細

#### Goイディオムチェック
- エラーは戻り値で返す（panicは避ける）
- `defer` の適切な使用
- 不要なelse節の削除
- Interface は小さく保つ
- receiver名は1-2文字の略称

#### パフォーマンスチェック
- 不要な `append` のアロケーション
- `strings.Builder` vs `+= string`
- `sync.Pool` の活用機会
- Goroutine leak の可能性

#### セキュリティチェック
- ユーザー入力の検証
- パス traversal の可能性
- 機密情報のログ出力

### ベストプラクティス
- 建設的なトーンで指摘（批判ではなく改善提案）
- 問題だけでなく、良い点も指摘
- 具体的なコード例を示す
- 重要度を明確にして優先順位をつける
