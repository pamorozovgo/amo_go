package main

import (
	"fmt"
	"./handlers"
	"encoding/json"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"net/http"

	"sync"
)


/**
 * Наша структура счетчика, содержащая само значение счетчика и шаг прибавления
 */
type Counter struct {
	counterMutex *sync.Mutex `json:"-"`
	Value int `json:"value"`
	Step int `json:"step"`
}

/**
 * Получение текущего значение счетчика
 */
func (c *Counter) get() int {
	c.counterMutex.Lock()
	defer c.counterMutex.Unlock()

	return c.Value
}

/**
 * Инкремент счетчика
 */
func (c *Counter) add() int {
	c.counterMutex.Lock()
	c.Value += c.Step
	c.counterMutex.Unlock()

	return c.Value
}


func main() {
	fmt.Println("Start on localhost:3000")

	var C = Counter{new(sync.Mutex), 0, 1}
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout", // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	staticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOptions))

	// Запросы на корень сайта обрабатывает IndexHandler c помощью martini
	m.Get("/", handlers.IndexHandler)

	//Обработчик post запросов на увеличение значения счетчика
	m.Post("/add", func () int {
		C.add()
		fmt.Println(C.get())

		return http.StatusOK
	})

	//Обработчик вывода значения счетчика
	m.Get("/show", func (rnd render.Render) {
		count := C.get()

		if count >= 0 {
			to_json, err := json.Marshal(count)
			if err != nil {
				fmt.Println(err)
			}
			rnd.Data(200, to_json)
		}
	})

	m.Run()
}