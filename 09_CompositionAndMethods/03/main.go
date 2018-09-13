package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"CS50", "Intro to Programming", "4"},
				{"BLM2622", "Electronic Circuits", "6"},
				{"BLM2021", "Low Level Programming", "3"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"CSCI-50", "Advanced Go", "5"},
				{"CSCI-190", "Advanced Web Programming with Go", "5"},
				{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)

	if err != nil {
		log.Fatalln(err)
	}
}
