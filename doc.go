// Copyright 2015 Simon HEGE. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package cc provides functionalities related to coding challenges, where
problems are read from a text file and solutions are written back in an
other file.

The package name "cc" stands for "Coding Challenge"

See sources at http://github.com/xeonx/cc

For example to create a program testing whether the input is odd or even:

	package main

	import(
		"github.com/xeonx/cc"
	)

	func main() {
		cc.Run(func() cc.Problem{
			return &problem{}
		})
	}

	type problem struct {
		a int
	}

	func (pb *problem) Init(lines []cc.String) int
		pb.a = lines[0].Int()
		return 1
	}

	func (pb *problem) Solve() interface{}
		if a%2 == 0 {
			return "ODD"
		}
		return "EVEN"
	}


*/
package cc
