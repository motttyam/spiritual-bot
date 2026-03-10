# 要件定義書 — MIXI2 占いBot

## インテント分析

| 項目 | 内容 |
|------|------|
| ユーザーリクエスト | GitHub Actions の cron で5分に1回 MIXI2 に占い投稿を行う Bot の作成 |
| リクエストタイプ | New Project（Greenfield） |
| スコープ | System-wide（新規プロジェクト一式） |
| 複雑度 | Moderate |

---

## 機能要件

### FR-1: 占いコンテンツ生成
- ランダムに1つの星座（12星座：牡羊座〜魚座）を選択する
- 以下の2行フォーマットで投稿本文を生成する:
  - **1行目**: 呼びかけ（超具体的な状態・特徴の組み合わせ）＋結果（断言形式）
  - **2行目**: ラッキー行動（儀式：無害・具体・回数つき）
- 辞書（states / traits / verdicts / rituals / zodiac）＋テンプレート合成で本文を生成する
- LLM は使用しない（将来的に挟めるよう生成ロジックと投稿ロジックを分離する）

**例:**
```
あごひげをたくわえていて、今歯をみがいている乙女座のあなた、今日、運勢が一回だけ成仏します。
ラッキー行動：コップの水に「ありがとう」を3回言ってから一口だけ飲む
```

### FR-2: 投稿コンテンツのトーン
- めちゃスピリチュアルで、やりすぎて面白い方向に振り切る
- ユーザー本人に対して直接「死ぬ」等の断定は避ける
- 代替表現を使用する（例: 「運勢が成仏する」「不運が転生する」「現実がログアウトする」「宇宙が優先処理する」）

### FR-3: 文字数制限対応
- MIXI2 の投稿文字数制限: **149文字**
- 超過時は1行目の修飾（状態・特徴）を削って短縮する
- 2行目（儀式）はできるだけ残す
- ログに生成した本文を出力する

### FR-4: MIXI2 への投稿
- `CreatePost` RPC を使用してテキスト投稿を行う
- 投稿は1回のみ実行して終了する（単発実行）
- OAuth 2.0 認証によるアクセス（MIXI2 公式 Go SDK を使用）

### FR-5: 定期実行
- GitHub Actions の `schedule` トリガーで `*/5 * * * *`（5分ごと）実行する
- 実行ごとに1件の占い投稿を行って終了する

---

## 非機能要件

### NFR-1: 認証・セキュリティ
- MIXI2 の認証情報（クライアントクレデンシャル等）は GitHub Secrets に格納する
- 環境変数経由でプログラムに渡す
- シークレットをコードにハードコードしない

### NFR-2: 拡張性
- 生成ロジック（`internal/fortune/`）と投稿クライアント（`internal/mixi2/`）を分離する
- 将来的なLLM統合、メンション対応などのイベント駆動拡張を想定した構造にする

### NFR-3: 保守性
- ログに生成した投稿本文を出力する
- GitHub Actions のログで動作確認できる

---

## 技術スタック

| 項目 | 選定内容 |
|------|---------|
| 言語 | Go |
| MIXI2 SDK | mixi2-application-sdk-go（公式） |
| 実行環境 | GitHub Actions（schedule/cron） |
| 認証情報管理 | GitHub Secrets → 環境変数 |
| LLM | なし（ルールベース） |

---

## ディレクトリ構成

```text
<project-root>/
├── cmd/
│   └── bot/
│       └── main.go               # エントリポイント（単発実行）
├── internal/
│   ├── fortune/                  # 占いコンテンツ生成ロジック
│   │   ├── generator.go          # テンプレート合成・文字数調整
│   │   └── dictionary.go         # states/traits/verdicts/rituals/zodiac 辞書
│   └── mixi2/                    # MIXI2 API クライアント
│       └── client.go             # CreatePost ラッパー
├── .github/
│   └── workflows/
│       └── fortune.yml           # cron スケジュール定義
├── go.mod
├── go.sum
└── README.md                     # Secrets 設定方法・ローカル実行方法
```

---

## GitHub Actions 仕様

```yaml
# .github/workflows/fortune.yml の概要
on:
  schedule:
    - cron: '*/5 * * * *'
```

- 必要な GitHub Secrets:
  - `MIXI2_CLIENT_ID`
  - `MIXI2_CLIENT_SECRET`（または SDK が要求する認証情報）

---

## API 制約（MIXI2）

| 項目 | 制限 |
|------|------|
| 投稿文字数上限 | 149 文字 |
| 使用 RPC | `CreatePost`（テキストのみ） |
| 認証方式 | OAuth 2.0 |
| プロトコル | gRPC（Go SDK 経由） |

---

## スコープ外（初期リリース）

- イベント受信（Webhook / gRPC Stream）
- メンション対応
- DM 対応
- LLM による文章生成
- 複数投稿や画像添付
