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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Group     []member
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	mustafa := member{
		First: "Mustafa",
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
	ford := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}
	toyota := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	members := []member{mustafa, mert, alperen}
	cars := []car{ford, toyota}

	data := items{
		Group:     members,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
