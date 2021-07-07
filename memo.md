# コマンドメモ
```
curl -s http://localhost/api/tasks
```

# goについて
## 参考URL
* 言語の基本
    * https://qiita.com/megu_ma/items/6fad257345af2d748dbe
* 文法
    * https://qiita.com/megu_ma/items/f5e66293b2934e117731
    * https://qiita.com/k-penguin-sato/items/1d0e1c6b4bf937996cd3
* net/http goでhttpを扱うようにするパッケージ
    * https://www.yoheim.net/blog.php?q=20170403
## fresh
* goのライブラリでソースコードを変更後に自動コンパイルする
* runner-buildはfreshのふぁいる？
* http://tsujitaku50.hatenablog.com/entry/2017/10/30/232659

## 文法
変数宣言と初期化を関数の内部で行う場合は、var と型宣言を省略し、 := という記号を用いることができます。

```
var message string = "hello world"
message := "hello world"
上下は一緒のこと
```

## echoについて
* 公式ドキュメント
    * https://echo.labstack.com/guide/routing
    * https://echo.labstack.com/guide
### echoのテストについて
* https://qiita.com/kiyc/items/c20ac7bb6997c0753314
### assets
* https://qiita.com/codehex/items/a6cb72251b95c8f5baa1

## restfulAPI
* https://dev.classmethod.jp/articles/web-api_rest-api_url-design/
# mysqlについて

## 基本文法
* https://qiita.com/knife0125/items/bb095a85d1a5d3c8f706
* https://qiita.com/knife0125/items/bb095a85d1a5d3c8f706#%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB%E3%81%AB%E3%83%87%E3%83%BC%E3%82%BF%E3%82%92%E6%8C%BF%E5%85%A5insert
## "database/sql"について
* SQL に関する汎用的な機能を提供してます。 コネクションの管理や、クエリの発行、トランザクションなど
* https://blog.p1ass.com/posts/go-database-sql-wrapper/
## go-sql-driver/mysqlについて
* https://noumenon-th.net/programming/2019/09/20/go-sql-driver/ 
# Vueについて
* 公式ドキュメント
    * https://jp.vuejs.org/v2/guide/
* todoアプリ
    * https://qiita.com/sayama0402/items/1ac51dfcc01144c5e874
## 非同期処理
### axios
* ブラウザやnode.js上で動くPromiseベースのHTTPクライアントです
* 非同期にHTTP通信を行いたいときに容易に実装できます。 Vue.jsでは非同期通信を行うのにaxiosを使うのがスタンダード
* https://github.com/axios/axios
* https://qiita.com/ksh-fthr/items/2daaaf3a15c4c11956e9
* https://www.willstyle.co.jp/blog/2751/

* axiosのパラメータ指定方法まとめ
    * https://qiita.com/taroc/items/f22f7dd5d6d5443c72a4
## 何でもメモ

# 参考URL

* Goで作成したAPIサーバをdockerで複数コンテナ運用
    * https://qiita.com/yuukiyuuki327/items/9d07a04c87316d3fd23c
* docker-compose公式ドキュメント
    * https://docs.docker.jp/compose/compose-file.html#command
* docer公式ドキュメント
    * https://docs.docker.jp/compose/compose-file.html#depends-ons

# メモURL
* Go(Echo) + Vue.js + nginx の環境をDocker Composeで立てる。
    * https://qiita.com/69incat/items/9bbfbf8566535dc266c6
* Vue.jsの書き方実例集(随時追加)※逆引きリファレンス的な
    * https://qiita.com/Yorinton/items/a0144c34e4edb0777493
* 初心者でもDockerでAmazonLinux+Go(Echo)+Mysqlの環境構築しつつ簡単なAPIを作って動かす
    * https://qiita.com/sola-msr/items/eb3ae3d3cbd608b06cfe#docker-composeyml
* Goで作成したAPIサーバをdockerで複数コンテナ運用
    * https://qiita.com/yuukiyuuki327/items/9d07a04c87316d3fd23c#docker-composeyml%E3%81%AE%E6%9B%B8%E3%81%8D%E6%96%B9
* 【Go】echoを使ってさくっとAPIサーバを構築する
    * https://qiita.com/yagi_eng/items/b06722dbd7a5652ec239
* Go/echo 入門
    * https://qiita.com/pylor1n/items/36912a47c893ea5782cc
* golang echo 超入門
    * https://hawksnowlog.blogspot.com/2019/04/tutorial-golang-echo.html

* mySQLの連携で参考
* https://qiita.com/sola-msr/items/eb3ae3d3cbd608b06cfe#docker-composeyml
* https://qiita.com/Le0tk0k/items/c2945c260f28f7ee2d47#docker-composeyml