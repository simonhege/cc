// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
cc_template generates a new go package usable as a base for coding challenges.

The generated package is based on github.com/xeonx/cc and includes
an empty test case.

Usage:
    cc_template name

*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	flag.Parse() // Scans the arg list and sets up flags

	//First argument is package name
	if flag.NArg() < 1 {
		log.Fatal("Undefined target package name")
	}
	pkgName := flag.Arg(0)

	//Create directory in %GOPATH%/src/packagename
	gopath := os.Getenv("GOPATH")
	pkgPath := fmt.Sprintf("%s%csrc%c%s", gopath, os.PathSeparator, os.PathSeparator, pkgName)
	err := os.MkdirAll(pkgPath, os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}

	//Create main file
	filePath := fmt.Sprintf("%s%c%s.go", pkgPath, os.PathSeparator, pkgName)
	log.Print(filePath)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		log.Fatal("File already exists: ", err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("main").Parse(mainTemplate)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(f, pkgName)
	if err != nil {
		log.Fatal(err)
	}

	//Create test file
	testPath := fmt.Sprintf("%s%c%s_test.go", pkgPath, os.PathSeparator, pkgName)
	log.Print(filePath)
	if _, err := os.Stat(testPath); !os.IsNotExist(err) {
		log.Fatal("File already exists: ", err)
	}

	ft, err := os.Create(testPath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = template.New("test").Parse(testTemplate)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(ft, pkgName)
	if err != nil {
		log.Fatal(err)
	}
}

const mainTemplate = `package main

import(
	"log"
	"github.com/xeonx/cc"
)

func main() {
	cc.Run(func() cc.Problem{
		return &problem{}
	})
}
	
type problem struct {
}

func (pb *problem) Init(lines []cc.String) int {

	return 1
}

func (pb *problem) Solve() interface{} {
	log.Println("Solving", pb)
	res := 0
	
	return res
}`

const testTemplate = `package main

import (
	"testing"
	"bytes"
	
	"github.com/xeonx/cc"
)

func TestSample(t *testing.T) {
	in := ` + "`" + `0
` + "`" + `
	out := ` + "``" + `
	
	inBuffer := bytes.NewBufferString(in)
	outBuffer := bytes.Buffer{}
	
	cc.RunFrom(inBuffer, &outBuffer, func() cc.Problem { return &problem{}; })
	
	result := string(outBuffer.Bytes())
	if(result != out) {
		t.Errorf("%v, want %v", result, out)
	}
}`
