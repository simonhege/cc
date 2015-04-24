// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cc

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	//DefaultInput is the default name for input files
	DefaultInput = "sample.in"
	//PrintFormat is the format for writing the result in the output file
	PrintFormat = "Case #%d: %v\n"
)

//Problem represents a item of a challenge that must be solved
type Problem interface {
	Init([]String) int
	Solve() interface{}
}

//A ProblemFactory is a factory for creating Problems
type ProblemFactory func() Problem

//Run reads from file, solves and writes problem results in files.
func Run(factory ProblemFactory) {

	flag.Parse() // Scans the arg list and sets up flags

	//if file defined in source, use it, else use DefaultInput
	input := DefaultInput
	if flag.NArg() > 0 {
		input = flag.Arg(0)
	}
	log.Printf("INFO: Selected file: %s\n", input)
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	//Create output writer
	output := fmt.Sprintf("%s.%s.out", input, time.Now().Format("2006-01-02T15-04-05"))
	w, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	RunFrom(f, w, factory)

	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}

	//Exit
	log.Printf("THE END (result written in %s)\n", output)
}

//RunFrom reads from io.Reader, solve and writes problem results in io.Writer
func RunFrom(r io.Reader, w io.Writer, factory ProblemFactory) {

	//Read file
	sdata, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	linesOriginal := strings.Split(string(sdata), "\n")
	iLineCount := len(linesOriginal)
	if len(linesOriginal[iLineCount-1]) == 0 {
		iLineCount--
	}
	log.Printf("INFO: %d readed lines\n", iLineCount)

	//Check line count
	if iLineCount < 1 {
		log.Fatal("ERROR: No lines\n")
	}

	//Convert to cc.String
	lines := make([]String, iLineCount, iLineCount)
	for i, line := range linesOriginal {
		if i < iLineCount {
			lines[i] = String(line)
		}
	}

	//Pb count is on first line. Check total line count
	pbCount := lines[0].Int()
	log.Printf("INFO: %d problems\n", pbCount)

	if (pbCount + 1) != iLineCount {
		log.Fatalf("ERROR: Wrong lines count: %d read\n", iLineCount)
	}

	//Read pb data
	data := make([]Problem, pbCount, pbCount)
	cursor := 1
	for i := 0; i < pbCount; i++ {
		data[i] = factory()
		cursor += data[i].Init(lines[cursor:])
	}

	result := make([]interface{}, pbCount, pbCount)
	log.Printf("INFO: Start solving problems\n")
	for i := 0; i < pbCount; i++ {
		result[i] = data[i].Solve()
	}
	log.Printf("INFO: End solving problems\n")

	for i := 0; i < pbCount; i++ {
		fmt.Fprintf(w, PrintFormat, i+1, result[i])
	}

}
