# Domain Entities — fortune-bot

## Dictionary（辞書）

占い本文を生成するための語彙セット。各フィールドは JSON ファイルから `//go:embed` でロードする。

```go
type Dictionary struct {
    Zodiac   []string `json:"zodiac"`   // 12星座
    States   []string `json:"states"`   // 状態（今〜している）
    Traits   []string `json:"traits"`   // 特徴（外見・習慣）
    Verdicts []string `json:"verdicts"` // 結果（断言）
    Rituals  []string `json:"rituals"`  // 儀式（ラッキー行動）
}
```

### 各フィールドの意味と制約

| フィールド | 意味 | 例 | 最大文字数 |
|-----------|------|-----|-----------|
| `Zodiac` | 星座名 | 乙女座、牡羊座 | 4文字 |
| `States` | 今この瞬間の状態 | 今歯をみがいている | 20文字 |
| `Traits` | 外見・習慣の特徴 | あごひげをたくわえていて | 20文字 |
| `Verdicts` | 運勢の断言（スピリチュアル過激） | 運勢が一回だけ成仏します | 25文字 |
| `Rituals` | ラッキー行動（具体・回数つき・無害） | コップの水に「ありがとう」を3回言ってから一口だけ飲む | 40文字 |

### 初期エントリ数

各辞書 **20〜30件** を初期搭載する（5分おき投稿でのローテーション耐性を確保するため）。

---

## FortuneGenerator

```go
type FortuneGenerator struct {
    dict Dictionary
    rng  *rand.Rand // math/rand/v2 — 明示的シードで初期化
}

// rng が nil の場合は time.Now().UnixNano() シードで自動初期化（本番用）
// テスト時は固定シードの rng を渡して再現性を確保する
func NewGenerator(rng *rand.Rand) (*FortuneGenerator, error)
```

---

## FortuneContent（生成結果の中間表現）

テンプレート合成後・トリミング前の状態を表す内部型。

```go
type FortuneContent struct {
    Zodiac  string
    State   string
    Trait   string
    Verdict string
    Ritual  string
}
```

`Generate()` 内部でのみ使用し、外部には最終的な `string` のみを返す。

---

## 文字数カウントの単位

MIXI2 の 149 文字制限は **rune（Unicode コードポイント）単位**で計算する。

```go
// 正しい計算
length := len([]rune(text)) // ✅ 日本語1文字 = 1

// 誤った計算
length := len(text) // ❌ 日本語1文字 = 3バイト
```
