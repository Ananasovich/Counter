package main

import "testing"

type testpair struct {
	value  string
	result int
}

var tests = []testpair{
	{"Line1\n Text before ending line2!\n some additional text for line3,\n ending", 4},
	{"Test one line string!", 1},
}

func TestCountLines(t *testing.T) {
	for _, pair := range tests {
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
