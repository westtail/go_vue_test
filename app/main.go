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

   e.GET("/", func(c echo.Context) error {
      return c.String(http.StatusOK, "Hello, World!")// 
   }) // GETリクエスト 例

   e.GET("/test", getTest) // GETリクエスト 別の書き方もできる
   e.GET("/test/:path", testParams)

   e.GET("/hoge", hogeHandler) // パラメーター取得

   e.Logger.Fatal(e.Start(":8082"))
   // このAPIの出力ポート
   // e.Startの中はdocker-composeのgoコンテナで設定したportsを指定してください。
}

func getTest(c echo.Context) error {
   return c.String(http.StatusOK, "Hello, World! test")
}

func testParams(c echo.Context) error {
   return c.String(http.StatusOK, c.Param("path"))
}

func hogeHandler(c echo.Context) error {
   return c.String(http.StatusOK, c.QueryParam("key"))
}