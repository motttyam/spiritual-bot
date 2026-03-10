# Business Logic Model — fortune-bot

## 1. 辞書ロードフロー

```
起動
  |
  └─ LoadDictionary()
        |
        ├─ //go:embed data/zodiac.json   → []string
        ├─ //go:embed data/states.json   → []string
        ├─ //go:embed data/traits.json   → []string
        ├─ //go:embed data/verdicts.json → []string
        └─ //go:embed data/rituals.json  → []string
              |
              └─ Dictionary{} を返す
```

---

## 2. 占い本文生成フロー

```
Generate()
  |
  ├─ 1. 各辞書からランダムに1つ選択
  |       zodiac  = Zodiac[rand]
  |       state   = States[rand]
  |       trait   = Traits[rand]
  |       verdict = Verdicts[rand]
  |       ritual  = Rituals[rand]
  |
  ├─ 2. FortuneContent{} に格納
  |
  ├─ 3. テンプレートに当てはめて2行生成
  |       line1 = "{trait}、{state}{zodiac}のあなた、今日、{verdict}。"
  |       line2 = "ラッキー行動：{ritual}"
  |
  ├─ 4. truncate(line1, line2, 149) でトリミング
  |
  └─ 5. "{trimmed_line1}\n{line2}" を返す
```

---

## 3. トリミングアルゴリズム（truncate）

**前提**: 辞書設計により通常は 149 rune 以内に収まる。トリミングは例外的ケースのみ。

```
truncate(content FortuneContent, maxLen=149) string
  |
  ├─ Step 0: フル形が <= 149 → そのまま返す
  |           "{trait}、{state}{zodiac}のあなた、今日、{verdict}。\nラッキー行動：{ritual}"
  |
  ├─ Step 1: trait を削除して試みる
  |           "{state}{zodiac}のあなた、今日、{verdict}。\nラッキー行動：{ritual}"
  |           <= 149 → 返す
  |
  ├─ Step 2: state も削除して試みる（最小形）
  |           "{zodiac}のあなた、今日、{verdict}。\nラッキー行動：{ritual}"
  |           <= 149 → 返す
  |
  └─ Step 3: 最終手段（辞書設計・検証が正しければ到達しないはず）
              line1_minimal = "{zodiac}のあなた、今日、{verdict}。"
              remaining = maxLen - len([]rune(line1_minimal)) - 1（\n分）- 7（"ラッキー行動："分）- 1（"…"分）
              ritual_short = ritual[:remaining] + "…"
              → "{line1_minimal}\nラッキー行動：{ritual_short}"
              ※ line2（儀式）を優先して残す。改行・"ラッキー行動："は必ず維持する
```

### トリミング優先度（削る順）

1. **trait（特徴）** を先に削る — state の方が「今この瞬間感」が出るため残す
2. **state（状態）** を次に削る
3. **ritual を末尾 rune 単位で短縮 + `…`** — line2 の構造（`ラッキー行動：`）は必ず残す

---

## 4. ランダム選択

```go
// math/rand/v2 を使用（Go 1.22+）
// シードを明示的に与える（math/rand v1 との混同・固定シード事故を防ぐ）
rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
index := rng.IntN(len(slice))
item := slice[index]
```

- `FortuneGenerator` 生成時に `rng` を初期化し、フィールドとして保持する
- `time.Now().UnixNano()` をシードにすることで起動ごとに必ず異なる結果が保証される
- テスト時は固定シード（例: `rand.NewPCG(42, 0)`）を渡せるよう `NewGenerator(rng)` の形を取る

---

## 5. 生成例

| パーツ | 選択値 |
|--------|--------|
| zodiac | 乙女座 |
| trait | あごひげをたくわえていて |
| state | 今歯をみがいている |
| verdict | 運勢が一回だけ成仏します |
| ritual | コップの水に「ありがとう」を3回言ってから一口だけ飲む |

**生成結果（81 rune）:**
```
あごひげをたくわえていて、今歯をみがいている乙女座のあなた、今日、運勢が一回だけ成仏します。
ラッキー行動：コップの水に「ありがとう」を3回言ってから一口だけ飲む
```
