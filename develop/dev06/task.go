package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type flags struct {
	fields    int
	delimiter string
	separated bool
}

//инициализация флагов
func flagsInit() flags {
	fl := flags{}
	flag.IntVar(&fl.fields, "f", 0, "'fields' - выбрать поля (колонки)")
	flag.StringVar(&fl.delimiter, "d", "\t", "'delimiter' - использовать другой разделитель")
	flag.BoolVar(&fl.separated, "s", false, "'separated' - только строки с разделителем")
	flag.Parse()

	return fl
}
func main() {

	fl := flagsInit()
	args := flag.Args()

	if fl.fields == 0 {
		log.Fatalln("поля нумеруются с 1")
	}
	if len(args) < 1 {
		log.Fatalln("Пример использования: [флаги] [название файла]")
	}
	fileName := args[len(args)-1]
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	str := Cut(string(file), fl)
	fmt.Println(str)
}

//функция разбивки строк по столбцам
func Cut(text string, fl flags) string {
	var temp string
	//сплитим файл по строкам, а потом используем cut для каждой из них
	splitString := strings.Split(text, "\n")
	// проходим по всем строкам и для каждой вызываем метод Cut
	for _, str := range splitString {
		if res, ok := findInLine(str, fl); ok {
			temp += "\n" + res
		}
	}
	return temp[1:]
}

//поиск определенного столбца в строке
func findInLine(str string, fl flags) (string, bool) {
	// ПРоверка на флаг -s, пропускаем строки без разделителя
	if fl.separated && !strings.Contains(str, fl.delimiter) {
		return "", false
	}
	//Сплитим строку разделителем и выводим нужный столбец
	splitStr := strings.Split(str, fl.delimiter)
	if fl.fields <= len(splitStr) {
		return splitStr[fl.fields-1], true
	}
	return "", false
}
