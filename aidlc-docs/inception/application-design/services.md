# Services — MIXI2 占いBot

## FortunePostingService（main.go が担当）

このプロジェクトは単純な単発実行 CLI のため、独立したサービス層は設けず、`cmd/bot/main.go` がオーケストレーター兼サービス層を兼務する。

### 処理フロー

```
main()
  |
  ├─ 1. 環境変数から Config を読み込む
  |       MIXI2_CLIENT_ID, MIXI2_CLIENT_SECRET,
  |       MIXI2_TOKEN_URL, MIXI2_API_ADDRESS
  |
  ├─ 2. fortune.NewGenerator() で FortuneGenerator を初期化
  |
  ├─ 3. generator.Generate() で占い本文を生成
  |       └─ 内部: LoadDictionary() → ランダム選択 → テンプレート合成 → truncate()
  |
  ├─ 4. log.Printf("生成本文: %s", text) でログ出力
  |
  ├─ 5. mixi2.NewClient(cfg) で MIXI2Client を初期化
  |
  ├─ 6. client.Post(ctx, text) で MIXI2 へ投稿
  |
  └─ 7. エラー発生時は log.Fatal → os.Exit(1)
         正常終了時は os.Exit(0)（暗黙）
```

### 責務の境界

| 責務 | 担当 |
|------|------|
| 設定の読み込み | main.go |
| 占いコンテンツ生成 | internal/fortune/ |
| MIXI2 API 通信 | internal/mixi2/ |
| オーケストレーション | main.go |
| エラーハンドリング・終了 | main.go |
