# API リファレンス (/docs/reference/api-document)



共通仕様 [#共通仕様]

プロトコル [#プロトコル]

gRPC を使用します。

Protocol Buffers 定義は以下のリポジトリで公開しています。

* [mixi2-api](https://github.com/mixigroup/mixi2-api)

認証 [#認証]

OAuth 2.0 による認証が必要です。認証情報の取得方法は[クイックスタート](/getting-started/quickstart)を参照してください。

RPC一覧 [#rpc一覧]

GetUsers [#getusers]

指定したユーザーIDリストに対応するユーザー情報を取得します。

GetUsersRequest [#getusersrequest]

ユーザー情報取得リクエストです。

| Field          | Type            | Description           |
| -------------- | --------------- | --------------------- |
| user\_id\_list | repeated string | 取得対象のユーザーIDを指定してください。 |

GetUsersResponse [#getusersresponse]

ユーザー情報取得レスポンスです。

| Field | Type          | Description  |
| ----- | ------------- | ------------ |
| users | repeated User | ユーザー情報の一覧です。 |

GetPosts [#getposts]

指定したポストIDリストに対応するポスト情報を取得します。

GetPostsRequest [#getpostsrequest]

ポスト情報取得リクエストです。

| Field          | Type            | Description          |
| -------------- | --------------- | -------------------- |
| post\_id\_list | repeated string | 取得対象のポストIDを指定してください。 |

GetPostsResponse [#getpostsresponse]

ポスト情報取得レスポンスです。

| Field | Type          | Description |
| ----- | ------------- | ----------- |
| posts | repeated Post | ポスト情報の一覧です。 |

CreatePost [#createpost]

ポストを作成します（返信/引用/メディア添付等に対応）。

CreatePostRequest [#createpostrequest]

in\_reply\_to\_post\_id と quoted\_post\_id は同時に指定できません。

| Field                   | Type                        | Description                  |
| ----------------------- | --------------------------- | ---------------------------- |
| text                    | string                      | ポストの本文を指定してください。             |
| in\_reply\_to\_post\_id | optional string             | 返信先ポストIDを指定してください（任意）。       |
| quoted\_post\_id        | optional string             | 引用対象ポストIDを指定してください（任意）。      |
| media\_id\_list         | repeated string             | 添付するメディアID一覧を指定してください（最大4件）。 |
| post\_mask              | optional PostMask           | ポストに適用するマスクを指定してください（任意）。    |
| publishing\_type        | optional PostPublishingType | ポストの配信設定を指定してください。           |

CreatePostResponse [#createpostresponse]

ポスト作成レスポンスです。

| Field | Type | Description   |
| ----- | ---- | ------------- |
| post  | Post | 作成されたポスト情報です。 |

InitiatePostMediaUpload [#initiatepostmediaupload]

ポストやメッセージ（ルーム送信/DM）に添付するメディアのアップロードを開始し、アップロード先URLを発行します。

InitiatePostMediaUploadRequest [#initiatepostmediauploadrequest]

メディアアップロード開始リクエストです。

| Field         | Type            | Description                        |
| ------------- | --------------- | ---------------------------------- |
| content\_type | string          | アップロードするデータのContent-Typeを指定してください。 |
| data\_size    | uint64          | アップロードするデータサイズ（バイト）を指定してください。      |
| media\_type   | Type            | メディア種別を指定してください。                   |
| description   | optional string | メディアの説明を指定してください（任意）。              |

InitiatePostMediaUploadResponse [#initiatepostmediauploadresponse]

メディアアップロード開始レスポンスです。

| Field       | Type   | Description                                 |
| ----------- | ------ | ------------------------------------------- |
| media\_id   | string | アップロード状況確認や、ポスト/メッセージに送信時にメディアを添付するためのIDです。 |
| upload\_url | string | メディアデータをアップロードするためのURLです。                   |

GetPostMediaStatus [#getpostmediastatus]

指定したメディアIDのアップロード/処理状況を取得します。

GetPostMediaStatusRequest [#getpostmediastatusrequest]

メディアアップロード状況取得リクエストです。

| Field     | Type   | Description                      |
| --------- | ------ | -------------------------------- |
| media\_id | string | アップロード状況を確認する対象のメディアIDを指定してください。 |

GetPostMediaStatusResponse [#getpostmediastatusresponse]

メディアアップロード状況取得レスポンスです。

| Field  | Type   | Description         |
| ------ | ------ | ------------------- |
| status | Status | メディアのアップロード/処理状況です。 |

SendChatMessage [#sendchatmessage]

指定したルームにチャットメッセージを送信します（テキスト/メディア添付）。

SendChatMessageRequest [#sendchatmessagerequest]

text または media\_id のいずれかは必須です。

| Field     | Type            | Description              |
| --------- | --------------- | ------------------------ |
| room\_id  | string          | 送信先ルームIDを指定してください。       |
| text      | optional string | 送信するテキストを指定してください（任意）。   |
| media\_id | optional string | 添付するメディアIDを指定してください（任意）。 |

SendChatMessageResponse [#sendchatmessageresponse]

チャットメッセージ送信レスポンスです。

| Field   | Type        | Description   |
| ------- | ----------- | ------------- |
| message | ChatMessage | 送信されたメッセージです。 |

GetStamps [#getstamps]

スタンプ一覧を取得します。

GetStampsRequest [#getstampsrequest]

スタンプ一覧取得リクエストです。

| Field                     | Type                  | Description                                             |
| ------------------------- | --------------------- | ------------------------------------------------------- |
| official\_stamp\_language | optional LanguageCode | 取得する公式スタンプの言語を指定してください（任意）。<br>未指定の場合、公式スタンプ一覧は空で返されます。 |

GetStampsResponse [#getstampsresponse]

スタンプ一覧取得レスポンスです。

| Field                 | Type                      | Description         |
| --------------------- | ------------------------- | ------------------- |
| official\_stamp\_sets | repeated OfficialStampSet | 指定言語の公式スタンプセット一覧です。 |

AddStampToPost [#addstamptopost]

指定したポストにスタンプを付与します。

AddStampToPostRequest [#addstamptopostrequest]

ポストへのスタンプ付与リクエストです。

| Field     | Type   | Description                                                                |
| --------- | ------ | -------------------------------------------------------------------------- |
| post\_id  | string | スタンプを付与する対象のポストIDを指定してください。指定可能なポストIDは次の通りです。<br>• アプリケーションにメンションしているポストID |
| stamp\_id | string | 付与するスタンプIDを指定してください。指定可能なスタンプIDは次の通りです。<br>• 公式スタンプID                      |

AddStampToPostResponse [#addstamptopostresponse]

ポストへのスタンプ付与レスポンスです。

| Field | Type | Description |
| ----- | ---- | ----------- |
| post  | Post | 更新されたポストです。 |

SubscribeEvents [#subscribeevents]

イベントをストリーミングで購読します。

SubscribeEventsRequest [#subscribeeventsrequest]

イベント購読リクエストです。

*(no fields)*

SubscribeEventsResponse [#subscribeeventsresponse]

イベント購読レスポンスです。

| Field  | Type           | Description    |
| ------ | -------------- | -------------- |
| events | repeated Event | 受信したイベントの情報です。 |

メッセージ型 [#メッセージ型]

ChatMessageReceivedEvent [#chatmessagereceivedevent]

チャットメッセージを受信したことを通知するイベントです。

| field\_name         | field\_type          | field\_description |
| ------------------- | -------------------- | ------------------ |
| event\_reason\_list | repeated EventReason | イベントが発生した理由を示します。  |
| message             | ChatMessage          | 受信したメッセージの情報です。    |
| issuer              | User                 | メッセージを送信したユーザーです。  |

Event [#event]

アプリケーションが受信するイベントを表します。

| field\_name                    | field\_type                           | field\_description                                             |
| ------------------------------ | ------------------------------------- | -------------------------------------------------------------- |
| event\_id                      | string                                | イベントIDです。                                                      |
| event\_type                    | EventType                             | イベントの種別です。                                                     |
| ping\_event                    | oneof (body) PingEvent                | event\_type が EVENT\_TYPE\_PING の場合に設定されます。                    |
| post\_created\_event           | oneof (body) PostCreatedEvent         | event\_type が EVENT\_TYPE\_POST\_CREATED の場合に設定されます。           |
| chat\_message\_received\_event | oneof (body) ChatMessageReceivedEvent | event\_type が EVENT\_TYPE\_CHAT\_MESSAGE\_RECEIVED の場合に設定されます。 |

PingEvent [#pingevent]

疎通確認用のイベントです。

*(no fields)*

PostCreatedEvent [#postcreatedevent]

ポストが作成されたことを通知するイベントです。

| field\_name         | field\_type          | field\_description |
| ------------------- | -------------------- | ------------------ |
| event\_reason\_list | repeated EventReason | イベントが発生した理由を示します。  |
| post                | Post                 | 作成されたポストの情報です。     |
| issuer              | User                 | ポストしたユーザーの情報です。    |

Media [#media]

メディアを表します。

| field\_name | field\_type                | field\_description                           |
| ----------- | -------------------------- | -------------------------------------------- |
| media\_type | MediaType                  | メディアの種別です。                                   |
| image       | oneof (content) MediaImage | media\_type が MEDIA\_TYPE\_IMAGE の場合に設定されます。 |
| video       | oneof (content) MediaVideo | media\_type が MEDIA\_TYPE\_VIDEO の場合に設定されます。 |

MediaImage [#mediaimage]

画像の情報を表します。

| field\_name              | field\_type | field\_description    |
| ------------------------ | ----------- | --------------------- |
| large\_image\_url        | string      | 大きいサイズの画像のURLです。      |
| large\_image\_mime\_type | string      | 大きいサイズの画像のMIMEタイプです。  |
| large\_image\_height     | uint32      | 大きいサイズの画像の高さ（ピクセル）です。 |
| large\_image\_width      | uint32      | 大きいサイズの画像の幅（ピクセル）です。  |
| small\_image\_url        | string      | 小さいサイズの画像のURLです。      |
| small\_image\_mime\_type | string      | 小さいサイズの画像のMIMEタイプです。  |
| small\_image\_height     | uint32      | 小さいサイズの画像の高さ（ピクセル）です。 |
| small\_image\_width      | uint32      | 小さいサイズの画像の幅（ピクセル）です。  |

MediaStamp [#mediastamp]

スタンプ画像の情報を表します。

| field\_name | field\_type | field\_description |
| ----------- | ----------- | ------------------ |
| url         | string      | スタンプ画像のURLです。      |
| mime\_type  | string      | スタンプ画像のMIMEタイプです。  |
| height      | uint32      | スタンプ画像の高さ（ピクセル）です。 |
| width       | uint32      | スタンプ画像の幅（ピクセル）です。  |

MediaVideo [#mediavideo]

動画の情報を表します。

| field\_name                | field\_type | field\_description     |
| -------------------------- | ----------- | ---------------------- |
| video\_url                 | string      | 動画のURLです。              |
| video\_mime\_type          | string      | 動画のMIMEタイプです。          |
| video\_height              | uint32      | 動画の高さ（ピクセル）です。         |
| video\_width               | uint32      | 動画の幅（ピクセル）です。          |
| preview\_image\_url        | string      | 動画のプレビュー画像のURLです。      |
| preview\_image\_mime\_type | string      | 動画のプレビュー画像のMIMEタイプです。  |
| preview\_image\_height     | uint32      | 動画のプレビュー画像の高さ（ピクセル）です。 |
| preview\_image\_width      | uint32      | 動画のプレビュー画像の幅（ピクセル）です。  |
| duration                   | float       | 動画の再生時間（秒）です。          |

ChatMessage [#chatmessage]

チャットメッセージを表します。

| field\_name | field\_type     | field\_description    |
| ----------- | --------------- | --------------------- |
| room\_id    | string          | メッセージが送信されたルームのIDです。  |
| message\_id | string          | メッセージIDです。            |
| creator\_id | string          | メッセージ送信者のユーザーIDです。    |
| text        | string          | メッセージのテキストです。         |
| created\_at | Timestamp       | メッセージ送信日時です。          |
| media\_list | repeated Media  | メッセージに添付されたメディア一覧です。  |
| post\_id    | optional string | メッセージに引用されているポストIDです。 |

Post [#post]

ポストを表します。

| field\_name             | field\_type        | field\_description                                            |
| ----------------------- | ------------------ | ------------------------------------------------------------- |
| post\_id                | string             | ポストIDです。                                                      |
| is\_deleted             | bool               | ポストが削除されているかどうかを示します。削除されている場合、post\_id 以外のフィールドはデフォルト値を返します。 |
| creator\_id             | string             | ポスト作成者のユーザーIDです。                                              |
| text                    | string             | ポストの本文です。                                                     |
| created\_at             | Timestamp          | ポスト作成日時です。                                                    |
| post\_media\_list       | repeated PostMedia | ポストに添付されたメディア一覧です。                                            |
| in\_reply\_to\_post\_id | optional string    | 返信先のポストIDです。                                                  |
| post\_mask              | optional PostMask  | ポストに適用されるマスク情報です。                                             |
| visibility              | PostVisibility     | ポストを閲覧可能かどうかを示します。                                            |
| access\_level           | PostAccessLevel    | ポストの公開設定を示します。                                                |
| stamps                  | repeated PostStamp | ポストに付与されたスタンプの一覧です。                                           |
| reader\_stamp\_id       | optional string    | 現在のアプリケーションがすでにこのポストに付与したスタンプIDです。                            |

PostMask [#postmask]

ポストに適用されるマスク情報を表します。

| field\_name | field\_type  | field\_description |
| ----------- | ------------ | ------------------ |
| mask\_type  | PostMaskType | マスクのタイプです。         |
| caption     | string       | マスクのキャプションです。      |

PostMedia [#postmedia]

ポストに添付されたメディアを表します。

| field\_name | field\_type                    | field\_description                                 |
| ----------- | ------------------------------ | -------------------------------------------------- |
| media\_type | PostMediaType                  | メディアの種別です。                                         |
| image       | oneof (content) PostMediaImage | media\_type が POST\_MEDIA\_TYPE\_IMAGE の場合に設定されます。 |
| video       | oneof (content) PostMediaVideo | media\_type が POST\_MEDIA\_TYPE\_VIDEO の場合に設定されます。 |

PostMediaImage [#postmediaimage]

ポストに添付された画像の情報を表します。

| field\_name              | field\_type | field\_description    |
| ------------------------ | ----------- | --------------------- |
| large\_image\_url        | string      | 大きいサイズの画像のURLです。      |
| large\_image\_mime\_type | string      | 大きいサイズの画像のMIMEタイプです。  |
| large\_image\_height     | uint32      | 大きいサイズの画像の高さ（ピクセル）です。 |
| large\_image\_width      | uint32      | 大きいサイズの画像の幅（ピクセル）です。  |
| small\_image\_url        | string      | 小さいサイズの画像のURLです。      |
| small\_image\_mime\_type | string      | 小さいサイズの画像のMIMEタイプです。  |
| small\_image\_height     | uint32      | 小さいサイズの画像の高さ（ピクセル）です。 |
| small\_image\_width      | uint32      | 小さいサイズの画像の幅（ピクセル）です。  |

PostMediaVideo [#postmediavideo]

ポストに添付された動画の情報を表します。

| field\_name                | field\_type | field\_description     |
| -------------------------- | ----------- | ---------------------- |
| video\_url                 | string      | 動画のURLです。              |
| video\_mime\_type          | string      | 動画のMIMEタイプです。          |
| video\_height              | uint32      | 動画の高さ（ピクセル）です。         |
| video\_width               | uint32      | 動画の幅（ピクセル）です。          |
| preview\_image\_url        | string      | 動画のプレビュー画像のURLです。      |
| preview\_image\_mime\_type | string      | 動画のプレビュー画像のMIMEタイプです。  |
| preview\_image\_height     | uint32      | 動画のプレビュー画像の高さ（ピクセル）です。 |
| preview\_image\_width      | uint32      | 動画のプレビュー画像の幅（ピクセル）です。  |
| duration                   | float       | 動画の再生時間（秒）です。          |

PostStamp [#poststamp]

ポストに付与されたスタンプを表します。

| field\_name | field\_type | field\_description |
| ----------- | ----------- | ------------------ |
| stamp       | MediaStamp  | スタンプの情報です。         |
| count       | uint64      | スタンプが押された回数です。     |

OfficialStamp [#officialstamp]

公式スタンプを表します。

| field\_name  | field\_type     | field\_description      |
| ------------ | --------------- | ----------------------- |
| stamp\_id    | string          | スタンプIDです。               |
| index        | uint32          | スタンプセット（スプライト）内での並び順です。 |
| search\_tags | repeated string | スタンプの検索用タグの一覧です。        |
| url          | string          | スタンプの画像のURLです。          |

OfficialStampSet [#officialstampset]

公式スタンプセットを表します。

| field\_name      | field\_type            | field\_description                           |
| ---------------- | ---------------------- | -------------------------------------------- |
| name             | string                 | スタンプセットの名前です。                                |
| sprite\_url      | string                 | スタンプセットのスプライト画像のURLです。                       |
| stamps           | repeated OfficialStamp | スタンプセットに含まれるスタンプ一覧です。                        |
| stamp\_set\_id   | string                 | スタンプセットIDです。                                 |
| start\_at        | optional Timestamp     | スタンプセットが利用可能になる開始日時です。未指定の場合、開始日時は限定されません。   |
| end\_at          | optional Timestamp     | スタンプセットが利用可能でなくなる終了日時です。未指定の場合、終了日時は限定されません。 |
| stamp\_set\_type | StampSetType           | スタンプセットのタイプです。                               |

User [#user]

ユーザーを表します。

| field\_name   | field\_type     | field\_description                         |
| ------------- | --------------- | ------------------------------------------ |
| user\_id      | string          | ユーザーIDです。                                  |
| is\_disabled  | bool            | ユーザーが無効化されているかどうかを示します（無効化は退会やBANなどを含みます）。 |
| name          | string          | ユーザーの名前です。                                 |
| display\_name | string          | ユーザーの表示名です。                                |
| profile       | string          | ユーザーのプロフィールです。                             |
| user\_avatar  | UserAvatar      | ユーザーのアバター情報です。                             |
| visibility    | UserVisibility  | ユーザーの情報を閲覧可能かどうかを示します。                     |
| access\_level | UserAccessLevel | ユーザーの公開設定を示します。                            |

UserAvatar [#useravatar]

ユーザーのアバター画像の情報を表します。

| field\_name              | field\_type | field\_description        |
| ------------------------ | ----------- | ------------------------- |
| large\_image\_url        | string      | 大きいサイズのアバター画像のURLです。      |
| large\_image\_mime\_type | string      | 大きいサイズのアバター画像のMIMEタイプです。  |
| large\_image\_height     | uint32      | 大きいサイズのアバター画像の高さ（ピクセル）です。 |
| large\_image\_width      | uint32      | 大きいサイズのアバター画像の幅（ピクセル）です。  |
| small\_image\_url        | string      | 小さいサイズのアバター画像のURLです。      |
| small\_image\_mime\_type | string      | 小さいサイズのアバター画像のMIMEタイプです。  |
| small\_image\_height     | uint32      | 小さいサイズのアバター画像の高さ（ピクセル）です。 |
| small\_image\_width      | uint32      | 小さいサイズのアバター画像の幅（ピクセル）です。  |

列挙型 [#列挙型]

EventReason [#eventreason]

イベントの発生理由を示す列挙型

| field                                    | type | description          |
| ---------------------------------------- | ---- | -------------------- |
| EVENT\_REASON\_UNSPECIFIED               | 0    | 未指定                  |
| EVENT\_REASON\_PING                      | 1    | 接続確認                 |
| EVENT\_REASON\_POST\_REPLY               | 2    | ポストに返信された            |
| EVENT\_REASON\_POST\_MENTIONED           | 3    | ポストでメンションされた         |
| EVENT\_REASON\_POST\_QUOTED              | 4    | ポストが引用された            |
| EVENT\_REASON\_DIRECT\_MESSAGE\_RECEIVED | 8    | チャット/ダイレクトメッセージを受信した |

EventType [#eventtype]

イベントの種別を示す列挙型

| field                                | type | description              |
| ------------------------------------ | ---- | ------------------------ |
| EVENT\_TYPE\_UNSPECIFIED             | 0    | 未指定                      |
| EVENT\_TYPE\_PING                    | 1    | 接続確認                     |
| EVENT\_TYPE\_POST\_CREATED           | 2    | ポスト作成                    |
| EVENT\_TYPE\_CHAT\_MESSAGE\_RECEIVED | 4    | メッセージ受信（チャット/ダイレクトメッセージ） |

LanguageCode [#languagecode]

言語コードを示す列挙型

| field                       | type | description |
| --------------------------- | ---- | ----------- |
| LANGUAGE\_CODE\_UNSPECIFIED | 0    | 未指定         |
| LANGUAGE\_CODE\_JP          | 1    | 日本語         |
| LANGUAGE\_CODE\_EN          | 2    | 英語          |

MediaType [#mediatype]

メッセージに添付されるメディア種別を示す列挙型

| field                    | type | description |
| ------------------------ | ---- | ----------- |
| MEDIA\_TYPE\_UNSPECIFIED | 0    | 未指定         |
| MEDIA\_TYPE\_IMAGE       | 1    | 画像          |
| MEDIA\_TYPE\_VIDEO       | 2    | 動画          |

PostAccessLevel [#postaccesslevel]

ポストの公開設定を示す列挙型

| field                            | type | description        |
| -------------------------------- | ---- | ------------------ |
| POST\_ACCESS\_LEVEL\_UNSPECIFIED | 0    | 未指定                |
| POST\_ACCESS\_LEVEL\_PUBLIC      | 1    | 公開                 |
| POST\_ACCESS\_LEVEL\_PRIVATE     | 2    | 非公開（特定のユーザーのみ閲覧可能） |

PostMaskType [#postmasktype]

ポストに適用するマスク種別を示す列挙型

| field                         | type | description       |
| ----------------------------- | ---- | ----------------- |
| POST\_MASK\_TYPE\_UNSPECIFIED | 0    | 未指定               |
| POST\_MASK\_TYPE\_SENSITIVE   | 1    | 刺激的なコンテンツに対する注意喚起 |
| POST\_MASK\_TYPE\_SPOILER     | 2    | ネタバレ防止のための注意喚起    |

PostMediaType [#postmediatype]

ポストに添付されるメディア種別を示す列挙型

| field                          | type | description |
| ------------------------------ | ---- | ----------- |
| POST\_MEDIA\_TYPE\_UNSPECIFIED | 0    | 未指定         |
| POST\_MEDIA\_TYPE\_IMAGE       | 1    | 画像          |
| POST\_MEDIA\_TYPE\_VIDEO       | 2    | 動画          |

PostPublishingType [#postpublishingtype]

ポストの投稿先設定を示す列挙型

| field                                   | type | description             |
| --------------------------------------- | ---- | ----------------------- |
| POST\_PUBLISHING\_TYPE\_UNSPECIFIED     | 0    | 未指定（自分のフォロワーのタイムラインに公開） |
| POST\_PUBLISHING\_TYPE\_NOT\_PUBLISHING | 1    | ポストを自分のプロフィールにのみ公開      |

PostVisibility [#postvisibility]

ポストを閲覧できるかどうかを示す列挙型

| field                         | type | description |
| ----------------------------- | ---- | ----------- |
| POST\_VISIBILITY\_UNSPECIFIED | 0    | 未指定         |
| POST\_VISIBILITY\_VISIBLE     | 1    | ポストを閲覧できる   |
| POST\_VISIBILITY\_INVISIBLE   | 2    | ポストを閲覧できない  |

StampSetType [#stampsettype]

公式スタンプセットの種別を示す列挙型

| field                         | type | description   |
| ----------------------------- | ---- | ------------- |
| STAMP\_SET\_TYPE\_UNSPECIFIED | 0    | 未指定           |
| STAMP\_SET\_TYPE\_DEFAULT     | 1    | デフォルトのスタンプセット |
| STAMP\_SET\_TYPE\_SEASONAL    | 2    | 季節限定のスタンプセット  |

UserAccessLevel [#useraccesslevel]

ユーザーの公開設定を示す列挙型

| field                            | type | description |
| -------------------------------- | ---- | ----------- |
| USER\_ACCESS\_LEVEL\_UNSPECIFIED | 0    | 未指定         |
| USER\_ACCESS\_LEVEL\_PUBLIC      | 1    | 公開ユーザー      |
| USER\_ACCESS\_LEVEL\_PRIVATE     | 2    | 非公開ユーザー     |

UserVisibility [#uservisibility]

ユーザーを閲覧できるか示す列挙型

| field                         | type | description |
| ----------------------------- | ---- | ----------- |
| USER\_VISIBILITY\_UNSPECIFIED | 0    | 未指定         |
| USER\_VISIBILITY\_VISIBLE     | 1    | ユーザーを閲覧できる  |
| USER\_VISIBILITY\_INVISIBLE   | 2    | ユーザーを閲覧できない |

GetPostMediaStatusResponse.Status [#getpostmediastatusresponsestatus]

メディアのアップロード/処理状況を表します。

| field                   | type | description |
| ----------------------- | ---- | ----------- |
| STATUS\_UNSPECIFIED     | 0    | 未指定         |
| STATUS\_UPLOAD\_PENDING | 1    | アップロード待機中   |
| STATUS\_PROCESSING      | 2    | 処理中         |
| STATUS\_COMPLETED       | 3    | 完了          |
| STATUS\_FAILED          | 4    | 失敗          |

InitiatePostMediaUploadRequest.Type [#initiatepostmediauploadrequesttype]

アップロードするメディアの種別を指定してください。

| field             | type | description |
| ----------------- | ---- | ----------- |
| TYPE\_UNSPECIFIED | 0    | 未指定         |
| TYPE\_IMAGE       | 1    | 画像          |
| TYPE\_VIDEO       | 2    | 動画          |
