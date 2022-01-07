package main

import (
	"html/template"
	"os"
)

type User struct {
	Name        string
	Bio         string
	Age         int
	ContactInfo map[string]contactInfo
}

type contactInfo struct {
	Phone string
	Email string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Dalton",
		Bio:  "<script>alert(\"XSS attempt\");</alert>",
		Age:  99,
		ContactInfo: map[string]contactInfo{
			"Work": contactInfo{
				Phone: "123-456-0000",
				Email: "workExample@example.com",
			},
			"Home": contactInfo{
				Phone: "123-456-1111",
				Email: "homeExample@example.com",
			},
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
