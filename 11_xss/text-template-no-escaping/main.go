package main

import (
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

// We are using Template from text/template
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	home := Page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped with text/template",
		Input: `<script type="text/javascript">
				alert("No-No-No");
				</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)

	if err != nil {
		log.Fatalln(err)
	}
}
