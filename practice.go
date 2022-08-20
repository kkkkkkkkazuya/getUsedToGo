package main

import (
	"fmt"
)

const h string = "constant"
const n = 500000

func all(num1, num2 int) int {
	return num1 + num2
}

func allString(word string) string {
	return word
}

func calAll(x, y int) (int, int) {
	return (x * y), (x / y)
}

func wordAll(red, blue string) (string, string) {
	return red, blue
}

// func numCal(o, p int) (int, int) {
// 	a := o * p
// 	b := o / p
// 	return a, b
// }

func numCal(o, p int) (a int, b int) {
	a = o * p
	b = o / p
	return
}

// 構造体
type Student struct {
	name          string
	math, english float64
}

type User struct {
	gender string
	age    int
}

func main() {
	// var a = "initial"
	// fmt.Println(a)

	// var b, c int = 1, 2
	// fmt.Println(b, c)

	// var d = true
	// fmt.Println(d)

	// var e int
	// fmt.Println(e)

	// f := "Apple"
	// fmt.Println(f)

	// var g int
	// fmt.Println(g)

	// fmt.Println(h)

	// const z = 3e20 / n
	// fmt.Println(z)

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	var arr [5]string = [5]string{"a", "b", "c", "d", "e"}

	for n := 0; n < 5; n++ {
		fmt.Println(arr[n])
	}

	cal := all(2, 3)
	calString := allString("practice")
	fmt.Println(calString, cal)

	result1, result2 := calAll(6, 3)

	fmt.Println(result1, result2)

	wordAll1, wordAll2 := wordAll("赤", "青")

	fmt.Println(wordAll1, wordAll2)

	numCalResult1, numCalResult2 := numCal(12, 3)

	fmt.Println(numCalResult1, numCalResult2)

	//無名関数
	hello := func(greeting string) {
		fmt.Println(greeting)
	}

	hello("Good morning")

	func(greetingNight string) {
		fmt.Println(greetingNight)
	}("Good evening")

	//構造体の初期化

	// s := Student{"kojima", 80, 70}
	// fmt.Println(s)

	s := Student{name: "kazuya", english: 80}
	fmt.Println(s)

	//:=を使用するバージョン
	t := User{gender: "male", age: 20}
	fmt.Println(t)

	//varを使って直接代入バージョン
	var u User
	u.gender = "female"
	u.age = 23
	fmt.Println(u)
}
