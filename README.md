# CC

Package cc provides functionalities related to coding challenges, where
problems are read from a text file and solutions are written back in an
other file.

The package name `cc` stands for "Coding Challenge"

## Install

	go get github.com/xeonx/cc

## Docs

<http://godoc.org/github.com/xeonx/cc>

## Usage
Implement the `cc.Problem` interface and call `cc.Run` in the main function.

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
			return "EVEN"
		}
		return "ODD"
	}
	
	
## License

This code is licensed under the MIT license. See [LICENSE](https://github.com/xeonx/geodesic/blob/master/LICENSE).
