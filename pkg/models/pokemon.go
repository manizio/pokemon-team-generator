package models

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Pokemon struct {
	ID int `json:"id"`
	Species struct {
		Name string `json:"name"`
	} `json:"species"`
	Sprite struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

func (p *Pokemon) FormatName() {
	p.Species.Name = strings.Replace(p.Species.Name, "-", " ", -1)
	p.Species.Name = cases.Title(language.English, cases.Compact).String(p.Species.Name)
}
