# Component Methods — MIXI2 占いBot

## internal/fortune/

### Dictionary 構造体（dictionary.go）

```go
type Dictionary struct {
    Zodiac   []string `json:"zodiac"`
    States   []string `json:"states"`
    Traits   []string `json:"traits"`
    Verdicts []string `json:"verdicts"`
    Rituals  []string `json:"rituals"`
}

// JSON ファイル群を //go:embed でロードして Dictionary を返す
func LoadDictionary() (Dictionary, error)
```

### FortuneGenerator 構造体（generator.go）

```go
type FortuneGenerator struct {
    dict Dictionary
}

// FortuneGenerator を初期化する
func NewGenerator() (*FortuneGenerator, error)

// 占い本文を生成して返す（149文字以内に収める）
// 生成ロジックと投稿ロジックを分離するため、文字列を返すのみ
func (g *FortuneGenerator) Generate() (string, error)

// 本文が maxLen 文字を超えていた場合にトリミングする（内部メソッド）
// 2行目（儀式）を優先的に残し、1行目の修飾語を削る
func (g *FortuneGenerator) truncate(line1, line2 string, maxLen int) string
```

---

## internal/mixi2/

### Config 構造体（client.go）

```go
type Config struct {
    ClientID     string // MIXI2_CLIENT_ID
    ClientSecret string // MIXI2_CLIENT_SECRET
    TokenURL     string // MIXI2_TOKEN_URL
    APIAddress   string // MIXI2_API_ADDRESS
}
```

### MIXI2Client 構造体（client.go）

```go
type Client struct {
    // mixi2-application-sdk-go の内部クライアントを保持
}

// 環境変数ベースの Config から Client を初期化する
func NewClient(cfg Config) (*Client, error)

// テキストを MIXI2 へ投稿する（CreatePost RPC を呼び出す）
func (c *Client) Post(ctx context.Context, text string) error
```

---

## cmd/bot/

### main.go

```go
// エントリポイント
// 1. 環境変数から Config を構築
// 2. FortuneGenerator.Generate() で本文を生成
// 3. ログに本文を出力
// 4. MIXI2Client.Post() で投稿
// 5. エラー時は log.Fatal（内部で os.Exit(1)）
func main()
```

---

## 設計上の注意

- `Generate()` は文字列を返すのみ。MIXI2 への投稿は知らない（関心の分離）
- `Post()` はテキストを受け取るのみ。生成ロジックは知らない（関心の分離）
- 将来 LLM を挟む場合は `Generate()` の戻り値を LLM に通してから `Post()` する形で拡張可能
