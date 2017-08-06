package main

import (
	. "./models"
	"./controllers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)


func main() {
	fmt.Println("Start on localhost:8080")

	//beego controller
	bController := new(controllers.MainController)

	// Объявляю счетчик, который сразу будет обрабатывать данные с каналов
	var c = NewCounter(1)

	// Запросы на корень сайта обрабатывает IndexHandler c помощью martini
	beego.Router("/", &controllers.MainController{}, "get:Index")

	//Обработчик post запросов на увеличение значения счетчика
	beego.Post("/add", func(ctx *context.Context) {
		bController.Add(c)

		ctx.Output.JSON(http.StatusOK, true, true)
	})

	//Обработчик вывода значения счетчика
	beego.Get("/show", func(ctx *context.Context) {
		to_json := bController.Show(c)

		ctx.Output.JSON(to_json, true, true)
	})

	beego.Run()
}
