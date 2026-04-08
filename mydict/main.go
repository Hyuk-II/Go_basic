package main

import (
	"fmt"

	"github.com/Hyuk-II/Go_basic/mydict"
)


func main() {
	dictionary := mydict.Dictionary{}
	// definition, err := dictionary.Search("first")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// word := "Hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// hello, err := dictionary.Search(word)
	// fmt.Println("found hello definition:", hello)
	
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// baseWord := "Hello"
	// dictionary.Add(baseWord, "First")
	// err := dictionary.Update("asddas", "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)

	// fmt.Println(word)

	baseWord := "Hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(word)
}