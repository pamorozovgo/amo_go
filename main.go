package main

import (
	"fmt"
	"./handlers"
	"encoding/json"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"runtime"
)


/**
 * Наша структура счетчика, содержащая само значение счетчика и шаг прибавления,
 * а также каналы под изменение счетчика и получение текущего значения
 */
type Counter struct {
	value 			int
	Step 			int `json:"step"`
	addChan 		chan int
	curCountChan 	chan int
}

/**
 * Получение текущего значение счетчика
 */
func (c *Counter) get() int {
	return <-c.curCountChan
}

/**
 * Инкремент счетчика
 */
func (c *Counter) add() {
	c.addChan <-c.Step
}

/**
 * Применение изменений счетчика
 */
func (c *Counter) apply(step int) {
	fmt.Println("Прибавляем ", step)

	c.value += step
}


/**
 * Создание нового счетчика
 */
func NewCounter(step int) (c *Counter) {
	c = &Counter{
		value:			0,
		Step:			step,
		addChan: 		make(chan int),
		curCountChan:	make(chan int),
	}

	go c.run()

	return
}

/**
 * Бесконечный цикл-обработчик изменений значения счетчика и получения текущего значения
 */
func (c *Counter) run() {
	//Единожды инициализируем переменную для хранения шага счетчика
	var step int
	for {
		select {
			//При увеличении счетчика с канала приходит шаг изменения счетчика
			case step = <-c.addChan:
				c.apply(step)
			case c.curCountChan <- c.value:
				//в case`е уже записали значение текущего счетчика в канал
				fmt.Println(`Текущее значение счетчика`, c.value)
		}
		//Даю возможность runtime`у свободно переключаться между потоками без опасений блокировок
		runtime.Gosched()
	}

}

func main() {
	fmt.Println("Start on localhost:3000")

	// Объявляю счетчик, который сразу будет обрабатывать данные с каналов
	var c = NewCounter(1)

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
		c.add()

		return http.StatusOK
	})

	//Обработчик вывода значения счетчика
	m.Get("/show", func (rnd render.Render) {
		to_json, err := json.Marshal(c.get())
		if err != nil {
			fmt.Println(err)
		}
		rnd.Data(200, to_json)
	})

	m.Run()
}