package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func getTime() (time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return time, err
}

func main() {
	time, err := getTime()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println("Точное время:")
	fmt.Println(time.Hour(), "часов", time.Minute(), "минут", time.Second(), "секунд")
}
