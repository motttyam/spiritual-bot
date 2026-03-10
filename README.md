# spiritual-bot

MIXI2 に 5 分おきで占い投稿を行う Go 製 CLI Bot です。

## 仕様

- 投稿は 2 行フォーマット
- 149 文字制限を rune 単位で保証
- 辞書は `internal/fortune/data/*.json` を `//go:embed` で読み込み
- GitHub Actions `cron: */5 * * * *` で単発実行

## 必要な Secrets

GitHub リポジトリの `Settings > Secrets and variables > Actions` に以下を登録してください。

- `MIXI2_CLIENT_ID`
- `MIXI2_CLIENT_SECRET`
- `MIXI2_TOKEN_URL`
- `MIXI2_API_ADDRESS`

## ローカル実行

1. Go 1.24.6 以上を用意
2. 環境変数を設定
3. 依存を解決
4. 実行

```bash
export MIXI2_CLIENT_ID=...
export MIXI2_CLIENT_SECRET=...
export MIXI2_TOKEN_URL=...
export MIXI2_API_ADDRESS=...

go mod tidy
go run ./cmd/bot/main.go
```

## テスト

```bash
go test ./...
```

## ディレクトリ

```text
.
├── cmd/bot/main.go
├── internal/fortune/
│   ├── dictionary.go
│   ├── generator.go
│   ├── generator_test.go
│   └── data/*.json
├── internal/mixi2/client.go
└── .github/workflows/fortune.yml
```
