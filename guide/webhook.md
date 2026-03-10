# Webhook でイベントを受信する (/docs/guides/webhook)



Webhook 方式では、mixi2 からアプリケーションサーバーの HTTPS エンドポイントにイベントが POST されます。本番環境やサーバーレス環境でのデプロイに適した受信方式です。

前提条件 [#前提条件]

* [SDK ガイド](/guides/sdk) で SDK のインストールと認証の実装が完了していること
* HTTPS で公開された URL が用意できること

Webhook 方式の特徴 [#webhook-方式の特徴]

**メリット:**

* サーバーレス環境に適合
* スケールアウトが容易
* 接続管理が不要

**デメリット:**

* HTTPS で公開された URL が必要
* 有効な CA 証明書が必要（自己署名証明書は不可）

<Callout type="info">
  1 つのアプリケーションに対して、同時に複数の Webhook URL を設定することはできません。
</Callout>

SDK による実装 [#sdk-による実装]

Webhook サーバーは以下のエンドポイントを提供します。

| エンドポイント        | 説明             |
| -------------- | -------------- |
| `POST /events` | イベント受信エンドポイント  |
| `GET /healthz` | ヘルスチェックエンドポイント |

```go
package main

import (
    "context"
    "crypto/ed25519"
    "encoding/base64"
    "log"
    "os"

    "github.com/mixigroup/mixi2-application-sdk-go/event/webhook"
    modelv1 "github.com/mixigroup/mixi2-application-sdk-go/gen/go/social/mixi/application/model/v1"
)

type MyHandler struct{}

func (h *MyHandler) Handle(ctx context.Context, ev *modelv1.Event) error {
    log.Printf("Received event: %v", ev)
    return nil
}

func main() {
    publicKeyBase64 := os.Getenv("SIGNATURE_PUBLIC_KEY")
    publicKey, err := base64.StdEncoding.DecodeString(publicKeyBase64)
    if err != nil {
        log.Fatal(err)
    }

    server := webhook.NewServer(
        ":8080",
        ed25519.PublicKey(publicKey),
        &MyHandler{},
    )

    if err := server.Start(); err != nil {
        log.Fatal(err)
    }
}
```

イベントハンドラ（`EventHandler` インターフェース）の詳細は [SDK ガイド - イベントハンドラの実装](/guides/sdk#イベントハンドラの実装) を、環境変数の一覧は [SDK ガイド - 環境変数](/guides/sdk#環境変数) を参照してください。

署名検証 [#署名検証]

Webhook 方式では、受信したリクエストが mixi2 サーバーから送信された正当なものであることを、Ed25519 署名で検証する必要があります。**署名検証は必須です。**

<Callout type="info">
  SDK を使用している場合は、イベント受信時の署名検証と Webhook URL 登録時の検証リクエスト応答がいずれも SDK 内部で処理されます。
</Callout>

SDK を使用しない場合は、以下の仕様に基づいて自前で実装してください。

検証に使用するヘッダとパラメータ [#検証に使用するヘッダとパラメータ]

| 項目         | 値                                                   |
| ---------- | --------------------------------------------------- |
| 署名ヘッダ      | `x-mixi2-application-event-signature`（Base64 エンコード） |
| タイムスタンプヘッダ | `x-mixi2-application-event-timestamp`（Unix 秒）       |
| 署名対象       | リクエストボディ + タイムスタンプ                                  |
| 許容時刻ズレ     | ±300 秒（5 分）                                         |
| 公開鍵形式      | Base64 エンコード                                        |

URL 検証時の署名検証 [#url-検証時の署名検証]

Webhook URL を登録して「有効化」すると、3 種類の Ping リクエストがランダムな順で送信されます。署名検証が正しく動作しているかを確認するため、それぞれ以下のレスポンスが期待されます。

| リクエスト     | 期待されるレスポンス |
| --------- | ---------- |
| 正しい署名     | HTTP 204   |
| 不正な署名     | HTTP 401   |
| 古いタイムスタンプ | HTTP 401   |

URL の登録手順や検証状態の確認方法は [アプリケーション開発 - Webhook の設定](/guides/application#webhook-の設定) を参照してください。

タイムアウトとリトライ [#タイムアウトとリトライ]

Webhook 方式では、以下のポリシーが適用されます。

| 項目     | 値      |
| ------ | ------ |
| タイムアウト | 3 秒    |
| リトライ回数 | 最大 3 回 |
| リトライ間隔 | 30 秒   |

イベントを受信したら速やかにレスポンスを返し、アプリケーションのロジックは非同期で処理してください。3 秒以内にアプリケーションの処理を完了させる必要はありません。

Webhook URL の設定 [#webhook-url-の設定]

URL の要件 [#url-の要件]

* HTTPS 必須
* 有効な CA 証明書（自己署名証明書は不可）
* SDK を使用している場合、パスは `/events` です（例: `https://YOUR_HOST_NAME/events`）

URL の登録 [#url-の登録]

Webhook URL の登録手順は [アプリケーション開発 - Webhook の設定](/guides/application#webhook-の設定) を参照してください。

デプロイ [#デプロイ]

Webhook 方式を使用するには、HTTPS エンドポイントを用意できる環境が必要です。クラウドサービス・自前のサーバーなど任意の環境にデプロイしてください。

* [サンプルアプリケーション](https://github.com/mixigroup/mixi2-application-sample-go)には [Vercel](https://vercel.com) にデプロイするためのコードが含まれています。README の手順に沿ってすぐにデプロイが可能です。
* 構成・環境変数の詳細は [SDK ガイド - サンプルアプリケーション](/guides/sdk#サンプルアプリケーション) を参照してください。

<Callout type="warning">
  **レスポンス後にプロセスが終了する環境での注意**

  [タイムアウトとリトライ](/guides/webhook#タイムアウトとリトライ) で案内している「速やかにレスポンスを返し、ロジックは非同期で処理する」というやり方は、レスポンスを返した時点でプロセスが終了する環境では行えません。そのような環境にデプロイする場合は、3 秒以内に同期的に処理を完了できるアプリケーションが適しています。
</Callout>

次のステップ [#次のステップ]

* [gRPC ストリームでイベントを受信する](/guides/grpc-stream) - もう一つのイベント受信方式
* [イベント](/guides/events) - イベントの種類と構造
* [SDK ガイド](/guides/sdk) - SDK の基本（認証、イベントハンドラ）
* [アプリケーション開発](/guides/application) - Webhook URL の登録手順
* [API の使い方](/guides/api-usage) - イベントを受信した後の応答方法
