package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

// We are using Template from html/template package.
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	home := Page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input: `<script type="text/javascript">
				alert("No-No-No-No");
				</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)

	if err != nil {
		log.Fatalln(err)
	}
}
