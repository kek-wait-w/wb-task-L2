package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	After   int
	Before  int
	Context int
	Count   bool
	Ignore  bool
	Invert  bool
	Fixed   bool
	LineNum bool
}

type result struct {
	strs     [][]string
	lineNums []int
}

// Инициализация флагов
func flagsInit() flags {
	fl := flags{}
	flag.IntVar(&fl.After, "A", 0, "Print N lines after each match")
	flag.IntVar(&fl.Before, "B", 0, "Print N lines before each match")
	flag.IntVar(&fl.Context, "C", 0, "Print N lines before and after each match")
	flag.BoolVar(&fl.Count, "c", false, "Print only a count of selected lines")
	flag.BoolVar(&fl.Ignore, "i", false, "Ignore case distinctions")
	flag.BoolVar(&fl.Invert, "v", false, "Invert the sense of matching")
	flag.BoolVar(&fl.Fixed, "F", false, "Interpret pattern as a literal string")
	flag.BoolVar(&fl.LineNum, "n", false, "Print line number with output lines")
	flag.Parse()

	// Обработка флагов контекста
	if fl.Context > fl.After {
		fl.After = fl.Context
	}
	if fl.Context > fl.Before {
		fl.Before = fl.Context
	}

	return fl
}

func Grep(desired string, text []string, fl flags) result {
	res := result{}
	var condition bool

	for index, str := range text {
		// Применение флага игнорирования регистра
		if fl.Ignore {
			str = strings.ToLower(str)
			desired = strings.ToLower(desired)
		}

		// Проверка условий совпадения
		if fl.Fixed {
			condition = desired == str
		} else {
			condition = strings.Contains(str, desired)
		}

		// Применение флага инверсии совпадений
		if fl.Invert {
			condition = !condition
		}

		temp := make([]string, 0)
		if condition {
			// Определение диапазона строк для вывода в зависимости от флагов
			var upRange, downRange = 0, len(text) - 1
			if d := index - fl.Before; d > upRange {
				upRange = d
			}
			if d := index + fl.After; d < downRange {
				downRange = d
			}
			for i := upRange; i <= downRange; i++ {
				temp = append(temp, text[i])
			}
			res.lineNums = append(res.lineNums, index)
			res.strs = append(res.strs, temp)
		}
	}
	return res
}

func main() {
	fl := flagsInit()

	args := flag.Args()

	if len(args) < 2 {
		log.Fatalln("Usage: grep [OPTIONS] PATTERN [FILE]")
	}

	slicePhrase := args[:len(args)-1]
	desired := strings.Join(slicePhrase, " ")

	// Чтение файла
	file, err := ioutil.ReadFile(args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}

	// Разбивка файла на строки
	splitString := strings.Split(string(file), "\n")

	// Применение grep и вывод результатов
	res := Grep(desired, splitString, fl)
	if fl.Count {
		fmt.Printf("Matches: %v \n", len(res.lineNums))
	} else if fl.LineNum {
		for _, v := range res.lineNums {
			fmt.Println(v)
		}
	} else {
		fmt.Println(res.strs)
	}
}
