package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Button struct {
	command Command
}

type OnCommand struct {
	device Device
}

type Command interface {
	execute()
}

type OffCommand struct {
	device Device
}

type Device interface {
	on()
	off()
}

type Tv struct {
	isRunning bool
}

func (b *Button) press() {
	b.command.execute()
}

func (c *OnCommand) execute() {
	c.device.on()
}

func (c *OffCommand) execute() {
	c.device.off()
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}

/*
Описание:
Это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов,
ставить запросы в очередь, логировать их, а также поддерживать отмену операций.
Преимущества:
1. Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
2. Позволяет реализовать простую отмену и повтор операций.
3. Позволяет реализовать отложенный запуск операций.
4. Позволяет собирать сложные команды из простых.
5. Реализует принцип открытости/закрытости.
Недостатки:
1. Усложняет код программы из-за введения множества дополнительных классов.
*/
