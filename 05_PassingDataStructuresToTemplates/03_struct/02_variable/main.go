package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type members struct {
	First string
	Last  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	mustafa := members{
		First: "Mustafa",
		Last:  "Katipoğlu",
	}

	err := tpl.Execute(os.Stdout, mustafa)

	if err != nil {
		log.Fatalln(err)
	}
}
