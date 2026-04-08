package main

import (
	"fmt"
	"strings"
)

func variables() {
	const name string = "nico"

	name2 := "hyuk" // var name2 string = "hyuk"
	// name2 = 2 불가능, 타입 한번 정해지면 고정

	x, _ := 10, 20 // _ 무시가능

	fmt.Println(name, name2, x)
}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, uppercase string) { // 반환변수 미리 만들어둬서 := 안됨
	defer fmt.Println("I am done")// return 할때 실행할 내용
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // 어떤 변수 반환할 지 정해놔서 return만 하면 됨
}

func repeat(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	sum := 0
	for index, number := range numbers{ // index 반환, range 많이 사용
		fmt.Println(index, number)
		sum += number
	}

	for i := 0; i < len(numbers); i++ { // C 스타일
		fmt.Println(i, numbers[i])
	}

	return sum
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge  < 20 { // variable expression, if 시 변수 생성
		return false
	} else {
		fmt.Println(age)
	}
	return true	
}

func canIDrinkSwitch(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}

	switch {
	case age < 10:
		return false

	case age > 18:
		return true
	}

	return true
}

func point() {
	a := 2
	b := a // cpy value
	a = 10
	fmt.Println(a, b)
	fmt.Println(&a, &b) // access to memory

	c := &a
	fmt.Println(&a, c)
	fmt.Println(a, *c)

	*c = 5
	fmt.Println(a)	
}

func array() {
	names := [5]string{"a", "b", "c"} // []에 길이 없으면 계속 추가 가능
	names[4] = "hyuk"
	names[3] = "asd"
	fmt.Println(names)

	nums := []int{3, 1, 4}
	nums = append(nums, 1) // 변형한 리스트를 다시 입력
	fmt.Println(nums)
}

func maps() {
	nico := map[string]string{"name":"nico", "age":"12"}

	fmt.Println(nico)

	for k, v := range nico {
		fmt.Println(k, v)
	}
}

type person struct { // 생성자 없음, 직접 만들어야함
	name string
	age int
	favFood []string
}

// func main() {
// 	hyuk := person{name:"hyuk", age:25, favFood: []string{"a", "b"}} // attribute 안붙여도 되는데, 가독성상 쓰는거 선호
// 	fmt.Println(hyuk)
// }