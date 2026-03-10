# アプリケーション開発 (/docs/guides/application)



このガイドでは、アプリケーション設定から、デプロイ、テストまでの流れを解説します。

前提条件 [#前提条件]

* [開発者登録](/getting-started/registration)が完了していること
* [クイックスタート](/getting-started/quickstart)でアプリケーションを作成済みであること

イベント受信方式 [#イベント受信方式]

アプリケーションはイベントを受信して処理を行います。2 つのイベント受信方式から選択できます。

| 方式           | 推奨シーン           |
| ------------ | --------------- |
| gRPC ストリーム   | ローカル開発、プロトタイピング |
| HTTP Webhook | 本番環境、サーバーレス環境   |

各方式の詳細は [Webhook](/guides/webhook)、[gRPC ストリーム](/guides/grpc-stream) をそれぞれ参照してください。

Webhook の設定 [#webhook-の設定]

Webhook 方式を使用する場合、mixi2 Developer Platform で URL を登録し、検証を完了させる必要があります。

設定手順 [#設定手順]

1. mixi2 Developer Platform の管理画面でアプリケーションの「Webhook」を開きます
2. Webhook URL を登録します
3. 「有効化」ボタンを押して接続確認を実行します

<Callout type="info">
  URL の要件、署名検証の仕様、検証の仕組みの詳細は [Webhook でイベントを受信する](/guides/webhook) を参照してください。
</Callout>

URL の検証状態 [#url-の検証状態]

| 状態   | 説明               |
| ---- | ---------------- |
| 無効   | 初期状態、または無効化された状態 |
| 検証中  | 有効化ボタン押下後、検証処理中  |
| 検証失敗 | 検証に失敗した状態        |
| 検証済み | 正常に検証が完了した状態     |

<Callout type="warning">
  「検証済み」から「無効化」にするとステータスは「無効」に戻ります。エンドポイントを再設定した場合も「無効」に戻るため、再度「有効化」が必要です。
</Callout>

環境変数 [#環境変数]

アプリケーションサーバーの動作には、認証情報や接続先情報を環境変数として設定する必要があります。取得方法と設定手順は[クイックスタート](/getting-started/quickstart#step-2-認証情報を取得する)を参照してください。

デプロイ [#デプロイ]

Webhook 方式を使用する場合、HTTPS に対応した URL でリクエストを受信できる環境が必要です。デプロイの詳細は [Webhook でイベントを受信する - デプロイ](/guides/webhook#デプロイ) を参照してください。

エラーハンドリング [#エラーハンドリング]

API 呼び出し時にエラーが発生した場合は、エラーの種類に応じて適切に対処してください。

| エラー種別           | 対処法                           |
| --------------- | ----------------------------- |
| 認証エラー           | トークンを再取得してリトライ（SDK は自動で対応）    |
| レート制限超過         | `retry-after` ヘッダーに従って待機後リトライ |
| サーバーエラー         | 指数バックオフでリトライ（最大 3 回程度）        |
| クライアントエラー (4xx) | リトライせずログ出力、リクエスト内容を確認         |

レート制限の詳細（制限値、レスポンスヘッダー、エラーレスポンス）は[レート制限](/reference/rate-limits)を参照してください。

テスト方法 [#テスト方法]

開発中のテストは以下の手順で行います。

1. gRPC ストリーム方式でアプリケーションを起動
2. mixi2 アプリから実際にメンション・DM を送信
3. ログでイベント受信と処理結果を確認

<Callout type="info">
  テスト用のサンドボックス環境やイベントシミュレート機能は現在提供されていません。
</Callout>

開発リソース [#開発リソース]

アプリケーション開発に必要なリソースは GitHub で公開しています。詳細は [GitHub リポジトリ](/resource/github)を参照してください。

| リポジトリ                                                                                   | 説明                       |
| --------------------------------------------------------------------------------------- | ------------------------ |
| [mixi2-api](https://github.com/mixigroup/mixi2-api)                                     | API 定義（Protocol Buffers） |
| [mixi2-application-sdk-go](https://github.com/mixigroup/mixi2-application-sdk-go)       | Go 向け公式 SDK              |
| [mixi2-application-sample-go](https://github.com/mixigroup/mixi2-application-sample-go) | サンプルアプリケーション             |

次のステップ [#次のステップ]

* [Webhook でイベントを受信する](/guides/webhook) - Webhook 方式の実装・署名検証・デプロイ
* [gRPC ストリームでイベントを受信する](/guides/grpc-stream) - gRPC 方式の実装・再接続
* [イベント](/guides/events) - イベントの種類と構造
* [SDK ガイド](/guides/sdk) - SDK の基本（認証、イベントハンドラ）
* [API の使い方](/guides/api-usage) - ポストの作成、DM の送信、メディアのアップロード
* [API リファレンス](/reference/api-document) - API の完全な仕様
* [レート制限](/reference/rate-limits) - API のレート制限
