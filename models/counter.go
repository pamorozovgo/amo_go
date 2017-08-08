package models

import (
	"fmt"
	"sync/atomic"
)

/**
 * Наша структура счетчика, содержащая само значение счетчика и шаг прибавления,
 * а также каналы под изменение счетчика и получение текущего значения
 */
type Counter struct {
	value        uint64
	Step         uint64 `json:"step"`
}

/**
 * Получение текущего значение счетчика
 */
func (c *Counter) Get() uint64 {
	val := atomic.LoadUint64(&c.value)
	fmt.Println("Текущее значение счетчика ", val)

	return val
}

/**
 * Инкремент счетчика
 */
func (c *Counter) Add() {
	fmt.Println("Прибавляем ", c.Step)
	atomic.AddUint64(&c.value, c.Step)
}


/**
 * Создание нового счетчика
 */
func NewCounter(step uint64) (c *Counter) {
	c = &Counter{
		value:        0,
		Step:         step,
	}

	return
}