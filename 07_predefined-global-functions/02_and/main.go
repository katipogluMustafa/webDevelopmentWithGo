package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	First string
	Age   int
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	u1 := user{
		First: "Mustafa",
		Age:   20,
		Admin: false,
	}

	u2 := user{
		First: "Fatih",
		Age:   26,
		Admin: true,
	}

	u3 := user{
		First: "Ahmet",
		Age:   25,
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
