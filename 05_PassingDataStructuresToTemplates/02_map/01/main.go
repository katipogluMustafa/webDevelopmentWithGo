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
	members := map[string]string{
		"asdas": " sdasda",
		"asdad": " sdasda",
		"asdaf": " sdasda",
		"asdag": " sdasda",
		"asdah": " sdasda",
	}

	err := tpl.Execute(os.Stdout, members)
	if err != nil {
		log.Fatalln(err)
	}
}
