package main

import (
	"fmt"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
//функция проверки на анограммы с помощью мапы
func isAnagram(word1, word2 string) bool {
	if len(word2) != len(word1) {
		return false
	}
	chars := make(map[rune]int)
	for _, v := range word1 {
		chars[v]++
	}
	for _, v := range word2 {
		if _, ok := chars[v]; !ok {
			return false
		}
		if chars[v]-1 == 0 {
			delete(chars, v)
		}
		chars[v]--
	}
	return true
}

//сортировка
func sort(arr []string) []string {
	res := make([]string, len(arr))
	copy(res, arr)
	quickSort(res, 0, len(res)-1)
	return res
}

//алгоритм быстрой сортировки
func quickSort(arr []string, lowIndex, highIndex int) {
	if lowIndex >= highIndex {
		return
	}

	pivotIndex := highIndex - lowIndex/2 + lowIndex - 1
	pivot := arr[pivotIndex]
	arr[pivotIndex], arr[highIndex] = arr[highIndex], arr[pivotIndex]

	leftP := lowIndex
	rightP := highIndex - 1
	for leftP < rightP {
		for arr[leftP] <= pivot && leftP < rightP {
			leftP++
		}
		for arr[rightP] >= pivot && leftP < rightP {
			rightP--
		}
		arr[leftP], arr[rightP] = arr[rightP], arr[leftP]
	}

	if arr[leftP] > arr[highIndex] {
		arr[leftP], arr[highIndex] = arr[highIndex], arr[leftP]
	} else {
		leftP = highIndex
	}

	quickSort(arr, lowIndex, leftP-1)
	quickSort(arr, leftP+1, highIndex)

}

//функция создания множества
func uniqueStrs(s []string) []string {
	strs := make(map[string]int)
	sOut := make([]string, 0, cap(s))
	for _, v := range s {
		if _, ok := strs[v]; !ok {
			sOut = append(sOut, v)
			strs[v]++
		}
	}
	return sOut
}

func GetAnagramSet(strs []string) map[string][]string {
	vocab := make(map[string][]string)
	for _, v := range strs {
		notFound := true
		for k := range vocab {
			if isAnagram(strings.ToLower(k), strings.ToLower(v)) {
				vocab[k] = append(vocab[k], strings.ToLower(v))
				notFound = false
			}
		}
		if notFound {
			vocab[strings.ToLower(v)] = append(vocab[strings.ToLower(v)], strings.ToLower(v))
		}
	}
	for k, v := range vocab {
		if len(v) <= 1 {
			delete(vocab, k)
		}
		vocab[k] = uniqueStrs(v)
		vocab[k] = sort(vocab[k])
	}
	return vocab
}

func main() {
	strs := []string{"Столик", "СЛИТОК", "пятак", "ятпка", "СТОлик", "тяпка", "тяпка", "пятак"}
	out := GetAnagramSet(strs)
	fmt.Println(out)
}
