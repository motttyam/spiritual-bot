# Components — MIXI2 占いBot

## Component 1: FortuneGenerator
**パッケージ**: `internal/fortune/`

### 責務
- 辞書データ（JSON）をロードする
- ランダムに星座・状態・特徴・結果・儀式を選択する
- 2行フォーマットの投稿本文を合成する
- 149文字制限に収まるよう本文をトリミングする

### 管理ファイル
| ファイル | 役割 |
|---------|------|
| `generator.go` | FortuneGenerator 本体・Generate() メソッド |
| `dictionary.go` | Dictionary 構造体・JSON ロード処理 |
| `data/zodiac.json` | 12星座リスト |
| `data/states.json` | 状態ワードリスト（例: あごひげをたくわえていて）|
| `data/traits.json` | 特徴ワードリスト（例: 今歯をみがいている）|
| `data/verdicts.json` | 結果ワードリスト（例: 運勢が一回だけ成仏します）|
| `data/rituals.json` | 儀式ワードリスト（例: コップの水に...）|

---

## Component 2: MIXI2Client
**パッケージ**: `internal/mixi2/`

### 責務
- 環境変数から認証設定を受け取る
- mixi2-application-sdk-go を使い OAuth 2.0 認証を行う
- `CreatePost` RPC でテキスト投稿を実行する
- API エラー発生時にエラーを返す

### 管理ファイル
| ファイル | 役割 |
|---------|------|
| `client.go` | MIXI2Client 構造体・NewClient()・Post() |

---

## Component 3: Main（エントリポイント）
**パッケージ**: `cmd/bot/`

### 責務
- 環境変数から設定を読み込む
- FortuneGenerator・MIXI2Client を初期化する
- 占いコンテンツを生成して MIXI2 へ投稿する
- 成功時 `os.Exit(0)`・失敗時 `os.Exit(1)` で終了する
- 生成した投稿本文をログ出力する

### 管理ファイル
| ファイル | 役割 |
|---------|------|
| `main.go` | オーケストレーション・エラーハンドリング |
