package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Dog  string
	Age  int
}

func main() {
	template, err := template.ParseFiles("home.gohtml")

	if err != nil {
		panic(err)
	}

	data := User{Name: "Marvin Collins", Dog: "Derick", Age: 26}

	err = template.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "Gloria Glo"
	data.Dog = "John"
	data.Age = 24
	err = template.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
