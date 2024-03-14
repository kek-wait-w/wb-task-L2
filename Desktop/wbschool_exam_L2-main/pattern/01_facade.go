package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type SubsystemA struct{}

func (sa *SubsystemA) OperationA() {
	fmt.Println("SubsystemA: OperationA")
}

type SubsystemB struct{}

func (sb *SubsystemB) OperationB() {
	fmt.Println("SubsystemB: OperationB")
}

type SubsystemC struct{}

func (sc *SubsystemC) OperationC() {
	fmt.Println("SubsystemC: OperationC")
}

type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
	subsystemC *SubsystemC
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
		subsystemC: &SubsystemC{},
	}
}

func (f *Facade) Operation1() {
	fmt.Println("Facade: Operation1")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
}

func (f *Facade) Operation2() {
	fmt.Println("Facade: Operation2")
	f.subsystemB.OperationB()
	f.subsystemC.OperationC()
}

/*
Применимость:

Используется, когда нужно предоставить простой интерфейс для сложной системы с множеством взаимосвязанных классов.
Подходит, когда необходимо разделить подсистему на уровни, и предоставить уровень, через который остальная часть системы будет взаимодействовать с подсистемой.

Плюсы:
Уменьшает зависимость между клиентом и подсистемой, так как клиент взаимодействует только с фасадом.
Упрощает работу с подсистемой, скрывая её сложность и предоставляя простой интерфейс.

Минусы:
Может привести к увеличению числа классов в системе, так как для каждой подсистемы нужно создать свой фасад.
Фасад может стать узким местом, если он становится единственной точкой взаимодействия с подсистемой.

Примеры:
Веб-приложение, использующее различные сервисы (например, авторизация, отправка электронной почты, управление данными), может иметь фасад для упрощения взаимодействия с этими сервисами.
Графический пользовательский интерфейс, который скрывает сложность управления виджетами и элементами управления.
*/
