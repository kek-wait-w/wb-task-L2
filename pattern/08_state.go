package pattern

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	ExecuteState(c *Context)
}

type RedState struct{}

func (r *RedState) ExecuteState(c *Context) {
	fmt.Println("Светофор переключен в красный")
	time.Sleep(3 * time.Second) // Подождем 3 секунды
	c.SetState(&GreenState{})   // Переключаемся на следующее состояние
}

type GreenState struct{}

func (g *GreenState) ExecuteState(c *Context) {
	fmt.Println("Светофор переключен в зеленый")
	time.Sleep(3 * time.Second) // Подождем 3 секунды
	c.SetState(&YellowState{})  // Переключаемся на следующее состояние
}

type YellowState struct{}

func (y *YellowState) ExecuteState(c *Context) {
	fmt.Println("Светофор переключен в желтый")
	time.Sleep(3 * time.Second) // Подождем 3 секунды
	c.SetState(&RedState{})     // Переключаемся на следующее состояние
}

type Context struct {
	state State
}

func (c *Context) SetState(s State) {
	c.state = s
}

func (c *Context) Request() {
	c.state.ExecuteState(c)
}

/*
Применимость паттерна:

Когда объект должен менять свое поведение в зависимости от своего состояния.
Когда код содержит множество условных операторов, связанных с различными состояниями объекта, и нужно упростить их управление.

Плюсы:
Улучшает читаемость и поддерживаемость кода, за счет выноса логики состояний из основного класса в отдельные классы.
Упрощает добавление новых состояний и изменение поведения объекта без изменения существующего кода.

Минусы:
Может привести к увеличению количества классов в программе из-за создания отдельных классов для каждого состояния.
Усложняет взаимодействие между объектами из-за необходимости обмена сообщениями между состояниями.

Примеры использования:
Управление заказами , где состояния заказа могут быть "новый", "оплаченный", "отгруженный", "доставленный" и т.д.
*/
