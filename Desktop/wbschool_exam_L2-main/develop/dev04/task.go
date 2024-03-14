package main

import (
	"sort"
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

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	anagramGroups := make(map[string][]string)

	for _, word := range words {
		// Приводим слово к нижнему регистру и сортируем его буквы.
		sortedWord := sortString(strings.ToLower(word))

		// Добавляем слово в соответствующую группу анаграмм.
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}

	for _, group := range anagramGroups {
		if len(group) > 1 {
			// Сортируем слова в группе перед добавлением в результат.
			sort.Strings(group)
			anagrams[group[0]] = group
		}
	}

	return anagrams
}

// sortString сортирует символы в строке.
func sortString(s string) string {
	sortedRunes := []rune(s)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток"}
	anagrams := findAnagrams(words)

	for key, value := range anagrams {
		println(key, ":", strings.Join(value, ", "))
	}
}
