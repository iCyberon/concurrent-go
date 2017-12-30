package main

import "github.com/icyberon/concurrent"
import "fmt"

func divide(i int) int {
	return i / 2
}

func main() {
	a := func(name string) string {
		s := ("Hello " + name)
		return s
	}

	b := func(i int) int {
		return 2 * i
	}

	cg := concurrent.New()
	res, err := cg.Add(a, "").
		Add(b, 10).
		Add(divide, 10).
		Exec()

	if err != nil {
		fmt.Println(err) // []error
	} else {
		fmt.Println(res) // [][]interface{}
	}
}
