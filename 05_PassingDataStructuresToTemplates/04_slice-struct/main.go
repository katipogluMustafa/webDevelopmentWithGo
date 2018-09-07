package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type member struct {
	First string
	Age   int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	mustafa := member{
		First: "Mustafa",
		Age:   20,
	}
	ahmet := member{
		First: "Ahmet",
		Age:   24,
	}
	fatih := member{
		First: "Fatih",
		Age:   20,
	}
	mert := member{
		First: "Mert",
		Age:   20,
	}
	alperen := member{
		First: "Alperen",
		Age:   21,
	}

	members := []member{mustafa, ahmet, fatih, mert, alperen}

	err := tpl.Execute(os.Stdout, members)

	if err != nil {
		log.Fatalln(err)
	}
}
