# クイックスタート (/docs/getting-started/quickstart)











このガイドでは、mixi2 Developer Platform でアプリケーション（Bot）を作成し、イベントを受信して応答する基本的な流れを説明します。

前提条件 [#前提条件]

* [開発者登録](/getting-started/registration)が完了していること
* [アプリケーションの概念](/getting-started/concepts)を理解していること

Step 1: アプリケーションを作成する [#step-1-アプリケーションを作成する]

mixi2 Developer Platform にログインし、「新規アプリケーション」をクリックします。

<img alt="mixi2 Developer Platform ホーム画面" src={__img0} placeholder="blur" />

以下の情報を入力してアプリケーションを作成します。

<img alt="mixi2 Developer Platform アプリケーション作成画面" src={__img1} placeholder="blur" />

| 項目  | 説明                                                 |
| --- | -------------------------------------------------- |
| ID  | mixi2 上で表示される ID です。ユーザー全体でユニークである必要があり、後から変更できません |
| 表示名 | アプリケーションの表示名です。後から変更できます                           |

<Callout type="info">
  アプリケーションを作成すると、mixi2 上に指定した ID のアカウントが生成されます。mixi2 アプリで `@ID` を検索すると、アカウントが作成されていることが確認できます。
</Callout>

Step 2: 認証情報を取得する [#step-2-認証情報を取得する]

アプリケーションを作成すると、アプリケーション詳細画面に遷移します。

<img alt="mixi2 Developer Platform アプリケーション詳細画面" src={__img2} placeholder="blur" />

サイドバーで「認証情報」を選択し、 Client Secret を生成します。

<img alt="mixi2 Developer Platform アプリケーション認証情報画面" src={__img3} placeholder="blur" />

以下の情報が、後の手順で必要になります。

OAuth 2.0 クライアント認証用 [#oauth-20-クライアント認証用]

| 項目            | 説明                              |
| ------------- | ------------------------------- |
| Client ID     | OAuth 2.0 認証に使用するクライアント ID      |
| Client Secret | OAuth 2.0 認証に使用するシークレット。再発行可能です |

接続先情報 [#接続先情報]

| 項目             | 説明                       |
| -------------- | ------------------------ |
| Token URL      | アクセストークン取得用のエンドポイント URL  |
| Stream Address | gRPC ストリーミング接続用のサーバーアドレス |

<Callout type="warning">
  Client Secret は秘密情報です。ソースコードにハードコーディングしたり、ログ出力したり、リポジトリにコミットしないでください。
</Callout>

Step 3: アプリケーションサーバーのセットアップ [#step-3-アプリケーションサーバーのセットアップ]

アプリケーションサーバーのサンプルコードをセットアップします。

サンプルコードリポジトリをクローン [#サンプルコードリポジトリをクローン]

```bash
git clone https://github.com/mixigroup/mixi2-application-sample-go.git
cd mixi2-application-sample-go
```

<Callout type="info">
  このクイックスタートではサンプルコードリポジトリを使用しますが、SDK や API 定義など他の公開リポジトリもあります。詳しくは [GitHub リポジトリ](/resource/github)を参照してください。
</Callout>

認証情報を設定 [#認証情報を設定]

リポジトリ内の `.env.example` を `.env` にコピーし、Step 2 で生成した認証情報を設定します。

```bash
cp .env.example .env
```

`.env` ファイルを編集し、以下の環境変数を設定してください。

| 変数名              | 説明                                                 |
| ---------------- | -------------------------------------------------- |
| `CLIENT_ID`      | mixi2 Developer Platform で発行した OAuth2 クライアント ID    |
| `CLIENT_SECRET`  | mixi2 Developer Platform で発行した OAuth2 クライアントシークレット |
| `TOKEN_URL`      | mixi2 Developer Platform で確認したトークンエンドポイント URL      |
| `STREAM_ADDRESS` | mixi2 Developer Platform で確認した Stream サーバーアドレス     |

サンプルコードの詳細は[SDK ガイド](/guides/sdk#サンプルアプリケーション)を参照してください。

Step 4: アプリケーションサーバーを起動する [#step-4-アプリケーションサーバーを起動する]

アプリケーションサーバーと mixi2 のサーバーを接続する方法は、gRPC ストリーム接続と Webhook URL 登録の 2 種類があります。

ローカル環境から検証する場合は、外部からアクセス可能な URL が不要な gRPC ストリーム接続が便利です。以下のコマンドでアプリケーションサーバーを起動します。

```bash
source .env
go run cmd/stream/main.go
```

起動すると、mixi2 のサーバーに gRPC ストリーム接続を確立し、イベントの待ち受けを開始します。

Step 5: 動作を確認する [#step-5-動作を確認する]

アプリケーションサーバーが正しく動作しているか確認します。

mixi2 アプリで、Step 1 で作成したアプリケーションに DM を送信してください。
ターミナルにイベント受信のログが表示され、送信した内容と同じメッセージが返ってくれば成功です。

次のステップ [#次のステップ]

* [アプリケーション開発](/guides/application) - アプリケーションの詳細な開発方法
* [Webhook でイベントを受信する](/guides/webhook) - Webhook 方式の実装・署名検証・デプロイ
* [gRPC ストリームでイベントを受信する](/guides/grpc-stream) - gRPC 方式の実装・再接続
* [GitHub リポジトリ](/resource/github) - SDK・API 定義・サンプルコードの公開リポジトリ一覧
