# Component Dependency — MIXI2 占いBot

## 依存関係マトリクス

| コンポーネント | 依存先 | 依存の種類 |
|--------------|--------|----------|
| `cmd/bot/main.go` | `internal/fortune/` | 直接呼び出し |
| `cmd/bot/main.go` | `internal/mixi2/` | 直接呼び出し |
| `internal/fortune/` | `data/*.json`（embed） | ファイル埋め込み |
| `internal/fortune/` | 標準ライブラリのみ | — |
| `internal/mixi2/` | `mixi2-application-sdk-go` | 外部ライブラリ |
| `internal/mixi2/` | 標準ライブラリのみ | — |

## データフロー

```
[環境変数]
    |
    v
[cmd/bot/main.go]
    |                          |
    v                          v
[internal/fortune/]     [internal/mixi2/]
    |                          |
    v                          v
[data/*.json]      [mixi2-application-sdk-go]
  (//go:embed)               |
                             v
                       [MIXI2 API]
```

## 依存の方向性

- `main.go` → `fortune` / `mixi2`（一方向）
- `fortune` と `mixi2` は互いに依存しない（完全分離）
- 将来 LLM を追加する場合は `main.go` で `Generate()` 結果を LLM に渡す形にする

## 外部依存

| ライブラリ | 用途 |
|-----------|------|
| `github.com/mixigroup/mixi2-application-sdk-go` | MIXI2 API クライアント・OAuth 2.0 認証 |
| Go 標準ライブラリ（`encoding/json`, `embed`, `math/rand`, `log`, `os`） | JSON読み込み・ランダム選択・ログ・終了処理 |
