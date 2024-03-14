package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct{}

func (cea *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(cea)
}

type ConcreteElementB struct{}

func (ceb *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(ceb)
}

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteVisitor1 struct{}

func (cv1 *ConcreteVisitor1) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor1: VisitConcreteElementA")
}

func (cv1 *ConcreteVisitor1) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor1: VisitConcreteElementB")
}

type ConcreteVisitor2 struct{}

func (cv2 *ConcreteVisitor2) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor2: VisitConcreteElementA")
}

func (cv2 *ConcreteVisitor2) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor2: VisitConcreteElementB")
}

/*
Применимость:
Используется, когда нужно добавить новую функциональность к классам, но изменение их кода нежелательно или невозможно.
Полезен, когда есть набор объектов различных типов и требуется выполнить различные операции над ними без изменения их классов.

Плюсы:
Позволяет добавлять новые операции, не изменяя классы элементов.
Упрощает процесс добавления новых операций, так как все они объединены в одном месте - в посетителе.

Минусы:
Может привести к усложнению кода, особенно если в системе большое количество классов и операций.
Необходимость реализации методов Accept в каждом классе элементов может привести к нарушению инкапсуляции.

Примеры использования на практике:
Обработка структуры документа (например, HTML-документа) различными операциями, такими как отображение, сохранение в файл и т. д.
Посещение различных узлов дерева (например, DOM-дерева веб-страницы) для выполнения различных операций, таких как обход, изменение стилей и т. д.
Генерация кода из абстрактного синтаксического дерева (AST) различными посетителями для различных языков программирования.
*/
