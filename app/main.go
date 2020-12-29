package main

//echoのデータを取得
import (
       "net/http"
       "github.com/labstack/echo"
)

// メインの関数
func main () {
     e := echo.New() //Echoインスタンスを作成します｡
     e.Static("/","public/")// 静的ファイル
     e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
     })
     e.Logger.Fatal(e.Start(":8082"))
     // e.Startの中はdocker-composeのgoコンテナで設定したportsを指定してください。
}