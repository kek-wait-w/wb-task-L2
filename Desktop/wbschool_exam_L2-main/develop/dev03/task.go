/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Флаги для настройки сортировки
	column := flag.Int("k", 1, "Номер колонки для сортировки")
	numeric := flag.Bool("n", false, "Сортировка по числовому значению")
	reverse := flag.Bool("r", false, "Сортировка в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	// Считывание ввода со стандартного потока ввода
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
		os.Exit(1)
	}

	// Разбиение строк на колонки
	rows := make([][]string, len(lines))
	for i, line := range lines {
		rows[i] = strings.Fields(line)
	}

	// Сортировка
	sort.SliceStable(rows, func(i, j int) bool {
		columnIndex := *column - 1
		if columnIndex >= len(rows[i]) || columnIndex >= len(rows[j]) {
			return false // Пропускаем строки с недостаточным количеством колонок
		}

		// Применение ключа -n для сортировки по числовому значению
		if *numeric {
			val1, err1 := strconv.Atoi(rows[i][columnIndex])
			val2, err2 := strconv.Atoi(rows[j][columnIndex])
			if err1 == nil && err2 == nil {
				if val1 != val2 {
					return val1 < val2
				}
			}
		}

		// Обратный порядок сортировки
		if *reverse {
			return rows[i][columnIndex] > rows[j][columnIndex]
		}

		return rows[i][columnIndex] < rows[j][columnIndex]
	})

	// Удаление повторяющихся строк
	if *unique {
		uniqueRows := make(map[string]bool)
		uniqueData := make([][]string, 0)

		for _, row := range rows {
			str := strings.Join(row, " ")
			if !uniqueRows[str] {
				uniqueRows[str] = true
				uniqueData = append(uniqueData, row)
			}
		}
		rows = uniqueData
	}

	// Вывод отсортированных строк
	for _, row := range rows {
		fmt.Println(strings.Join(row, " "))
	}
}
