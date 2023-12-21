package main

import "fmt"

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func addSuffix(suffix string) func(string) string {
	return func(s string) string {
		return s + suffix
	}
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	p1 := Person{
		name: "小王子",
		age:  18,
	}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?

	// 闭包
	add := addSuffix(".jpg")
	s := add("test")
	fmt.Println(s) // test.jpg

	// defer机制
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
