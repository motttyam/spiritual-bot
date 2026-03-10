# SDK ガイド (/docs/guides/sdk)



mixi2 Developer Platform では、アプリケーション開発を効率化するための公式 SDK を提供しています。SDK を使用すると、認証管理やイベント受信の実装を簡略化できます。

前提条件 [#前提条件]

* Go 1.24.6 以上
* 開発者登録が完了していること
* クライアント ID・クライアントシークレットを取得済みであること

インストール [#インストール]

```bash
go get github.com/mixigroup/mixi2-application-sdk-go
```

パッケージ構造 [#パッケージ構造]

SDK は以下のパッケージで構成されています。

| パッケージ           | 機能                                                   |
| --------------- | ---------------------------------------------------- |
| `auth`          | OAuth2 Client Credentials 認証（アクセストークンの取得・キャッシュ・自動更新） |
| `event`         | イベントハンドリングのインターフェース定義                                |
| `event/webhook` | HTTP Webhook サーバーによるイベント受信                           |
| `event/stream`  | gRPC ストリーミングによるイベント受信                                |
| `gen`           | Protocol Buffers 生成コード（型定義、gRPC クライアント等）             |

```
github.com/mixigroup/mixi2-application-sdk-go/
├── auth/             # 認証（OAuth2 Client Credentials）
├── event/            # EventHandler インターフェース定義
│   ├── webhook/      # Webhook サーバー
│   └── stream/       # gRPC ストリーミング
└── gen/              # protobuf 生成コード
```

SDK が提供する機能 [#sdk-が提供する機能]

認証管理 [#認証管理]

| 機能           | 説明                                                 |
| ------------ | -------------------------------------------------- |
| OAuth 2.0 認証 | Client Credentials フローによるアクセストークン取得                |
| トークンキャッシュ    | 有効期限の 1 分前までキャッシュし、自動更新                            |
| gRPC メタデータ付与 | `Authorization` ヘッダに `Bearer {access_token}` を自動設定 |

イベント受信 [#イベント受信]

SDK は 2 つのイベント受信方式をサポートしています。各方式の詳細な実装方法は個別のガイドを参照してください。

| 方式           | パッケージ           | ガイド                               |
| ------------ | --------------- | --------------------------------- |
| HTTP Webhook | `event/webhook` | [Webhook](/guides/webhook)        |
| gRPC ストリーム   | `event/stream`  | [gRPC ストリーム](/guides/grpc-stream) |

SDK は以下の処理を内部で行うため、開発者が個別に実装する必要はありません。

| 機能                    | 方式         | 説明                              |
| --------------------- | ---------- | ------------------------------- |
| Ed25519 署名検証          | Webhook    | リクエストの改ざん検知。検証失敗時は HTTP 401 を返却 |
| タイムスタンプ検証             | Webhook    | ±5 分以上ずれたリクエストを拒否（リプレイ攻撃対策）     |
| Webhook URL 検証リクエスト応答 | Webhook    | 3 種類の Ping に対して適切なレスポンスを自動返却    |
| 再接続                   | gRPC ストリーム | 接続切断時に指数バックオフで自動再接続（最大 3 回）     |
| Ping イベント処理           | 両方         | SDK 内部で処理し、`Handle` には渡さない      |

認証の実装 [#認証の実装]

`auth.Authenticator` は OAuth2 Client Credentials フローによるアクセストークンの取得・キャッシュ・自動更新を行います。

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/mixigroup/mixi2-application-sdk-go/auth"
)

func main() {
    authenticator, err := auth.NewAuthenticator(
        os.Getenv("CLIENT_ID"),
        os.Getenv("CLIENT_SECRET"),
        os.Getenv("TOKEN_URL"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // gRPC リクエスト用のコンテキストを取得
    ctx, err := authenticator.AuthorizedContext(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    // ctx を使って gRPC リクエストを送信
    _ = ctx
}
```

イベントハンドラの実装 [#イベントハンドラの実装]

イベントハンドラは `Handle` メソッドを実装したインターフェースです。Webhook・gRPC ストリームの両方式で共通して使用します。

```go
import modelv1 "github.com/mixigroup/mixi2-application-sdk-go/gen/go/social/mixi/application/model/v1"

type EventHandler interface {
    Handle(ctx context.Context, event *modelv1.Event) error
}
```

注意事項 [#注意事項]

| 項目             | 説明                                                           |
| -------------- | ------------------------------------------------------------ |
| 並列実行           | `Handle` は複数の goroutine から並列に呼び出される可能性があります。スレッドセーフに実装してください |
| Ping イベント      | SDK 内部で処理されるため、`Handle` には渡されません                             |
| Best-effort 処理 | `Handle` の結果にかかわらず、Webhook は常に `204 No Content` を返します        |
| エラーハンドリング      | `Handle` から返されたエラーはログに出力されますが、処理は継続します                       |

サンプルアプリケーション [#サンプルアプリケーション]

SDK の使い方を理解するためのサンプルアプリケーションを提供しています。実際に動作するコードを参考に開発を始めることができます。

プロジェクト構成 [#プロジェクト構成]

| ディレクトリ         | 説明                            |
| -------------- | ----------------------------- |
| `cmd/stream/`  | gRPC ストリーミングモードで動作するアプリケーション  |
| `cmd/webhook/` | HTTP Webhook モードで動作するアプリケーション |
| `api/`         | Vercel サーバーレス関数のエントリーポイント     |
| `handler/`     | イベントを処理するハンドラ                 |
| `config/`      | 環境変数から設定を読み込む                 |

環境変数 [#環境変数]

| 変数名                    | 必須 | 説明                                          |
| ---------------------- | -- | ------------------------------------------- |
| `CLIENT_ID`            | ○  | OAuth2 クライアント ID                            |
| `CLIENT_SECRET`        | ○  | OAuth2 クライアントシークレット                         |
| `TOKEN_URL`            | ○  | トークンエンドポイント URL                             |
| `API_ADDRESS`          | ○  | API サーバーアドレス                                |
| `STREAM_ADDRESS`       | △  | Stream サーバーアドレス ※ gRPC ストリーミング時のみ必須         |
| `SIGNATURE_PUBLIC_KEY` | △  | イベント署名検証用の公開鍵（Base64）※ Webhook・Vercel でのみ必須 |
| `PORT`                 |    | Webhook サーバーポート（デフォルト: `8080`）              |

対応イベント [#対応イベント]

サンプルアプリケーションは以下のイベントに反応します。

| イベント                               | 動作                         |
| ---------------------------------- | -------------------------- |
| `EVENT_TYPE_POST_CREATED`          | ポスト作成イベントをログ出力             |
| `EVENT_TYPE_CHAT_MESSAGE_RECEIVED` | 受信した DM のテキストをそのまま返信（Echo） |

使い方 [#使い方]

```bash
# リポジトリをクローン
git clone https://github.com/mixigroup/mixi2-application-sample-go.git
cd mixi2-application-sample-go

# 環境変数を設定
cp .env.example .env
# .env を編集して環境変数を設定
source .env
go mod tidy

# gRPC ストリームモード（ローカル開発向け）
go run cmd/stream/main.go

# HTTP Webhook モード（デプロイ向け）
go run cmd/webhook/main.go
```

<Callout type="info">
  ローカル開発では gRPC ストリームモードがおすすめです。外部公開 URL なしでイベントを受信できます。
</Callout>

SDK を使わない開発 [#sdk-を使わない開発]

SDK を使用せずに開発する場合は、Protocol Buffers 定義から直接コードを生成できます。

```bash
# API Proto リポジトリをクローン
git clone https://github.com/mixigroup/mixi2-api.git

# buf を使用してコード生成
cd mixi2-api
buf generate
```

<Callout type="warning">
  SDK を使用しない場合、認証管理、署名検証、再接続処理などを自前で実装する必要があります。
</Callout>

リポジトリ [#リポジトリ]

| リポジトリ                                                                    | 説明                  |
| ------------------------------------------------------------------------ | ------------------- |
| [SDK（Go）](https://github.com/mixigroup/mixi2-application-sdk-go)         | 公式 SDK              |
| [サンプルアプリケーション](https://github.com/mixigroup/mixi2-application-sample-go) | サンプルアプリケーション        |
| [API Proto](https://github.com/mixigroup/mixi2-api)                      | Protocol Buffers 定義 |

次のステップ [#次のステップ]

* [Webhook でイベントを受信する](/guides/webhook) - Webhook 方式の実装・署名検証・デプロイ
* [gRPC ストリームでイベントを受信する](/guides/grpc-stream) - gRPC 方式の実装・再接続
* [イベント](/guides/events) - イベントの種類・構造・配信仕様
* [API の使い方](/guides/api-usage) - 各 API の使用方法
