package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unPack(str string) string {
	var errFlag, escapeFlag bool
	var res, buf string
	unp := strings.Split(str, "")
	// основной алгоритм распаковки
	for _, i := range unp {

		if i == "\\" && !escapeFlag {
			escapeFlag = true
			continue
		}

		cap, err := strconv.Atoi(i)
		if err != nil && i != "\\" {
			res += i
			buf = i
			errFlag = true
			continue
		}

		if escapeFlag {
			res += i
			buf = i
			escapeFlag = false
			continue
		}

		for j := 1; j < cap; j++ {
			res += buf
		}
	}

	//обработка ошибки, если одни числа
	if !errFlag {
		log.Fatal("Error: No letter, wrong data")
	}

	return res
}

func main() {
	var str string
	fmt.Scanln(&str)
	// проверка на пустую строку
	if str != "" {
		fmt.Println(unPack(str))
	} else {
		fmt.Println("")
	}
}
