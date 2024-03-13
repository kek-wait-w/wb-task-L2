package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Payment interface {
	Pay()
}

func buyProduct(payment Payment) {
	payment.Pay()
}

type Card struct {
}

func (p *Card) Pay() {
	fmt.Println("Card transactions")
}

type Cash struct {
}

func (p *Cash) Pay() {
	fmt.Println("Cash transactions")
}

type Bitcoin struct {
}

func (p *Bitcoin) Pay() {
	fmt.Println("Bitcoin transactions")
}

func main() {
	var payment Payment

	payment = &Card{}
	buyProduct(payment)

	payment = &Cash{}
	buyProduct(payment)

	payment = &Bitcoin{}
	buyProduct(payment)

}

/*
	Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает
	каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
Плюсы:
	1. Горячая замена алгоритмов на лету.
	2. Изолирует код и данные алгоритмов от остальных классов.
	3. Уход от наследования к делегированию.
	4. Реализует принцип открытости/закрытости.
Минусы:
	1. Усложняет программу за счёт дополнительных классов.
	2. Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/
