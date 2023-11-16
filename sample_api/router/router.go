package router

import (
	"net/http"
	"sampleApi/handler"

	"github.com/labstack/echo/v4"
)

// 新しいリクエスト先が必要な場合はここに追加する
func Route(e *echo.Echo) {
	/*
		httpメソッドは主にGET,POST,PUT,DELETEが使用されます。
		GET: DBからデータを取得する
		POST: DBのデータを更新する
		PUT: DBにデータを新規作成する
		DELTE: DBからデータを削除する

		golangのechoではe.〇〇の〇〇にGET～DELETEを入れることで
		クライアント（ブラウザからの操作）からDBに対して
		どのような処理を実行させたいか決定します。
	*/
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Accessible")
	})

	// クライアントからhttp://{バックエンドのURL}:{バックエンドのポート}/helloが
	// リクエストされたときに、handler/hello.goのHello関数を実行する
	e.GET("/hello", handler.Hello)

	// サーバーを起動して、以下のコマンドをターミナルから実行しましょう
	// curl -X PUT -H "Content-Type: application/json" -d "{\"Name\": \"smart phone\"}" localhost:1324/item
	e.PUT("/item", handler.CreateItem)

	// 課題１ handlerディレクトリにget_item.goを作成し、ブラウザに文字列"smart phone"を表示させる
	e.GET("/get_item", handler.GetItem)
	// 課題２ entityに新しく構造体を作成して、任意のテーブルを作成し、データを登録する

	// 課題３ 課題２で作成したテーブルにデータを入力し、そのデータを更新する処理を作成する
}
