# よくある質問 (/docs/resource/faq)



import { Accordion, Accordions } from 'fumadocs-ui/components/accordion';
import { Callout } from 'fumadocs-ui/components/callout';

アカウント・ログイン [#アカウントログイン]

<Accordions>
  <Accordion title="ログインできません">
    以下のパターンが考えられます。

    | 原因                   | 対処法                                            |
    | -------------------- | ---------------------------------------------- |
    | 開発者登録が完了していない        | [開発者登録](/getting-started/registration)を行ってください |
    | 別の MIXI ID でログインしている | 開発者登録時に使用した MIXI ID でログインしてください                |

    <Callout type="info">
      mixi2 に登録しているアカウントで、別途[開発者登録](/getting-started/registration)が必要です。
    </Callout>

    <Callout type="warning">
      複数のメールアドレスで mixi2 を利用している場合、それぞれ別の MIXI ID となります。開発者登録した MIXI ID と異なる MIXI ID でログインしようとするとエラーになります。

      また、株式会社MIXI が提供する他サービスに別の MIXI ID でログイン済みの場合、その MIXI ID で認証されてエラーになることがあります。

      現在ログイン中の MIXI ID は [MIXI ID セキュリティ設定](https://account.mixi.com/id/security)から確認できます。
    </Callout>
  </Accordion>

  <Accordion title="MIXI ID とは何ですか？">
    mixi2 アカウントと MIXI ID の関係は、以下のヘルプページを参照してください。

    [mixi2アカウントとMIXI IDの関係](https://support.mixi.social/support/solutions/articles/154000212165-mixi2%E3%82%A2%E3%82%AB%E3%82%A6%E3%83%B3%E3%83%88%E3%81%A8mixi-id%E3%81%AE%E9%96%A2%E4%BF%82)
  </Accordion>

  <Accordion title="開発者登録の審査にはどのくらい時間がかかりますか？">
    申請の到着順に審査を実施しますが、順番が前後する場合があります。システムの状況により、審査に時間がかかる場合や、新規受付を一時停止する場合があります。

    詳細は[開発者登録](/getting-started/registration)をご確認ください。
  </Accordion>

  <Accordion title="審査完了の通知はどこに届きますか？">
    mixi2 上で [mixi2 公式アカウント](https://mixi.social/@mixi2)からメッセージで審査完了の連絡が届きます。

    <Callout type="warning">
      メッセージはアプリ版でのみ確認できます（Web 版では確認できません）。
    </Callout>
  </Accordion>

  <Accordion title="ログインするには電話番号が必要ですか？">
    はい、電話番号の登録が必要です。

    mixi2 はメールアドレスのみでアカウントを作成できますが、mixi2 Developer Platform を利用するには電話番号の登録が必要です。初回ログイン時に電話番号が未登録の場合は、登録を求められます。
  </Accordion>
</Accordions>

イベント受信 [#イベント受信]

<Accordions>
  <Accordion title="gRPC ストリームと Webhook のどちらを使うべきですか？">
    用途に応じて選択してください。

    | 方式         | 推奨シーン                       |
    | ---------- | --------------------------- |
    | gRPC ストリーム | ローカル開発、プロトタイピング、小規模アプリケーション |
    | Webhook    | 本番環境、サーバーレス環境               |

    詳細は [Webhook](/guides/webhook)、[gRPC ストリーム](/guides/grpc-stream) の各ガイドをご確認ください。
  </Accordion>

  <Accordion title="Webhook URL には何を設定すればよいですか？">
    以下の要件を満たす URL を設定してください。

    * HTTPS の公開 URL であること
    * 有効な CA 証明書を使用していること（自己署名証明書は不可）
    * SDK を使用している場合、パスは `/events` になります（例: `https://YOUR_HOST_NAME/events`）
  </Accordion>

  <Accordion title="Webhook のタイムアウトは何秒ですか？">
    3 秒です。イベントを受信したら速やかにレスポンスを返し、アプリケーションのロジックは非同期で処理してください。

    <Callout type="info">
      SDK を使用している場合は、レスポンスが自動で返されるため、タイムアウトを意識する必要はありません。
    </Callout>
  </Accordion>

  <Accordion title="Webhook が失敗した場合、リトライされますか？">
    はい、リトライされます。

    | 項目     | 値    |
    | ------ | ---- |
    | リトライ回数 | 3 回  |
    | リトライ間隔 | 30 秒 |
  </Accordion>
</Accordions>

API 利用 [#api-利用]

<Accordions>
  <Accordion title="ポストの最大文字数は？">
    149 文字です。
  </Accordion>

  <Accordion title="ポストに添付できるメディアの最大数は？">
    4 件までです。
  </Accordion>

  <Accordion title="返信と引用を同時に指定できますか？">
    いいえ、同時に指定できません。

    `in_reply_to_post_id` と `quoted_post_id` はどちらか一方のみ指定可能です。
  </Accordion>

  <Accordion title="DM を先送りで送ることはできますか？">
    いいえ、先送りはできません。

    ユーザーからの DM を受信した後のみ返信が可能です。
  </Accordion>

  <Accordion title="メディアのアップロードサイズ制限は？">
    | メディアタイプ | 最大サイズ |
    | ------- | ----- |
    | 画像      | 15MB  |
    | 動画      | 50MB  |
  </Accordion>

  <Accordion title="メディアアップロードの有効期限は？">
    | メディアタイプ | 有効期限  |
    | ------- | ----- |
    | 画像      | 200 秒 |
    | 動画      | 600 秒 |
  </Accordion>

  <Accordion title="メディアアップロードが失敗した場合は？">
    メディアは再利用できません。`InitiatePostMediaUpload` からやり直してください。
  </Accordion>

  <Accordion title="スタンプはどのポストにも付与できますか？">
    いいえ、以下の制限があります。

    | 制限項目     | 内容                      |
    | -------- | ----------------------- |
    | 対象ポスト    | アプリケーションにメンションしているポストのみ |
    | 使用可能スタンプ | 公式スタンプのみ                |
  </Accordion>

  <Accordion title="一度付与したスタンプを取り消せますか？">
    いいえ、取り消せません。

    現状、スタンプの取り消し機能は提供されていません。
  </Accordion>

  <Accordion title="自分（Bot）の投稿に対するイベントを受信しますか？">
    いいえ、受信しません。

    自身の投稿のイベントは送信されないように mixi2 側で制御されています。
  </Accordion>
</Accordions>

SDK [#sdk]

<Accordions>
  <Accordion title="Go 以外の SDK はありますか？">
    現状は Go のみです。
  </Accordion>

  <Accordion title="SDK を使わずに開発できますか？">
    はい、可能です。

    [mixi2-api](https://github.com/mixigroup/mixi2-api) リポジトリから proto ファイルを取得し、buf コマンドでコード生成できます。

    <Callout type="info">
      SDK を使用しない場合は、認証管理や署名検証などを自前で実装する必要があります。
    </Callout>
  </Accordion>
</Accordions>

Webhook 署名検証 [#webhook-署名検証]

<Accordions>
  <Accordion title="Webhook の署名検証は必須ですか？">
    はい、必須です。

    SDK を使用している場合は自動で検証されます。SDK を使用しない場合は、署名検証を自前で実装する必要があります。

    Webhook URL が設定された際に、アプリケーションサーバーが正しく実装されているか確認するために、正しい Ping イベントと無効な Ping イベントの両方が送信されます。署名検証が正しくハンドリングされている必要があります。
  </Accordion>

  <Accordion title="署名検証用の公開鍵はどこで取得できますか？">
    mixi2 Developer Platform のアプリケーション設定画面から取得できます。
  </Accordion>
</Accordions>

テスト・開発環境 [#テスト開発環境]

<Accordions>
  <Accordion title="テスト用のサンドボックス環境はありますか？">
    現状、サンドボックス環境は提供されていません。
  </Accordion>

  <Accordion title="イベントをシミュレートする API はありますか？">
    現状、イベントシミュレート用の API やツールは提供されていません。
  </Accordion>

  <Accordion title="ローカル開発でのテスト方法は？">
    1. gRPC ストリーム方式で接続
    2. mixi2 アプリから実際にメンションや DM を送信

    詳細は[クイックスタート](/getting-started/quickstart)をご確認ください。
  </Accordion>
</Accordions>

関連ページ [#関連ページ]

* [開発者登録](/getting-started/registration) - 開発者登録の詳細
* [クイックスタート](/getting-started/quickstart) - 最初のアプリケーション作成
* [SDK ガイド](/guides/sdk) - SDK の使い方
* [API の使い方](/guides/api-usage) - 各 API の使用方法
* [イベント](/guides/events) - イベントの種類と構造
* [Webhook](/guides/webhook) - Webhook 方式の実装
* [gRPC ストリーム](/guides/grpc-stream) - gRPC 方式の実装
* [アプリケーション開発](/guides/application) - アプリケーション構築の詳細
