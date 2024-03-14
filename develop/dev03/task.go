package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

type Comparator interface {
	compare(a, b string, f flags) int //реализуем перегрузку опертаоров через интерфейс
	//a>b - 1; a=b - 0; a<b - -1;
}

type CompareDefault struct{}

//метод дефолтного компоратора
func (c *CompareDefault) compare(a, b string, f flags) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

type CompareColumn struct{}

//метод компоратора сравнивания строк
func (c *CompareColumn) compare(a, b string, f flags) int {

	as := strings.Split(a, " ")[f.k-1]
	bs := strings.Split(b, " ")[f.k-1]

	if f.n {
		an, err := strconv.Atoi(as)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: converting to int\n")
			return -1
		}
		bn, err := strconv.Atoi(bs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: converting to int\n")
			return 1
		}
		if an > bn {
			return 1
		}
		if an < bn {
			return -1
		}
		if an == bn {
			return 0
		}
	}

	if as > bs {
		return 1
	}
	if as < bs {
		return -1
	}
	if as == bs {
		return 0
	}
	return 0
}

type CompareInt struct{}

//метод компоратора сравнинвания чисел
func (c *CompareInt) compare(a, b string, f flags) int {

	as := strings.Split(a, " ")
	bs := strings.Split(b, " ")
	aInt := make([]int, 0)
	bInt := make([]int, 0)
	for _, v := range as {
		val, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		aInt = append(aInt, val)
	}
	for _, v := range bs {
		val, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		bInt = append(bInt, val)
	}

	if len(aInt) <= len(bInt) {
		for i := 0; i < len(aInt); i++ {
			if aInt[i] < bInt[i] {
				return -1
			}
			if aInt[i] > bInt[i] {
				return 1
			}
		}
		if len(aInt) == len(bInt) {
			return 0
		}
		return -1
	} else if len(bInt) < len(aInt) {
		for i := 0; i < len(bInt); i++ {
			if aInt[i] < bInt[i] {
				return -1
			}
			if aInt[i] > bInt[i] {
				return 1
			}
		}
		return 1
	}
	return 0
}

type flags struct {
	k int
	n bool
	r bool
	u bool
	i string
}

//функция обработки флагов и соответсвующей логики
func Sort(arr []string, flags flags) []string {
	res := make([]string, len(arr))
	copy(res, arr)

	if flags.k <= 0 && !flags.n {
		quickSort(res, 0, len(res)-1, flags, &CompareDefault{})
	}
	if flags.k > 0 {
		quickSort(res, 0, len(res)-1, flags, &CompareColumn{})
	}
	if flags.n && flags.k <= 0 {
		quickSort(res, 0, len(res)-1, flags, &CompareInt{})
	}

	if flags.r {
		res = reverseString(res)
	}
	if flags.u {
		res = uniqueStrs(res)
	}
	return res
}

//функция переворачивания строки
func reverseString(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//функция создания множества
func uniqueStrs(s []string) []string {
	strs := make(map[string]int)
	res := make([]string, 0, cap(s))
	for _, v := range s {
		if _, ok := strs[v]; !ok {
			res = append(res, v)
			strs[v]++
		}
	}
	return res
}

//алгоритм быстрой сортировки
func quickSort(arr []string, lowIndex, highIndex int, fl flags, cmp Comparator) {
	if lowIndex >= highIndex {
		return
	}
	pivotIndex := (highIndex-lowIndex)/2 + lowIndex
	pivot := arr[pivotIndex]
	arr[pivotIndex], arr[highIndex] = arr[highIndex], arr[pivotIndex]

	leftP := lowIndex
	rightP := highIndex - 1
	for leftP < rightP {
		for cmp.compare(arr[leftP], pivot, fl) <= 0 && leftP < rightP {
			leftP++
		}
		for cmp.compare(arr[rightP], pivot, fl) >= 0 && leftP < rightP {
			rightP--
		}

		arr[leftP], arr[rightP] = arr[rightP], arr[leftP]
	}

	if cmp.compare(arr[leftP], arr[highIndex], fl) >= 1 {
		arr[leftP], arr[highIndex] = arr[highIndex], arr[leftP]
	} else {
		leftP = highIndex
	}

	quickSort(arr, lowIndex, leftP-1, fl, cmp)
	quickSort(arr, leftP+1, highIndex, fl, cmp)

}

//инициализация флагов
func flagsInit() *flags {
	f := new(flags)
	flag.StringVar(&f.i, "i", "./input.txt", "input file")
	flag.IntVar(&f.k, "k", -1, "sort column")
	flag.BoolVar(&f.n, "n", false, "sort by numeric value")
	flag.BoolVar(&f.r, "r", false, "sort reverse")
	flag.BoolVar(&f.u, "u", false, "no repeats")
	flag.Parse()
	return f
}

//чтение файла
func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	strs := make([]string, 0)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strs
}

func main() {
	fl := flagsInit()
	strs := readFile(fl.i)
	sorted := Sort(strs, *fl)
	for _, v := range sorted {
		fmt.Println(v)
	}

}
