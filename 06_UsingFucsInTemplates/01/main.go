package main

import (
	"log"
	"os"
	"strings"
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

/*
 * Create a FuncMap to register functions.
 */
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

//.New() method is used because if you use Parse(str string) method instead of ParseFiles, you need to name the template in order to execute it with tpl.ExecuteTemplate()
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
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

	data := struct {
		Group     []member
		Transport []car
	}{
		members,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
