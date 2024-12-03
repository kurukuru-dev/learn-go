package main

import (
	"fmt"
	// "errors"
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
	c := *b

	fmt.Println(a, b, c)
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

// method practice
type Vertext struct {
	s1 string
	s2 string
}

func (v Vertext) Print(s3 string) {
	fmt.Println(v.s1, v.s2, s3)
}

// pointer receiver
type Vertex2 struct {
	name string
	age int
}

func (v Vertex2) Print() {
	v.age++
}

func (v *Vertex2) Print2() {
	v.age++
	(*v).age++
}

// interface practice
type I interface {
	M()
}

type T struct {
	S string
	B bool
}

func (t *T) M() {
	fmt.Println(t.S)
}

// type assertion
type I2 interface {}

func typeAssertion(i I2) {
	s, ok := i.(string)
	fmt.Println(s, ok)
}

// switch with type
func do(i interface{}) {
	switch v := i.(type) {
		case int:
			fmt.Println("int", v)
		case string:
			fmt.Println("string", v)
		default:
			fmt.Println("unknown")
	}
}

// stringer interface
type Person struct {
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

// error interface
type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func validate(i int) error {
	if i < 0 {
		return fmt.Errorf("invalid value: %v", i) // fmt.Errorfはerrorを返すだけで、出力はしない
	}
	return nil
}

// go routine
func fibonacci(c, f chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <- f:
			fmt.Println("finish")
			return
		}
	}
}

// baffer channel
func bufferChannel() {
	c := make (chan int, 1)
	c <- 1
	c <- 2 // panic
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func main() {
	// slice()
	// mapPractice()
	// structPractice()
	// zeroValues()
	// ticTacToe()
	// appendPractice()
	// pointerPractice()
	// rangePractice()
	// var closure = closruePractice()
	// fmt.Println(closure()) // 1
	// fmt.Println(closure()) // 2

	// v := Vertext{"hello", "world"}
	// v.Print("!")

	// v := Vertex2{"komori", 1}
	// v.Print()
	// fmt.Println(v.age) // {komori 1}
	// v.Print2()
	// fmt.Println(v.age) // {komori 2}

	// i := &T{"Hello", true}
	// i.M()

	// typeAssertion("Hello")
	// typeAssertion(1) // panic

	// do(1)
	// do("Hello")
	// do(true)

	// p := Person{"komori"}
	// fmt.Println(p)

	// e := &MyError{"Error"}
	// fmt.Println(e)
	// var ErrNotUser = errors.New("user not found")
	// fmt.Println(ErrNotUser)
	// e := validate(-1)
	// fmt.Println(e)

	// c := make(chan int)
	// f := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	f <- 0
	// }()
	// fibonacci(c, f)

	// bufferChannel()
}