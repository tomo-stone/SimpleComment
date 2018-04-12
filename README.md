# SimpleComment
簡単なコメントページのWebアプリケーションサーバ

### サーバの立て方
1. `config/appConfig.go`にMySQLのユーザ・パスワード・データベース名を入力する
1. 使用するデータベースに、以下のようにしてテーブルを作成する
  `create table test1 (name varchar(32),body text);`
1. MySQLサービスが起動していることを確認する
1. `go run`コマンドで実行するか、ビルドして実行する
