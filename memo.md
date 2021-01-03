## fresh
* goのライブラリでソースコードを変更後に自動コンパイルする
* runner-buildはfreshのふぁいる？
* http://tsujitaku50.hatenablog.com/entry/2017/10/30/232659

## go
* 言語の基本
    * https://qiita.com/megu_ma/items/6fad257345af2d748dbe
* 文法
    * https://qiita.com/megu_ma/items/f5e66293b2934e117731
* net/http goでhttpを扱うようにするパッケージ
* https://www.yoheim.net/blog.php?q=20170403

### 文法
変数宣言と初期化を関数の内部で行う場合は、var と型宣言を省略し、 := という記号を用いることができます。

```
var message string = "hello world"
message := "hello world"
```

## 何でもメモ

### echoについて
* 公式ドキュメント
    * https://echo.labstack.com/guide/routing
### assets
* https://qiita.com/codehex/items/a6cb72251b95c8f5baa1
### axios
* ブラウザやnode.js上で動くPromiseベースのHTTPクライアントです
* 非同期にHTTP通信を行いたいときに容易に実装できます。 Vue.jsでは非同期通信を行うのにaxiosを使うのがスタンダード
* https://github.com/axios/axios
* https://qiita.com/ksh-fthr/items/2daaaf3a15c4c11956e9
* https://www.willstyle.co.jp/blog/2751/
# 参考URL

* Goで作成したAPIサーバをdockerで複数コンテナ運用
    * https://qiita.com/yuukiyuuki327/items/9d07a04c87316d3fd23c
* docker-compose公式ドキュメント
    * https://docs.docker.jp/compose/compose-file.html#command
* docer公式ドキュメント
    * https://docs.docker.jp/compose/compose-file.html#depends-ons

* メモURL
    * https://qiita.com/69incat/items/9bbfbf8566535dc266c6
    * https://qiita.com/Yorinton/items/a0144c34e4edb0777493
    * https://qiita.com/sola-msr/items/eb3ae3d3cbd608b06cfe#docker-composeyml
    * https://qiita.com/yuukiyuuki327/items/9d07a04c87316d3fd23c#docker-composeyml%E3%81%AE%E6%9B%B8%E3%81%8D%E6%96%B9
    * https://qiita.com/yagi_eng/items/b06722dbd7a5652ec239
    * https://qiita.com/pylor1n/items/36912a47c893ea5782cc

    * https://hawksnowlog.blogspot.com/2019/04/tutorial-golang-echo.html