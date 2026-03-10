# Functional Design Plan — fortune-bot

## 実行チェックリスト

- [x] ドメインエンティティの定義（domain-entities.md）
- [x] ビジネスロジックモデルの定義（business-logic-model.md）
- [x] ビジネスルールの定義（business-rules.md）

---

## 設計判断に関する質問

各 `[Answer]:` タグの後に選択肢の記号を入力してください。

---

## Question 1
1行目のテンプレート構造を教えてください。
traits（特徴）と states（状態）は **両方** 必ず入れますか？それとも片方だけでもOK？

A) 両方必ず入れる（例: `{traits}、{states}{zodiac}のあなた、今日、{verdict}。`）
B) どちらか1つだけランダムに選ぶ（例: `{states}{zodiac}のあなた` または `{traits}{zodiac}のあなた`）
C) Other (please describe after [Answer]: tag below)

[Answer]:
両方（traits＋states）必ず入れる。ここが“超具体”の核で、当たり判定っぽさと笑いが出やすい。
（ただし文字数超過時にだけ削る運用でOK）
---

## Question 2
149文字を超えたとき、1行目から削る順番はどちらがよいですか？

A) まず **traits（特徴）** を削る → まだ超過なら **states（状態）** も削る
B) まず **states（状態）** を削る → まだ超過なら **traits（特徴）** も削る
C) Other (please describe after [Answer]: tag below)

[Answer]:
まず traits（特徴）→ 次に states（状態） の順で削る。
理由：statesの方が「今この瞬間感」が出て“刺さりやすい”。traitsは外しても成立しやすい。

---

## Question 3
traits も states も削っても149文字を超えた場合（verdict や ritual が長い場合）、どう対応しますか？

A) verdict や ritual を末尾から切り捨てて `…` をつける
B) そのケースは辞書設計で起こらないようにする（ritual は短く設計する）
C) Other (please describe after [Answer]: tag below)

[Answer]:

辞書設計で起こらないようにする（ritualは短め固定、verdictも短く）。
本文の主役は「儀式」なので、切り捨てで壊したくない。
※保険として実装側は最終手段でA（末尾…）を持ってもいいが、基本方針はB。
---

## Question 4
各辞書の初期エントリ数の目安を教えてください。

A) 少なめ（各5〜10件）でまず動かす、後で追加
B) ある程度しっかり（各20〜30件）用意してから動かす
C) Other (please describe after [Answer]: tag below)

[Answer]:

最初から 各20〜30件くらい入れておく。
5分おき投稿だとすぐローテが見えるので、初動から“毎回違う感”を出したい。

---

全て回答したら「回答しました」と教えてください。
