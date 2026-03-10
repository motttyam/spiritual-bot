```text
# Requirements Clarification Questions — MIXI2 Bot（回答＋補足）

MIXI2 Bot の要件を明確にするための質問です。
各質問の `[Answer]:` タグの後に、選択肢の記号（A, B など）を入力してください。
選択肢に合うものがない場合は最後の選択肢（Other）を選び、その後に詳細を記述してください。

---

## Question 1
Bot の主な用途・目的は何ですか？

A) 自動返信Bot（メンションやDMに自動応答する）
B) 情報配信Bot（定期的に投稿や通知を行う）
C) 対話型Bot（会話フロー・コマンドに基づいて応答する）
D) 複合用途（上記の複数機能を持つ）
E) Other (please describe after [Answer]: tag below)

[Answer]: B

補足:
- 初期は「定期投稿」のみでOK。
- 5分に1回、MIXI2に占い投稿を行う。

---

## Question 2
Bot が反応するイベントはどれですか？（複数可 — 該当するものをカンマ区切りで記入）

A) メンション（@Bot 宛の投稿）
B) DM（ダイレクトメッセージ）
C) フォロー・アンフォロー
D) その他のイベント（投稿、スタンプなど）
E) Other (please describe after [Answer]: tag below)

[Answer]: E
Other:
- イベント反応は使わず、GitHub Actions の schedule（cron）で定期実行して投稿する

---

## Question 3
使用するプログラミング言語はどれですか？

A) Go（公式SDK あり）
B) TypeScript / JavaScript（コミュニティSDK あり）
C) その他の言語（SDKなし、手動実装）
D) Other (please describe after [Answer]: tag below)

[Answer]: A

補足:
- 公式SDKがある Go を採用。
- CLI的に単発実行（1回投稿して終了）できる構成にする。

---

## Question 4
イベント受信方式はどちらにしますか？

A) gRPC Stream（ローカル開発・プロトタイピング向け。サーバー不要）
B) Webhook（本番運用向け。HTTPS エンドポイントが必要）
C) まだ決めていない / 両方試したい
D) Other (please describe after [Answer]: tag below)

[Answer]: D
Other:
- 受信はしない（定期実行のみ）
- API呼び出しで投稿（CreatePost）だけ行う

---

## Question 5
デプロイ先・実行環境はどこですか？

A) ローカル環境のみ（開発・検証用）
B) クラウド（AWS / GCP / Azure など）
C) VPS・自前サーバー
D) サーバーレス（AWS Lambda など）
E) Other (please describe after [Answer]: tag below)

[Answer]: E
Other:
- GitHub Actions（schedule/cron）で5分に1回実行

---

## Question 6
Bot に AI / LLM 機能（自然言語応答など）を組み込みますか？

A) はい — AWS Bedrock / Claude API など特定のサービスを使いたい
B) はい — まだどのサービスを使うか未定
C) いいえ — ルールベースの固定応答のみ
D) Other (please describe after [Answer]: tag below)

[Answer]: C

補足:
- LLMは使わず、ルールベース（辞書＋テンプレ合成）で面白さを作る。
- 将来的に「言い回し整形だけ」LLMを挟めるように、生成ロジックと投稿ロジックは分離しておく。

---

## Question 7
セキュリティ・認証に関する要件はありますか？

A) MIXI2 SDK の Webhook 署名検証のみで十分
B) 追加の認証・認可機能が必要（管理者コマンドなど）
C) 特に要件なし
D) Other (please describe after [Answer]: tag below)

[Answer]: C

補足:
- Webhookは使わない。
- 認証はMIXI2のクライアントクレデンシャル等でAPI投稿ができればOK。
- シークレットは GitHub Secrets に格納。

---

## Question 8
将来的なスケールアップや拡張の予定はありますか？

A) ない（個人・小規模利用）
B) ある（将来的に機能追加・複数サーバー対応を想定）
C) 未定
D) Other (please describe after [Answer]: tag below)

[Answer]: B

補足:
- 将来的に「メンションで占う」などのイベント駆動に拡張する可能性あり。
- ただし初期は定期投稿のみ。

---

# 追加要件（Claude Code向け実装指示）

## 目的
- GitHub Actions の cron で 5分に1回 起動し、MIXI2 に 占い投稿を1件 行って終了する。

## 占い投稿のフォーマット（最小・2行固定）
- 1行目：呼びかけ（超具体）＋結果（断言）
- 2行目：ラッキー行動（儀式：無害・具体・回数つき）

例:
1行目: あごひげをたくわえていて、今歯をみがいている乙女座のあなた、今日、運勢が一回だけ成仏します。
2行目: ラッキー行動：コップの水に「ありがとう」を3回言ってから一口だけ飲む

## 星座
- 12星座（牡羊座〜魚座）を使う
- 星座決定は「毎回ランダム」でも「日替わり」でもよいが、最初はランダムでOK

## 面白さの方針（重要）
- めちゃスピリチュアルで、やりすぎて面白い方向に振り切る
- ただし、ユーザー本人に対して直接「死ぬ」等の断定は避ける
  代替例:
  - 運勢が成仏する
  - 不運が転生する
  - 現実がログアウトする
  - 宇宙が優先処理する

## 実装の最小要件
- Go プロジェクトとして作成
- 辞書（states/traits/verdicts/rituals/zodiac）＋テンプレ合成で本文生成
- MIXI2 API で投稿（CreatePost）
- 文字数制限がある場合に備え、安全に短縮する処理を入れる
  - 儀式（2行目）はできるだけ残す
  - 超過時は1行目の修飾（状態/特徴）を削る等で調整
- ログに「生成した本文」を出力

## GitHub Actions
- .github/workflows/fortune.yml を作成し、*/5 * * * * で実行
- 必要な認証情報は GitHub Secrets 経由で環境変数として渡す
- 単発実行（1回投稿して終了）でOK

## 望ましいディレクトリ構成（例）
- cmd/bot/main.go
- internal/fortune/（辞書＆生成ロジック）
- internal/mixi2/（投稿クライアント）
- .github/workflows/fortune.yml
- README.md（Secretsの設定方法・ローカル実行方法）
```
