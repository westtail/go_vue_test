// パッケージの宣言
package main

//ほかのパッケージを取り込むために使用
import (
       "net/http" //HTTPを扱うパッケージ
       "github.com/labstack/echo" //echoパッケージ
       "database/sql" // データベースドライバ
       //"fmt"

       _ "github.com/go-sql-driver/mysql" //database/sqlパッケージのMySQLドライバー
)


// JSON用の構造体を定義
type Task struct {
   ID   int    `json:"id"`
   Name string `json:"name"`
}

// メインの関数 基本的に最初に呼ばれる
func main () {
   e := echo.New() //Echoインスタンスを作成します｡ var e 型はわからない = echo.New()と同じ

   e.Static("/","public/")// 静的ファイル

   e.GET("/data", firstData) // GETリクエスト 別の書き方もできる

   e.GET("/params", returnParams) // パラメーター取得

   e.GET("/tasks", getTasks) // DBへのリクエスト

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

func dbConnect() *sql.DB {
   // DB接続処理
   db, err := sql.Open("mysql", "root:password@tcp(myapp_db)/test_todo")
   if err != nil {
       panic(err.Error())
   }
   return db
}

func getRows(db *sql.DB) *sql.Rows {
   rows, err := db.Query("SELECT id,name FROM tasks")
   if err != nil {
       panic(err.Error())
   }
   return rows
}

func getTasks(c echo.Context) error {
   // DB接続
   db := dbConnect()
   defer db.Close()
   rows := getRows(db)
   defer rows.Close()

   task := Task{}
   var results []Task
   for rows.Next() {
       err := rows.Scan(&task.ID, &task.Name)
       if err != nil {
           panic(err.Error())
       } else {
           results = append(results, task)
       }
   }

   return c.JSON(http.StatusOK, results)
}