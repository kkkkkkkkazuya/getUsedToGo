package main

import (
	"fmt"
)

const h string = "constant"
const n = 500000

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "Apple"
	fmt.Println(f)

	var g int
	fmt.Println(g)

	fmt.Println(h)

	const z = 3e20 / n
	fmt.Println(z)

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	var arr [5]string = [5]string{"a", "b", "c", "d", "e"}

	for n := 0; n < 5; n++ {
		fmt.Println(arr[n])
	}
}
