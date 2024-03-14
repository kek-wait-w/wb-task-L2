package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Visitor interface {
	visitForFord(*Ford)
	visitForToyota(*Toyota)
	visitForMB(*MB)
}

type Car interface {
	GetModel() string
	check(Visitor)
}

type Ford struct {
	model string
}

func (f *Ford) check(v Visitor) {
	v.visitForFord(f)
}

func (f *Ford) GetModel() string {
	return f.model
}

type Toyota struct {
	model string
	anime string
}

func (t *Toyota) check(v Visitor) {
	v.visitForToyota(t)
}

func (t *Toyota) GetModel() string {
	return t.model
}

type MB struct {
	model      string
	backgammon string
}

func (m *MB) check(v Visitor) {
	v.visitForMB(m)
}

func (m *MB) GetModel() string {
	return m.model
}

type Inspector struct {
}

func (i *Inspector) visitForFord(c *Ford) {
	fmt.Printf("Successful Ford check: %t \n", c.model == "Ford")

}

func (i *Inspector) visitForToyota(c *Toyota) {
	fmt.Printf("Successful Toyota check: %t \n", c.model == "Toyota")
}
func (i *Inspector) visitForMB(c *MB) {
	fmt.Printf("Successful MB check: %t \n", c.model == "MB")
}

func main() {
	Car1 := &Ford{"Ford"}
	Car2 := &Toyota{model: "Toyota"}
	Car3 := &MB{model: "BMW"}

	InspectorDPS := &Inspector{}

	Car1.check(InspectorDPS)
	Car2.check(InspectorDPS)
	Car3.check(InspectorDPS)
}

/*
	Паттерн "посетитель" позволяет нам добавить функционала к существующей структуре, не изменяя ее структуру.
Плюсы:
	1. Упрощает добавление операций, работающих со сложными структурами объектов.
	2. Объединяет родственные операции в одном классе.
	3. Посетитель может накапливать состояние при обходе структуры элементов.
Минусы:
	1. Паттерн не оправдан, если иерархия элементов часто меняется.
	2. Может привести к нарушению инкапсуляции элементов.
*/
