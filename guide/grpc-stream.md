# gRPC ストリームでイベントを受信する (/docs/guides/grpc-stream)



gRPC ストリーム方式では、mixi2 のサーバーに常時接続し、リアルタイムでイベントを受信します。外部公開 URL が不要なため、ローカル開発やプロトタイピングに適した受信方式です。

前提条件 [#前提条件]

* [SDK ガイド](/guides/sdk) で SDK のインストールと認証の実装が完了していること

gRPC ストリーム方式の特徴 [#grpc-ストリーム方式の特徴]

**メリット:**

* 外部公開 URL が不要
* ローカル開発で即座にテスト可能
* リアルタイム性が高い

**デメリット:**

* 常時接続が必要
* 接続が切れた場合、その間のイベントは失われる
* スケールアウトが難しい

**推奨シーン:** ローカル開発、プロトタイピング、単一インスタンスで十分な小規模アプリケーション

SDK による実装 [#sdk-による実装]

```go
package main

import (
    "context"
    "crypto/tls"
    "log"
    "os"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"

    "github.com/mixigroup/mixi2-application-sdk-go/auth"
    "github.com/mixigroup/mixi2-application-sdk-go/event/stream"
    application_streamv1 "github.com/mixigroup/mixi2-application-sdk-go/gen/go/social/mixi/application/service/application_stream/v1"
    modelv1 "github.com/mixigroup/mixi2-application-sdk-go/gen/go/social/mixi/application/model/v1"
)

type MyHandler struct{}

func (h *MyHandler) Handle(ctx context.Context, ev *modelv1.Event) error {
    log.Printf("Received event: %v", ev)
    return nil
}

func main() {
    authenticator, err := auth.NewAuthenticator(
        os.Getenv("CLIENT_ID"),
        os.Getenv("CLIENT_SECRET"),
        os.Getenv("TOKEN_URL"),
    )
    if err != nil {
        log.Fatal(err)
    }

    conn, err := grpc.NewClient(
        os.Getenv("STREAM_ADDRESS"),
        grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client := application_streamv1.NewApplicationServiceClient(conn)
    watcher := stream.NewStreamWatcher(client, authenticator)

    if err := watcher.Watch(context.Background(), &MyHandler{}); err != nil {
        log.Fatal(err)
    }
}
```

イベントハンドラ（`EventHandler` インターフェース）の詳細は [SDK ガイド - イベントハンドラの実装](/guides/sdk#イベントハンドラの実装) を参照してください。

再接続の仕様 [#再接続の仕様]

ストリーミング接続が切断された場合、SDK は自動的に再接続を試みます。

| 項目       | 値                   |
| -------- | ------------------- |
| 再接続方式    | 指数バックオフ（1秒, 2秒, 4秒） |
| 最大リトライ回数 | 3 回                 |

<Callout type="warning">
  再接続中に発生したイベントは失われます。イベントの厳密な到達保証が必要な場合は、[Webhook 方式](/guides/webhook)の使用を検討してください。
</Callout>

Ping イベント [#ping-イベント]

gRPC ストリーム方式では、イベントがない状態が 20 秒続くと Ping イベントが送信されます。Ping イベントは SDK 内部で処理されるため、`Handle` メソッドには渡されません。開発者が特別な対応をする必要はありません。

次のステップ [#次のステップ]

* [Webhook でイベントを受信する](/guides/webhook) - もう一つのイベント受信方式
* [イベント](/guides/events) - イベントの種類と構造
* [SDK ガイド](/guides/sdk) - SDK の基本（認証、イベントハンドラ）
* [API の使い方](/guides/api-usage) - イベントを受信した後の応答方法
