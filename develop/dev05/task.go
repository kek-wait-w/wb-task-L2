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

//инициализация флагов
func flagsInit() flags {
	fl := flags{}
	flag.IntVar(&fl.After, "A", 0, "'after' печатать +N строк после совпадения")
	flag.IntVar(&fl.Before, "B", 0, "'before' печатать +N строк до совпадения")
	flag.IntVar(&fl.Context, "C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&fl.Count, "c", false, "'count' (количество строк)")
	flag.BoolVar(&fl.Ignore, "i", false, "'ignore-case' (игнорировать регистр)")
	flag.BoolVar(&fl.Invert, "v", false, "'invert' (вместо совпадения, исключать)")
	flag.BoolVar(&fl.Fixed, "F", false, "'fixed', точное совпадение со строкой")
	flag.BoolVar(&fl.LineNum, "n", false, "'line num', печатать номер строки")
	flag.Parse()
	if fl.Context > fl.After {
		fl.After = fl.Context
	}
	if fl.Context > fl.Before {
		fl.Before = fl.Context
	}
	return fl
}
func main() {
	fl := flagsInit()

	//собираем аргументы, искомая строка и название файла
	args := flag.Args()

	//если аргументов нехватает пишем как использовать программу
	if len(args) < 2 {
		log.Fatalln("Чтобы начать поиск: [флаги] [искомая строка] [название файла]")
	}

	//искомая фраза
	slicePhrase := args[:len(args)-1]
	//объединяем в одну строку если фраза состоит из нескольких слов
	desired := strings.Join(slicePhrase, " ")

	//Считываем файл
	file, err := ioutil.ReadFile(args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}

	//сплитим файл построчно
	splitString := strings.Split(string(file), "\n")
	//Используем grep и записываем результат, затем печатаем его
	res := Grep(desired, splitString, fl)
	if fl.Count {
		fmt.Printf("Совпадений: %v \n", len(res.lineNums))

	} else if fl.LineNum {
		for _, v := range res.lineNums {
			fmt.Println(v)
		}

	} else {
		fmt.Println(res.strs)
	}
}

//функция поиска фразы или строки в файле с применением доп.условий
func Grep(desired string, text []string, fl flags) result {
	//слайс с результатами поиска. Это массив из узлов, каждый из которых массив ключ значение
	res := result{}
	var condition bool // условие сравнения

	//проходим построчно по файлу
	for index, str := range text {
		// если применен -i, убираем регистр
		if fl.Ignore {
			//Переводим текущую строку в нижний регистр
			str = strings.ToLower(str)
			//Переводим искомую фразу в нижний регистр
			desired = strings.ToLower(desired)
		}
		//Проверяем условия
		if fl.Fixed {
			condition = desired == str // полное совпадение строки
		} else {
			condition = strings.Contains(str, desired) // совпадение подстроки
		}

		//флаг исключения
		if fl.Invert {
			condition = !condition
		}

		//Создаем объект temp для добавления в результат
		temp := make([]string, 0)
		// если условие выполняется то значит в эту строку записываем
		if condition {
			//Определяем количество строк для печати в зависимости от флагов
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
