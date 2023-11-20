package main

import (
	"sampleApi/db"
	"sampleApi/entity"
	"sampleApi/router"
	"sampleApi/validate"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	  }
)
func main() {
	initDB()
	e := echo.New()

	e.Use(middleware.Logger())  // サーバーログを表示する
	e.Use(middleware.Recover()) // エラーでサーバーが落ちた時にリカバリーする

	// リクエストによって適切な処理を実行する
	router.Route(e)

	e.Validator = &validate.CustomValidator{Validator: validator.New()}

	// サーバーをポート番号1324で起動
	e.Logger.Fatal(e.Start(":1324"))
}


func initDB() {
	dbInstance := db.GetDB()
	dbInstance.AutoMigrate(
		&entity.Item{},
		&entity.User{},
	)
}
