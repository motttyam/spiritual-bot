# Code Generation Plan — fortune-bot

## ユニットコンテキスト

- **プロジェクトタイプ**: Greenfield / Go CLI
- **ワークスペースルート**: `/Users/kazuki.tsukamoto/Desktop/mixi2-development`
- **Go バージョン**: 1.24.6 以上（SDK 要件）
- **モジュール名**: `github.com/kazuki-tsukamoto/mixi2-fortune-bot`（要変更時は go.mod を修正）

---

## 生成ステップ

### Step 1: プロジェクト初期化
- [ ] `go.mod` を作成（module 宣言、Go バージョン指定）

### Step 2: 辞書 JSON ファイルの生成（各20〜30件）
- [ ] `internal/fortune/data/zodiac.json`（12星座）
- [ ] `internal/fortune/data/states.json`（状態ワードリスト）
- [ ] `internal/fortune/data/traits.json`（特徴ワードリスト）
- [ ] `internal/fortune/data/verdicts.json`（結果ワードリスト）
- [ ] `internal/fortune/data/rituals.json`（儀式ワードリスト）

### Step 3: 辞書ローダーの生成
- [ ] `internal/fortune/dictionary.go`
  - `Dictionary` 構造体
  - `//go:embed data/*.json`
  - `LoadDictionary()` — JSON パース + rune 長バリデーション + 最低件数チェック + 最悪ケースログ出力

### Step 4: 占いジェネレーターの生成
- [ ] `internal/fortune/generator.go`
  - `FortuneContent` 構造体
  - `FortuneGenerator` 構造体（`dict` + `rng *rand.Rand`）
  - `NewGenerator(rng *rand.Rand) (*FortuneGenerator, error)`
  - `Generate() (string, error)`
  - `truncate(content FortuneContent, maxLen int) string`（line2 構造優先）

### Step 5: 占いジェネレーターのユニットテスト
- [ ] `internal/fortune/generator_test.go`
  - 固定シードで再現可能なテスト
  - フル形が 149 rune 以内に収まることの確認
  - truncate の各 Step（0〜3）の動作確認
  - rune カウントの正確性確認

### Step 6: MIXI2 クライアントの生成
- [ ] `internal/mixi2/client.go`
  - `Config` 構造体（`MIXI2_CLIENT_ID` / `MIXI2_CLIENT_SECRET` / `MIXI2_TOKEN_URL` / `MIXI2_API_ADDRESS`）
  - `Client` 構造体
  - `NewClient(cfg Config) (*Client, error)` — `auth.NewAuthenticator` 初期化
  - `Post(ctx context.Context, text string) error` — `CreatePost` RPC 呼び出し

### Step 7: エントリポイントの生成
- [ ] `cmd/bot/main.go`
  - 環境変数から `Config` を構築
  - `fortune.NewGenerator(nil)` で本番用 rng 初期化
  - `Generate()` で本文生成 → `log.Printf("posting: %s", text)` でログ出力
  - `mixi2.NewClient(cfg)` → `client.Post(ctx, text)` で投稿
  - エラー時は `log.Fatal`（内部で `os.Exit(1)`）

### Step 8: GitHub Actions ワークフローの生成
- [ ] `.github/workflows/fortune.yml`
  - `on.schedule: cron: '*/5 * * * *'`
  - `concurrency: group: fortune-bot, cancel-in-progress: true`
  - `go-version: '1.24.6'`
  - GitHub Secrets から環境変数を注入
  - `go run cmd/bot/main.go` で実行

### Step 9: README の生成
- [ ] `README.md`
  - 概要（占いBot の説明）
  - 必要な GitHub Secrets の一覧と取得方法
  - ローカル実行方法（`.env` ファイル使用）
  - ディレクトリ構成

---

## 生成ファイル一覧

```
/Users/kazuki.tsukamoto/Desktop/mixi2-development/
├── go.mod
├── cmd/
│   └── bot/
│       └── main.go
├── internal/
│   ├── fortune/
│   │   ├── dictionary.go
│   │   ├── generator.go
│   │   ├── generator_test.go
│   │   └── data/
│   │       ├── zodiac.json
│   │       ├── states.json
│   │       ├── traits.json
│   │       ├── verdicts.json
│   │       └── rituals.json
│   └── mixi2/
│       └── client.go
├── .github/
│   └── workflows/
│       └── fortune.yml
└── README.md
```

---

## 依存ライブラリ（go.mod に追加）

| ライブラリ | 用途 |
|-----------|------|
| `github.com/mixigroup/mixi2-application-sdk-go` | MIXI2 API クライアント・OAuth 2.0 認証 |
