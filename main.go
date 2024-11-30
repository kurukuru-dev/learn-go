package main

import (
	"fmt"
)

// import ( "fmt" "math" )

// variable practice
func variable() {
	var a string = "Hello"
	b := "World"

	fmt.Println(a, b)
}

// type conversion
func typeConversion() {
	var a int = 1
	b := float64(a)

	fmt.Println(a, b)
}

// constant practice
func constant() {
	const a = 1
	const b = "Hello"

	// const c = make([int[2]]int)

	fmt.Println(a, b)
}

// loop practice
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("Infinite loop")
		break
	}
}

// if else practice
func ifElse() {
	a := 1

	if a == 1 {
		fmt.Println("a is 1")
	} else {
		fmt.Println("a is not 1")
	}
}

// switch practice
func switchPractice(a int) {
	switch a {
		case 1:
			fmt.Println("a is 1")
		case 2:
			fmt.Println("a is 2")
		default:
			fmt.Println("a is not 1 or 2")
	}
}

// defer practice
func deferPractice() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

// pointer practice
func pointerPractice() {
	a := 1
	// bにはaのアドレスが入る。つまり、bはポインタ
	b := &a

	fmt.Println(a, b)
}

// struct practice
func structPractice() {
	type Vertext struct {
		X int
		Y string
	}

	v := Vertext{1, "Hello"}

	fmt.Println(v)
	fmt.Println(v.X)
}

// zero values
func zeroValues() {
	var a int
	var b string
	var c bool
	d := make([]int, 0, 0)
	type Vertext struct {
	}
	// nilはポインタ、interface、slice、map、channel、functionのデフォルト値
	var e []int

	fmt.Println(a, b, c, d, e)
}

// Array practice
func array() {
	// Goでは配列は固定長、入る要素は同じ型でなければならないしプリみティブ型以外は使えない。
	var a = [2]string{"Hello", "World"}
	
	fmt.Println(a)
}

// slice practice
func slice() {
	var strSlice []string = make([]string, 3)

	strSlice[0] = "Hello"
	strSlice[1] = "World"
	strSlice[2] = "!"

	fmt.Println(strSlice)

	intArr := [4]int{1, 2, 3, 4}
	var intSlice []int = intArr[1: 3]
	intSlice[0] = 5
	fmt.Println(intSlice)

	hoge := []struct {
		i int
		b bool
		s string
	} {
		{1, true, "Hello"},
		{2, false, "World"},
		{3, true, "!"},
	}

	fmt.Println(hoge)
}

// tic tac toe
func ticTacToe() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"

	fmt.Println(board)
}

// append practice
func appendPractice() {
	var s []string
	s = append(s, "Hello")
	s = append(s, "World")

	fmt.Println(s[0], s[1])
}

// range practice for slice _を使う
func rangePractice() {
	a := []int{1, 2, 3, 4, 5}

	for i, v := range a {
		fmt.Println(i, v)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	for v := range a {
		fmt.Println(v)
	}
}

// map practice
func mapPractice() {
	m := make(map[string]int)
	
	m["hoge"] = 1
	m["fuga"] = 2
	m["foo"] = 3

	fmt.Println(m)

	delete (m, "hoge")

	fmt.Println(m)
}

// closure practice
func closruePractice() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	// slice()
	// mapPractice()
	// structPractice()
	// zeroValues()
	// ticTacToe()
	// appendPractice()
	// rangePractice()
	// var closure = closruePractice()
	// fmt.Println(closure()) // 1
	// fmt.Println(closure()) // 2
}