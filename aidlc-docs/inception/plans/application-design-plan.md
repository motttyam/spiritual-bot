# Application Design Plan — MIXI2 占いBot

## 実行チェックリスト

- [x] コンポーネント定義の作成（components.md）
- [x] メソッドシグネチャの定義（component-methods.md）
- [x] サービス層の定義（services.md）
- [x] 依存関係の定義（component-dependency.md）
- [x] 統合ドキュメントの作成（application-design.md）

---

## 設計判断に関する質問

以下の質問に回答してください。各 `[Answer]:` タグの後に選択肢の記号を入力してください。

---

## Question 1
辞書データ（states/traits/verdicts/rituals/zodiac）の管理方式はどれにしますか？

A) Go コードに直接定義（`[]string{...}` スライス）— シンプル・ファイル不要
B) JSON ファイル + `//go:embed`（コンパイル時埋め込み）— 内容の編集がしやすい
C) Other (please describe after [Answer]: tag below)

[Answer]:

辞書は増やしたり調整したりが頻繁に起きるはずなので、**JSON + //go:embed**で「コード触らずに辞書更新」できる形が相性いい。

---

## Question 2
API 投稿失敗時の挙動はどうしますか？

A) エラーログを出力して `os.Exit(1)` で終了（GitHub Actions でジョブ失敗として記録される）
B) エラーログを出力して `os.Exit(0)` で終了（失敗しても Actions は成功扱い）
C) Other (please describe after [Answer]: tag below)

[Answer]:

投稿失敗は運用上ちゃんと気づきたいので、os.Exit(1)でジョブ失敗にしてGitHub Actions上で赤くする。
（将来、安定運用で一時障害を許容したくなったらリトライやBに切り替えればOK）

---

## Question 3
認証情報の環境変数名はどうしますか？

A) MIXI2 SDK のデフォルト変数名に従う（SDK ドキュメント準拠）
B) 自分で名前を決めたい（例: `MIXI2_CLIENT_ID` / `MIXI2_CLIENT_SECRET`）
C) Other (please describe after [Answer]: tag below)

[Answer]:

GitHub Secretsとworkflowの可読性を優先して、独自命名にするのがラク。例：

MIXI2_CLIENT_ID

MIXI2_CLIENT_SECRET

MIXI2_TOKEN_URL

MIXI2_API_ADDRESS

---

全て回答したら「回答しました」と教えてください。
