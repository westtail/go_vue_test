// パッケージの宣言
package main

//ほかのパッケージを取り込むために使用
import (
       "net/http" //HTTPを扱うパッケージ
       "github.com/labstack/echo" //echoパッケージ
       //"fmt"
)

// メインの関数 基本的に最初に呼ばれる
func main () {
   e := echo.New() //Echoインスタンスを作成します｡ var e 型はわからない = echo.New()と同じ

   e.Static("/","public/")// 静的ファイル

   e.GET("/data", firstData) // GETリクエスト 別の書き方もできる

   e.GET("/params", returnParams) // パラメーター取得

   // サーバーを開始
   e.Logger.Fatal(e.Start(":8082"))
   // このAPIの出力ポート
   // e.Startの中はdocker-composeのgoコンテナで設定したportsを指定してください。
}

func firstData(c echo.Context) error {
   return c.String(http.StatusOK, "こんにちは 値を入力してください")
}


func returnParams(c echo.Context) error {
   param := c.QueryParam("key")
   return c.String(http.StatusOK, "key = " + param)
}