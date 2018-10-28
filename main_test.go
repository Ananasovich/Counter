package main

import "testing"

type testpair struct {
	value  string
	result int
}

var testsLine = []testpair{
	{"Line1\n Text before ending line2!\n some additional text for line3,\n ending", 4},
	{"Test one line string!", 1},
}

func TestCountLines(t *testing.T) {
	for _, pair := range testsLine {
		v := countLines(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"Expected", pair.result,
				"Got", v,
			)
		}
	}
}

var testsRune = []testpair{
	{"Such symbols much wow!", 22},
	{"Считаем кириллицу\n с разными символами \t", 40},
}

func TestCountRunes(t *testing.T) {
	for _, pair := range testsRune {
		v := countRunes(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"Expected", pair.result,
				"Got", v,
			)
		}
	}
}
