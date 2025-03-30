# syshack-backend

シスハック用バックエンド
# 仕様
APIのエンドポイントはポート`8080`番で待機します。

## エンドポイント一覧

| 実装済みか | メソッド | パス                       | 説明                                                                 |
| :--------- | :------- | :------------------------- | :------------------------------------------------------------------- |
|     ⭕️      | GET      | /                          | 正常にサーバが起動すればOKを返す                                                        |
|     ⭕️     | GET      | /users                     | ユーザ全件取得(サーバからの応答例: `{"user":[{"ID":1,"CreatedAt":"2025-03-30T07:02:14.195969Z","UpdatedAt":"2025-03-30T07:02:14.195969Z","DeletedAt":null,"user_id":"0f9e68c2-72f5-47c9-b313-b85a0368c880","name":"かきくけこ","google_id":"def","previous_upload_date":"0001-01-01T00:00:00Z","is_administrator":false,"count_post":0,"feeding_butterfly_id":0},{"ID":2,"CreatedAt":"2025-03-30T07:02:29.261427Z","UpdatedAt":"2025-03-30T07:02:29.261427Z","DeletedAt":null,"user_id":"2136ff68-cebb-4468-a3f7-e3c50f5db4fa","name":"あいうえお","google_id":"abc","previous_upload_date":"0001-01-01T00:00:00Z","is_administrator":true,"count_post":0,"feeding_butterfly_id":0}]}`)                                                         |
|     ⭕️      | GET      | /user/:user_id            | ユーザ１件取得取得(サーバからの応答例:`{"ID":2,"CreatedAt":"2025-03-30T07:02:29.261427Z","UpdatedAt":"2025-03-30T07:02:29.261427Z","DeletedAt":null,"user_id":"2136ff68-cebb-4468-a3f7-e3c50f5db4fa","name":"あいうえお","google_id":"abc","previous_upload_date":"0001-01-01T00:00:00Z","is_administrator":true,"count_post":0,"feeding_butterfly_id":0}`)                                                         |
|     ⭕️     | POST     | /user                      | ユーザ登録(google_idとname(ユーザ名)をPOSTリクエストボディのJSONに含めて送信してください,リクエストボディの例:`{"google_id" : "abc","name":"あいうえお","is_administrator" : true}`) |
|            | DELETE   | /user/:user_id            | user_idに指定されたユーザ削除                                          |
|            | PUT      | /user/:user_id/administrator |  管理者権限を追加、及び削除します。レスポンスボディのIs_administratorフィールドをtrueまたはfalseに設定することで変更できます。                        |
|            | GET      | /kakomons                  | 過去問一覧取得(取得する過去問の条件をGETパラメータで送信してください.複数の指定はできません) |
|            | GET      | /kakomon/:id              | 過去問指定取得(指定したidの過去問を取得します)                         |
|     ⭕️     | POST     | /kakomon                   | 過去問登録 次のような形でリクエストを送信してください: `file` : 過去問ファイル本体、`formData` : `{"grade": "B3",   "subject": "線形代数",   "title": "中間試験",   "year": 2024,   "teacher": "山田太郎",   "major": "kk", "upload_user_id" : "アップロードしたユーザのUUIDをここに書く"}`|
|            | DELETE   | /kakomon/:id              | 指定した過去問を削除します                                               |
|            | GET      | /butterflies/:feed_user_id | 蝶取得一覧                                                             |
|            | GET      | /butterfly/:id            | 蝶指定取得                                                             |
|            | POST     | /butterfly/:id           | 蝶登録                                                               |
|            | PUT      | /butterfly/:id            | 蝶更新                                                               |


以上、ざっくりとしたAPIエンドポイントの仕様です。
**(API仕様は以上のとおりですが、レスポンス周りの仕様について完全に実装が完了していないため詳細は控えさせてただきます。何か以上の設計で問題があれば教えてください)**
# セットアップ
使用する前にDockerとDocker Composeのインストールが必要です。Linuxでは、以下のSnapパッケージをインストールすることで両方インストールすることができます。
```bash
# Linux
sudo snap install docker
```

## サーバ立ち上げ方法  
```
docker-compose up  -d --build 
```
