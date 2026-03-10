# Execution Plan — MIXI2 占いBot

## 詳細分析サマリー

### Change Impact Assessment
- **User-facing changes**: Yes — MIXI2 タイムラインへの定期投稿
- **Structural changes**: No — 既存システムなし（Greenfield）
- **Data model changes**: No — 外部 DB・永続化なし
- **API changes**: No — MIXI2 API は外部（読み取り専用利用）
- **NFR impact**: Low — 文字数制限対応のみ（149文字トリミング）

### Risk Assessment
- **Risk Level**: Low
- **Rollback Complexity**: Easy（GitHub Actions ワークフローを無効化するだけ）
- **Testing Complexity**: Simple（単発実行・ローカルで確認可能）

---

## Workflow Visualization

```
INCEPTION PHASE
  [DONE] Workspace Detection
  [DONE] Requirements Analysis
  [SKIP] Reverse Engineering      -- Greenfield のため不要
  [SKIP] User Stories             -- 単一ユーザー・単純用途のため不要
  [EXEC] Workflow Planning        -- 現在実行中
  [EXEC] Application Design       -- コンポーネント間インターフェース定義
  [SKIP] Units Generation         -- 単一ユニットのため不要

CONSTRUCTION PHASE
  [EXEC] Functional Design        -- 辞書構造・テンプレートアルゴリズム設計
  [SKIP] NFR Requirements         -- 文字数制限は要件定義済み・追加NFRなし
  [SKIP] NFR Design               -- NFR Requirements スキップに伴いスキップ
  [SKIP] Infrastructure Design    -- GitHub Actions は単純設定のため不要
  [EXEC] Code Generation          -- 全コード・ワークフロー生成
  [EXEC] Build and Test           -- ビルド・テスト手順作成

OPERATIONS PHASE
  [----] Operations               -- Placeholder（対象外）
```

---

## 実行フェーズ一覧

### INCEPTION PHASE
- [x] Workspace Detection — COMPLETED
- [x] Requirements Analysis — COMPLETED
- [x] Workflow Planning — IN PROGRESS
- [ ] Application Design — **EXECUTE**
  - **Rationale**: 3つのコンポーネント（fortune / mixi2 / main）のインターフェースを明確に定義する。特に生成ロジックと投稿クライアントの分離境界が重要。

### CONSTRUCTION PHASE
- [ ] Functional Design — **EXECUTE**
  - **Rationale**: 辞書（states/traits/verdicts/rituals/zodiac）の構造設計と、テンプレート合成・文字数トリミングのアルゴリズム設計が必要。
- [ ] NFR Requirements — **SKIP**
  - **Rationale**: 文字数制限（149文字）は要件定義済み。パフォーマンス・スケーラビリティの追加要件なし。
- [ ] NFR Design — **SKIP**
  - **Rationale**: NFR Requirements スキップに伴いスキップ。
- [ ] Infrastructure Design — **SKIP**
  - **Rationale**: GitHub Actions の cron 設定はシンプルであり、複雑なインフラ設計は不要。
- [ ] Code Generation — **EXECUTE** (ALWAYS)
  - **Rationale**: 全コード（Go ソース・GitHub Actions ワークフロー・README）の生成。
- [ ] Build and Test — **EXECUTE** (ALWAYS)
  - **Rationale**: ビルド手順・テスト手順・ローカル実行手順の作成。

### OPERATIONS PHASE
- [ ] Operations — PLACEHOLDER

---

## 実行順序

1. **Application Design** — コンポーネント設計・インターフェース定義
2. **Functional Design** — 辞書構造・テンプレートアルゴリズム詳細設計
3. **Code Generation** — 実装（Go コード + GitHub Actions）
4. **Build and Test** — ビルド・テスト手順

---

## 成功基準

- **Primary Goal**: GitHub Actions で5分ごとに MIXI2 へ占い投稿が実行される
- **Key Deliverables**:
  - `cmd/bot/main.go`（単発実行エントリポイント）
  - `internal/fortune/`（辞書・生成ロジック）
  - `internal/mixi2/`（CreatePost クライアント）
  - `.github/workflows/fortune.yml`（cron スケジュール）
  - `README.md`（セットアップ手順）
- **Quality Gates**:
  - 投稿本文が常に 149 文字以内に収まること
  - ローカルで `go run cmd/bot/main.go` 実行後に MIXI2 投稿が確認できること
  - GitHub Secrets 設定後に Actions が正常動作すること
