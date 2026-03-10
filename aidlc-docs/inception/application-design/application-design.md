# Application Design — MIXI2 占いBot

## 概要

GitHub Actions の cron で5分ごとに起動し、MIXI2 に占い投稿を1件行って終了する単発実行 CLI Bot。

---

## ディレクトリ構成

```
<project-root>/
├── cmd/
│   └── bot/
│       └── main.go                  # エントリポイント・オーケストレーター
├── internal/
│   ├── fortune/
│   │   ├── generator.go             # FortuneGenerator・Generate()・truncate()
│   │   ├── dictionary.go            # Dictionary 構造体・LoadDictionary()
│   │   └── data/
│   │       ├── zodiac.json          # 12星座リスト
│   │       ├── states.json          # 状態ワードリスト
│   │       ├── traits.json          # 特徴ワードリスト
│   │       ├── verdicts.json        # 結果ワードリスト
│   │       └── rituals.json         # 儀式ワードリスト
│   └── mixi2/
│       └── client.go                # MIXI2Client・NewClient()・Post()
├── .github/
│   └── workflows/
│       └── fortune.yml              # cron: */5 * * * *
├── go.mod
├── go.sum
└── README.md
```

---

## コンポーネント一覧

### FortuneGenerator（internal/fortune/）

| 要素 | 内容 |
|------|------|
| 責務 | 辞書読み込み・本文生成・文字数トリミング |
| 入力 | なし（embedされたJSONを使用） |
| 出力 | `string`（投稿本文・149文字以内） |
| 外部依存 | なし（標準ライブラリのみ） |

**主要メソッド:**
```go
func NewGenerator() (*FortuneGenerator, error)
func (g *FortuneGenerator) Generate() (string, error)
func (g *FortuneGenerator) truncate(line1, line2 string, maxLen int) string
```

### MIXI2Client（internal/mixi2/）

| 要素 | 内容 |
|------|------|
| 責務 | OAuth 2.0 認証・CreatePost RPC 実行 |
| 入力 | Config（ClientID, ClientSecret, TokenURL, APIAddress） |
| 出力 | error |
| 外部依存 | mixi2-application-sdk-go |

**主要メソッド:**
```go
func NewClient(cfg Config) (*Client, error)
func (c *Client) Post(ctx context.Context, text string) error
```

### Main（cmd/bot/）

| 要素 | 内容 |
|------|------|
| 責務 | 設定読み込み・コンポーネント初期化・オーケストレーション |
| 成功時 | `os.Exit(0)`（暗黙） |
| 失敗時 | `log.Fatal` → `os.Exit(1)` |

---

## 環境変数

| 変数名 | 説明 |
|--------|------|
| `MIXI2_CLIENT_ID` | OAuth 2.0 クライアント ID |
| `MIXI2_CLIENT_SECRET` | OAuth 2.0 クライアントシークレット |
| `MIXI2_TOKEN_URL` | アクセストークン取得エンドポイント |
| `MIXI2_API_ADDRESS` | MIXI2 API サーバーアドレス |

---

## 投稿本文フォーマット

```
{星座}のあなた、{状態}、{特徴}、{結果（断言）}。
ラッキー行動：{儀式}
```

- 合計 149 文字以内
- 超過時は1行目の修飾語（状態・特徴）を削って調整
- 2行目（儀式）は最優先で残す

---

## 関心の分離

```
fortune.Generate()  ──────→  string
                                │
                   （将来LLMを挟める）
                                │
mixi2.Post(ctx, text)  ←──────
```

`fortune` と `mixi2` は互いを知らない。拡張は `main.go` のオーケストレーション部分だけで完結する。
