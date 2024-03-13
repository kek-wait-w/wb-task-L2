package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input  []string
	output map[string][]string
}{
	{[]string{"Столик", "СЛИТОК", "пятак", "ятпка", "СТОлик", "тяпка", "тяпка", "пятак"}, map[string][]string{
		"пятак":  {"пятак", "тяпка", "ятпка"},
		"столик": {"слиток", "столик"},
	}},
	{[]string{"Раки", "Каир", "ИРАК", "КУЛОН", "Клоун", "УКлон", "унокл"}, map[string][]string{
		"кулон": {"клоун", "кулон", "уклон", "унокл"},
		"раки":  {"ирак", "каир", "раки"},
	}},
}

func TestSort(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, GetAnagramSet(tt.input), tt.output)
	}
}
