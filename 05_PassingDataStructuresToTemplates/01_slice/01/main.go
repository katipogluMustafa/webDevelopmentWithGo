package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	members := []string{"Mustafa", "Ahmet", "Yusuf", "Alperen", "Fatih", "Mert"}
	err := tpl.Execute(os.Stdout, members)

	if err != nil {
		log.Fatalln(err)
	}
}
