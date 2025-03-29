# syshack-backend

シスハック用バックエンド
# 仕様
APIのエンドポイントはポート`8080`番で待機します。

## エンドポイント一覧

| メソッド | パス                       | 説明                                                                 |
| :------- | :------------------------- | :------------------------------------------------------------------- |
| GET      | /                          | Hello, World!                                                        |
| GET      | /users                     | ユーザ全件取得                                                         |
| GET      | /user/:user_id            | ユーザ１件取得                                                         |
| POST     | /user                      | ユーザ登録(google_idとユーザ名ををPOSTリクエストボディのJSONに含めて送信します) |
| DELETE   | /user/:user_id            | user_idに指定されたユーザ削除                                          |
| PUT      | /user/:user_id/administrator |  管理者権限を追加、及び削除します。レスポンスボディのIs_administratorフィールドをtrueまたはfalseに設定することで変更できます。                        |
| GET      | /kakomons                  | 過去問一覧取得(取得する過去問の条件をGETパラメータで送信してください.複数の指定はできません) |
| GET      | /kakomon/:id              | 過去問指定取得(指定したidの過去問を取得します)                         |
| POST     | /kakomon                   | 過去問登録(過去問情報はjsonで送信、ファイル本体はmultipart-formdataで送信します) |
| DELETE   | /kakomon/:id              | 指定した過去問を削除します                                               |
| GET      | /butterflies/:feed_user_id | 蝶取得一覧                                                             |
| GET      | /butterfly/:id            | 蝶指定取得                                                             |
| POST     | /butterfly/:id           | 蝶登録                                                               |
| PUT      | /butterfly/:id            | 蝶更新                                                               |
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
