# syshack-backend

シスハック用バックエンド
# 仕様
APIのエンドポイントはポート`8080(仮)`番で待機します。詳細な仕様は後ほど決定します。
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
