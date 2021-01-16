package main
// パッケージの宣言

//ほかのパッケージを取り込むために使用
import (
       "net/http" //HTTPを扱うパッケージ
       "github.com/labstack/echo" //echoパッケージ
       "database/sql" // データベースドライバ
       //"fmt"

       _ "github.com/go-sql-driver/mysql" //database/sqlパッケージのMySQLドライバー
)

// メインの関数 基本的に最初に呼ばれる
func main () {
   e := echo.New() //Echoインスタンスを作成します｡ var e 型はわからない = echo.New()と同じ

   e.Static("/","public/")// 静的ファイル

   e.GET("/params", returnParams) // パラメーター取得

   e.GET("/tasks", getTasks) // DBへのリクエスト

   e.GET("/connect", connectCheck) // コネクトチェック
   // webAPI
   //e.GET("/tasks", getTasks)            // 一覧
   //e.GET("/tasks/:id", getTask)         // 詳細
   //e.POST("/tasks", saveTask)           // 作成
   //e.PUT("/tasks/:id", updateTask)      // 編集
   //e.DELETE("//tasks/:id", deleteTask)  // 削除

   e.Logger.Fatal(e.Start(":8082"))
   // サーバーを開始 このAPIの出力ポート e.Startの中はdocker-composeのgoコンテナで設定したportsを指定
}

func connectCheck(c echo.Context) error {
   return c.String(http.StatusOK, "こんにちは 値を入力してください")
}

func returnParams(c echo.Context) error {
   param := c.QueryParam("key")
   return c.String(http.StatusOK, "key = " + param)
}

// JSON用の構造体を定義
type Task struct {
   ID   int    `json:"id"`
   Name string `json:"name"`
}

func dbConnect() *sql.DB {
   // DB接続処理
   //sql.DB DBを司るオブジェクト ＊はポインタ変数の宣言
   db, err := sql.Open("mysql", "root:password@tcp(myapp_db)/test_todo")
   //ローカルのmysqlの指定したデータベースにアクセス出来る
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
   db := dbConnect() // データベースの設定
   defer db.Close() // コネクション解放
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