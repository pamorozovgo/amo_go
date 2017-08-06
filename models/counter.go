package models

import (
	"fmt"
	"runtime"
)

/**
 * Наша структура счетчика, содержащая само значение счетчика и шаг прибавления,
 * а также каналы под изменение счетчика и получение текущего значения
 */
type Counter struct {
	value        int
	Step         int `json:"step"`
	addChan      chan int
	curCountChan chan int
}

/**
 * Получение текущего значение счетчика
 */
func (c *Counter) Get() int {
	return <-c.curCountChan
}

/**
 * Инкремент счетчика
 */
func (c *Counter) Add() {
	c.addChan <- c.Step
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
		value:        0,
		Step:         step,
		addChan:      make(chan int),
		curCountChan: make(chan int),
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
