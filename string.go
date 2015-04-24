// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cc

import (
	"log"
	"strconv"
	"strings"
)

//String provides helper function when reading problem data from string.
type String string

//String convert the value into a standard string
func (s String) String() string {
	return string(s)
}

//Int convert the value into an int
func (s String) Int() int {
	v, err := strconv.Atoi(string(s))
	if err != nil {
		log.Fatalf("ERROR: Conversion to int failed: %s read\n", s)
	}
	return v
}

//Int2 convert the value into two ints
func (s String) Int2() (int, int) {
	array := s.IntArray()
	if len(array) != 2 {
		log.Fatalf("ERROR: Conversion to 2 ints failed: %s read\n", s)
	}
	return array[0], array[1]
}

//Int3 convert the value into three ints
func (s String) Int3() (int, int, int) {
	array := s.IntArray()
	if len(array) != 3 {
		log.Fatalf("ERROR: Conversion to 3 ints failed: %s read\n", s)
	}
	return array[0], array[1], array[2]
}

//IntArray convert the value into an array of ints
func (s String) IntArray() []int {
	tokens := s.Split()
	result := make([]int, len(tokens))
	for i, v := range tokens {
		result[i] = v.Int()
	}
	return result
}

//SingleDigitIntArray convert the value into an array of ints. Each characters is converted individually.
func (s String) SingleDigitIntArray() []int {

	result := make([]int, len(s))

	for idx, sp := range string(s) {
		i, err := strconv.Atoi(string(sp))
		if err != nil {
			log.Fatal(err)
		}
		result[idx] = i
	}

	return result
}

//Int64 convert the value into an int64
func (s String) Int64() int64 {
	v, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		log.Fatalf("ERROR: Conversion to int failed: %s read\n", s)
	}
	return v
}

//Int642 convert the value into two int64
func (s String) Int642() (int64, int64) {
	array := s.Int64Array()
	if len(array) != 2 {
		log.Fatalf("ERROR: Conversion to 2 int64 failed: %s read\n", s)
	}
	return array[0], array[1]
}

//Int643 convert the value into three int64
func (s String) Int643() (int64, int64, int64) {
	array := s.Int64Array()
	if len(array) != 3 {
		log.Fatalf("ERROR: Conversion to 3 int64 failed: %s read\n", s)
	}
	return array[0], array[1], array[2]
}

//Int64Array convert the value into an array of int64
func (s String) Int64Array() []int64 {
	tokens := s.Split()
	result := make([]int64, len(tokens))
	for i, v := range tokens {
		result[i] = v.Int64()
	}
	return result
}

//Split the value into a array of String, using spaces as separator.
func (s String) Split() []String {
	tokens := strings.Fields(string(s))
	result := make([]String, len(tokens))
	for i, v := range tokens {
		result[i] = String(v)
	}
	return result

}
