package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

// Функция для определения индекса поля в списке
func getFieldIndex(fieldNum string, fieldsList []string) (int, bool) {
	fieldNum = strings.TrimSpace(fieldNum)
	for i := range fieldsList {
		if fieldNum == fmt.Sprintf("%d", i+1) {
			return i, true
		}
	}
	return 0, false
}

func main() {
	// Определение флагов
	fields := flag.String("f", "", "Fields to select (columns)")
	delimiter := flag.String("d", "\t", "Delimiter to use")
	separated := flag.Bool("s", false, "Only lines with delimiter")
	flag.Parse()

	// Считывание стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		fieldsList := strings.Split(line, *delimiter)

		// Выбор указанных полей
		selectedFields := make([]string, 0)
		if *fields != "" {
			fieldNumbers := strings.Split(*fields, ",")
			for _, fieldNum := range fieldNumbers {
				if i, ok := getFieldIndex(fieldNum, fieldsList); ok {
					selectedFields = append(selectedFields, fieldsList[i])
				}
			}
		} else {
			selectedFields = fieldsList
		}

		// Вывод выбранных полей
		fmt.Println(strings.Join(selectedFields, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}
