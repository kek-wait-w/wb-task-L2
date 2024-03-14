package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	Execute()
}

type Receiver struct {
	Action string
}

func (r *Receiver) ActionMethod() {
	fmt.Printf("Receiver executing action: %s\n", r.Action)
}

type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		receiver: receiver,
	}
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.ActionMethod()
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

/*
Применимость:
Паттерн команды полезен, когда нужно параметризовать объекты событиями или запросами.
Используется для реализации отмены операций, повтора операций или журналирования запросов.
Удобен для реализации очередей запросов или планирования операций.

Плюсы:
Уменьшает связность между отправителем и получателем команды.
Позволяет создавать составные команды из простых команд.
Поддерживает отмену операций и повтор выполнения.

Минусы:
Может привести к увеличению числа классов в системе.

Примеры использования на практике:
Обработка пользовательских действий в графическом интерфейсе, таких как кнопки и меню.
Управление заказами в интернет-магазине, где каждая команда представляет собой действие, которое нужно выполнить над заказом.
Реализация систем умного дома, где команды могут управлять устройствами, такими как свет, температура и безопасность.

*/
