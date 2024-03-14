package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
package builder

type Product struct {
	Part1 string
	Part2 string
	Part3 string
}

type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetProduct() *Product
}

type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		product: &Product{},
	}
}

func (cb *ConcreteBuilder) BuildPart1() {
	cb.product.Part1 = "Part 1"
}

func (cb *ConcreteBuilder) BuildPart2() {
	cb.product.Part2 = "Part 2"
}

func (cb *ConcreteBuilder) BuildPart3() {
	cb.product.Part3 = "Part 3"
}

func (cb *ConcreteBuilder) GetProduct() *Product {
	return cb.product
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
}

/*
Применимость:
Паттерн "Строитель" применяется, когда объект должен быть создан пошагово, с возможностью настройки его параметров.
Используется, когда создание объекта представляет собой сложный процесс, который можно разбить на отдельные шаги.

Плюсы:
Позволяет создавать сложные объекты пошагово и изменять порядок или конфигурацию шагов.
Изолирует код создания объекта от его представления, что позволяет использовать один и тот же процесс конструирования для различных представлений объекта.

Минусы:
Может привести к избыточному увеличению числа классов в системе, особенно если объект имеет много параметров.
Увеличивает сложность кода из-за наличия большого количества классов и интерфейсов.

Примеры использования на практике:
Построение сложных объектов, таких как документы, отчеты, графические интерфейсы и т. д.
Конструирование объектов с различными конфигурациями, такими как конфигурация компьютера или автомобиля.
Создание паттерна "Фасад", который может использовать паттерн "Строитель" для создания сложных объектов.
 */