package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool
	UpdateSource bool
}

type Device1 struct {
	Name string
	Next Service
}

type UpdateDataService struct {
	Name string
	Next Service
}

type DataService struct {
	Next Service
}

func (device *Device1) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Data from device [%s] already get. \n", device.Name)
		device.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device [%s]. \n", device.Name)
	data.GetSource = true
	device.Next.Execute(data)
}

func (device *Device1) SetNext(srv Service) {
	device.Next = srv
}

func (upd *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Data from device [%s] already update \n", upd.Name)
		upd.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from service [%s]. \n", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDataService) SetNext(srv Service) {
	upd.Next = srv
}

func (upd *DataService) Execute(data *Data) {
	if !data.GetSource {
		fmt.Printf("Data not update")
		return
	}
	fmt.Printf("Data save.")
}

func (upd *DataService) SetNext(srv Service) {
	upd.Next = srv
}

func main() {
	device := &Device1{Name: "Device-1"}
	updateSvc := &UpdateDataService{Name: "Update-1"}
	dataSvc := &DataService{}
	device.SetNext(updateSvc)
	updateSvc.SetNext(dataSvc)
	data := &Data{}
	device.Execute(data)
}

/*
Описание:
Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам
и стоит ли передавать запрос дальше по цепи.
Преимущества:
1. Уменьшает зависимость между клиентом и обработчиками.
2. Реализует принцип единственной обязанности.
3. Реализует принцип открытости/закрытости.
Недостатки:
1. Запрос может остаться никем не обработанным.
*/

/*
	Цепочка ответственностей может быть как линейной, так и разветвленной.
*/
