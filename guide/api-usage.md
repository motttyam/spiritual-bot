# API の使い方 (/docs/guides/api-usage)



このガイドでは、mixi2 API を使用してアプリケーションから各種操作を行う方法を解説します。

前提条件 [#前提条件]

* [開発者登録](/getting-started/registration)が完了していること
* [クイックスタート](/getting-started/quickstart)を完了していること

主要 API 一覧 [#主要-api-一覧]

| RPC                       | 説明                        |
| ------------------------- | ------------------------- |
| `CreatePost`              | ポストを作成（返信/引用/メディア添付対応）    |
| `SendChatMessage`         | チャットメッセージを送信（テキスト/メディア添付） |
| `InitiatePostMediaUpload` | メディアアップロードを開始             |
| `GetPostMediaStatus`      | メディアのアップロード/処理状況を取得       |
| `GetStamps`               | スタンプ一覧を取得                 |
| `AddStampToPost`          | ポストにスタンプを付与               |
| `GetUsers`                | ユーザー情報を取得                 |
| `GetPosts`                | ポスト情報を取得                  |

ポストの作成 [#ポストの作成]

`CreatePost` RPC を使用してポストを作成します。

基本的なポスト [#基本的なポスト]

```go
resp, err := client.CreatePost(ctx, &application_apiv1.CreatePostRequest{
    Text: "こんにちは！",
})
```

リプライ [#リプライ]

受信したポストに返信する場合は、`in_reply_to_post_id` を指定します。

```go
inReplyToPostId := event.Post.PostId
resp, err := client.CreatePost(ctx, &application_apiv1.CreatePostRequest{
    Text:            "返信ありがとうございます！",
    InReplyToPostId: &inReplyToPostId,
})
```

引用ポスト [#引用ポスト]

他のポストを引用する場合は、`quoted_post_id` を指定します。

```go
quotedPostId := originalPostId
resp, err := client.CreatePost(ctx, &application_apiv1.CreatePostRequest{
    Text:         "これは面白い投稿ですね！",
    QuotedPostId: &quotedPostId,
})
```

<Callout type="warning">
  `in_reply_to_post_id` と `quoted_post_id` は同時に指定できません。
</Callout>

マスク付きポスト [#マスク付きポスト]

センシティブなコンテンツやネタバレを含むポストには、マスクを適用できます。

```go
caption := "映画のネタバレ注意"
resp, err := client.CreatePost(ctx, &application_apiv1.CreatePostRequest{
    Text: "ネタバレを含む内容です...",
    PostMask: &modelv1.PostMask{
        MaskType: constv1.PostMaskType_POST_MASK_TYPE_SPOILER,
        Caption:  &caption,
    },
})
```

| マスク種別                      | 説明                |
| -------------------------- | ----------------- |
| `POST_MASK_TYPE_SENSITIVE` | 刺激的なコンテンツに対する注意喚起 |
| `POST_MASK_TYPE_SPOILER`   | ネタバレ防止のための注意喚起    |

**ユースケース:**

* センシティブな可能性があるコンテンツに予防的にマスクを適用
* ゲーム攻略・レビューボットでネタバレを含む返信にスポイラーマスクを適用
* ユーザーが「ネタバレあり」などのキーワードを含めた場合に自動でマスクを適用

配信設定 [#配信設定]

ポストの配信範囲を制御できます。

```go
publishingType := constv1.PostPublishingType_POST_PUBLISHING_TYPE_NOT_PUBLISHING
resp, err := client.CreatePost(ctx, &application_apiv1.CreatePostRequest{
    Text:           "このポストはプロフィールにのみ表示されます",
    PublishingType: &publishingType,
})
```

| 配信設定                                  | 説明                     |
| ------------------------------------- | ---------------------- |
| `POST_PUBLISHING_TYPE_UNSPECIFIED`    | フォロワーのタイムラインに公開（デフォルト） |
| `POST_PUBLISHING_TYPE_NOT_PUBLISHING` | プロフィールにのみ公開            |

<Callout type="info">
  `NOT_PUBLISHING` は、特定ユーザーへの返信時にフォロワー全体のタイムラインに流さない場合や、テスト投稿に便利です。
</Callout>

ポスト作成の制限 [#ポスト作成の制限]

| 項目        | 制限            |
| --------- | ------------- |
| テキスト最大文字数 | 149 文字        |
| メディア添付    | 最大 4 件        |
| メンション数    | 文字数に収まる限り制限なし |

DM の送信 [#dm-の送信]

`SendChatMessage` RPC を使用して DM を送信します。

```go
text := "メッセージを受け取りました！"
resp, err := client.SendChatMessage(ctx, &application_apiv1.SendChatMessageRequest{
    RoomId: event.Message.RoomId,
    Text:   &text,
})
```

<Callout type="info">
  アプリケーションから先に DM を送ることはできません。ユーザーから DM を受信した際に、イベントに含まれる `room_id` を使用して返信できます。
</Callout>

DM 送信の制限 [#dm-送信の制限]

| 項目      | 制限                          |
| ------- | --------------------------- |
| 先送り可否   | 不可（ユーザーからの DM 受信後のみ返信可能）    |
| 必須フィールド | `text` または `media_id` のいずれか |

メディアのアップロード [#メディアのアップロード]

ポストや DM にメディア（画像・動画）を添付するには、以下のフローで処理します。

Step 1: アップロードの開始 [#step-1-アップロードの開始]

`InitiatePostMediaUpload` を呼び出して、`media_id` と `upload_url` を取得します。

```go
description := "画像の説明（任意）"
resp, err := client.InitiatePostMediaUpload(ctx, &application_apiv1.InitiatePostMediaUploadRequest{
    MediaType:   application_apiv1.InitiatePostMediaUploadRequest_TYPE_IMAGE,
    ContentType: "image/jpeg",
    DataSize:    uint64(len(imageData)),
    Description: &description,
})

mediaId := resp.MediaId
uploadUrl := resp.UploadUrl
```

Step 2: メディアのアップロード [#step-2-メディアのアップロード]

取得した `upload_url` にメディアデータを POST で送信します。

```go
req, err := http.NewRequest(http.MethodPost, uploadUrl, bytes.NewReader(imageData))
req.Header.Set("Content-Type", "image/jpeg")
httpClient.Do(req)
```

Step 3: 処理状況の確認 [#step-3-処理状況の確認]

`GetPostMediaStatus` でメディアの処理状況を確認します。`STATUS_COMPLETED` になるまでポーリングしてください。

```go
for {
    statusResp, err := client.GetPostMediaStatus(ctx, &application_apiv1.GetPostMediaStatusRequest{
        MediaId: mediaId,
    })
    if statusResp.Status == application_apiv1.GetPostMediaStatusResponse_STATUS_COMPLETED {
        break
    }
    if statusResp.Status == application_apiv1.GetPostMediaStatusResponse_STATUS_FAILED {
        // エラー処理: アップロードをやり直す必要があります
        return errors.New("media processing failed")
    }
    time.Sleep(1 * time.Second)
}
```

| ステータス                   | 説明        |
| ----------------------- | --------- |
| `STATUS_UPLOAD_PENDING` | アップロード待機中 |
| `STATUS_PROCESSING`     | 処理中       |
| `STATUS_COMPLETED`      | 完了        |
| `STATUS_FAILED`         | 失敗        |

Step 4: ポストへの添付 [#step-4-ポストへの添付]

処理が完了したら、`CreatePost` で `media_id` を指定してポストを作成します。

```go
resp, err := client.CreatePost(ctx, &application_apiv1.CreatePostRequest{
    Text:        "画像を添付しました！",
    MediaIdList: []string{mediaId},
})
```

メディアアップロードの制限 [#メディアアップロードの制限]

| 項目             | 制限                     |
| -------------- | ---------------------- |
| 画像最大サイズ        | 15 MB                  |
| 動画最大サイズ        | 50 MB                  |
| 対応フォーマット       | JPEG, PNG, GIF, MP4 など |
| アップロード有効期限（画像） | 200 秒                  |
| アップロード有効期限（動画） | 600 秒                  |

<Callout type="warning">
  `STATUS_FAILED` になったメディアは再利用できません。`InitiatePostMediaUpload` からやり直してください。
</Callout>

スタンプの付与 [#スタンプの付与]

`AddStampToPost` RPC を使用して、ポストにスタンプを付与できます。

```go
// 利用可能なスタンプ一覧を取得
language := constv1.LanguageCode_LANGUAGE_CODE_JP
stampsResp, err := client.GetStamps(ctx, &application_apiv1.GetStampsRequest{
    OfficialStampLanguage: &language,
})

// スタンプを付与
_, err = client.AddStampToPost(ctx, &application_apiv1.AddStampToPostRequest{
    PostId:  postId,
    StampId: stampsResp.OfficialStampSets[0].Stamps[0].StampId,
})
```

スタンプ付与の制限 [#スタンプ付与の制限]

| 項目        | 制限                      |
| --------- | ----------------------- |
| 対象ポスト     | アプリケーションにメンションしているポストのみ |
| 使用可能なスタンプ | 公式スタンプのみ                |
| 付与回数      | 同じポストに複数回付与不可           |

<Callout type="info">
  アプリケーションが付与したスタンプを取り消す機能は現在提供されていません。
</Callout>

ユースケース [#ユースケース]

* ユーザーからのメンションに対して「確認しました」の意味でスタンプを付与
* 特定のキーワードを含むメンションにリアクションスタンプを付与
* 返信の代わりに軽量なリアクションとして使用

ユーザー情報の取得 [#ユーザー情報の取得]

`GetUsers` RPC を使用して、ユーザー情報を取得できます。

```go
resp, err := client.GetUsers(ctx, &application_apiv1.GetUsersRequest{
    UserIdList: []string{event.Post.CreatorId},
})

for _, user := range resp.Users {
    fmt.Printf("ユーザー名: %s\n", user.DisplayName)
}
```

ユースケース [#ユースケース-1]

* イベントで受信したユーザーの詳細情報（表示名、アイコン URL）を取得
* ポスト作成者の情報を取得してログ出力や管理画面に表示

<Callout type="info">
  アプリケーションがアクセス可能なユーザー情報のみ取得できます。
</Callout>

ポスト情報の取得 [#ポスト情報の取得]

`GetPosts` RPC を使用して、ポスト情報を取得できます。

```go
resp, err := client.GetPosts(ctx, &application_apiv1.GetPostsRequest{
    PostIdList: []string{event.Post.GetInReplyToPostId()},
})

for _, post := range resp.Posts {
    fmt.Printf("ポスト本文: %s\n", post.Text)
}
```

ユースケース [#ユースケース-2]

* リプライや引用ポストを受信した際に、元のポスト情報（本文、メディア、スタンプ等）を取得
* メンションされたポストの添付メディアやスタンプ情報を確認

<Callout type="info">
  アプリケーションがアクセス可能なポストのみ取得できます。
</Callout>

次のステップ [#次のステップ]

* [アプリケーション開発](/guides/application) - Webhook URL の設定、デプロイ
* [イベント](/guides/events) - イベントの種類と構造
* [Webhook でイベントを受信する](/guides/webhook) - Webhook 方式の実装
* [gRPC ストリームでイベントを受信する](/guides/grpc-stream) - gRPC 方式の実装
* [API リファレンス](/reference/api-document) - API の完全な仕様
