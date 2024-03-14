package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

type Server struct {
	Type   string
	Core   int
	Memory int
}

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

type Notebook struct {
	Type   string
	Core   int
	Memory int
}

func (sr Server) GetType() string {
	return sr.Type
}

func (sr Server) PrintDetails() {
	fmt.Printf("[%s] Core: [%d] Memory:[%d]\n", sr.Type, sr.Core, sr.Memory)
}

func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (pc PersonalComputer) GetType() string {
	return pc.Type
}

func (pc PersonalComputer) PrintDetails() {
	fmt.Printf("[%s] Core: [%d] Memory:[%d] Monitor:[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}

func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  32,
		Monitor: true,
	}
}

func (nt Notebook) GetType() string {
	return nt.Type
}

func (nt Notebook) PrintDetails() {
	fmt.Printf("[%s] Core: [%d] Memory:[%d]\n", nt.Type, nt.Core, nt.Memory)
}

func NewNotebook() Computer {
	return Server{
		Type:   NotebookType,
		Core:   4,
		Memory: 8,
	}
}

//фабрика(централизованный коструктор объектов)
func New(typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s Несуществующий тип объекта! \n ", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	}
}

var types = []string{PersonalComputerType, NotebookType, ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computer := New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}

/*
Краткое описание:
	Фабричный метод(виртуальный конструктор) — это порождающий паттерн проектирования, который определяет общий интерфейс
	для создания объектов в супер-классе, позволяя подклассам изменять тип создаваемых объектов.
Плюсы:
	1. Избавляет класс от привязки к конкретным классам продуктов.
	2. Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	3. Упрощает добавление новых продуктов в программу.
	4. Реализует принцип открытости/закрытости.
Минусы:
	1. Может привести к созданию больших параллельных иерархий классов,
		так как для каждого класса продукта надо создать свой подкласс создателя.
*/
