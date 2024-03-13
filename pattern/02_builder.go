package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
//Франкенштейна компьютера которую будем возвращать
type Frankenstein struct {
	Head  bool
	Hands bool
	Legs  bool
	Body  bool
}

//Интерфейс описывающий любого Франкенштейна который мы строим
//Каждый метод возвращает интерфейс чтобы можно было настраивать конфигурацию через точку
type FrankensteinBuilderInterface interface {
	Head() FrankensteinBuilderInterface
	Hands() FrankensteinBuilderInterface
	Legs() FrankensteinBuilderInterface
	Body() FrankensteinBuilderInterface
	Build() Frankenstein
}

//Структура реализующая интерфейс
type FrankensteinBuilder struct {
	head  bool
	hands bool
	legs  bool
	body  bool
}

func (f *FrankensteinBuilder) Head() FrankensteinBuilderInterface {
	f.head = true
	return f
}
func (f *FrankensteinBuilder) Hands() FrankensteinBuilderInterface {
	f.hands = true
	return f
}
func (f *FrankensteinBuilder) Legs() FrankensteinBuilderInterface {
	f.legs = true
	return f
}
func (f *FrankensteinBuilder) Body() FrankensteinBuilderInterface {
	f.body = true
	return f
}

//Возвращаем финальный объект Франкенштейна
func (f *FrankensteinBuilder) Build() Frankenstein {
	return Frankenstein{
		Head:  f.head,
		Hands: f.hands,
		Legs:  f.legs,
		Body:  f.body,
	}
}

/* Преимущества и недостатки
+ Позволяет создавать продукты пошагово.
+ Позволяет использовать один и тот же код для создания различных продуктов.
+ Изолирует сложный код сборки продукта от его основной бизнес-логики.
-Усложняет код программы из-за введения дополнительных классов
Паттерн Строитель также используется, когда нужный продукт сложный и требует нескольких шагов для построения.
В таких случаях несколько конструкторных методов подойдут лучше, чем один громадный конструктор.
При использовании пошагового построения объектов потенциальной проблемой является выдача клиенту частично построенного
нестабильного продукта. Паттерн "Строитель" скрывает объект до тех пор, пока он не построен до конца.
В этом примере мы можем создать Франкенштейна с разными параметрами */
