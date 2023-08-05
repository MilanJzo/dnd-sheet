package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Ability struct {
	Name        string
	Description string
	Link        string
}

type Class struct {
	Name      string
	Abilities []Ability
}

func main() {
	req, err := http.Get("https://www.d20srd.org/indexes/classes.htm")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	classesLinks := doc.Find("body > ul:nth-child(3) > li:nth-child(1) > ul").Find("a").Map(func(i int, s *goquery.Selection) string {
		link, _ := s.Attr("href")
		return "https://www.d20srd.org" + link
	})

	//filter out all links that include a #
	n := 0
	for _, link := range classesLinks {
		if !strings.Contains(link, "#") {
			classesLinks[n] = link
			n++
		}
	}
	classesLinks = classesLinks[:n]
	// log.Println(classesLinks)

	// classesLinks = classesLinks[:1]

	classes := make([]Class, 0)

	for i, class := range classesLinks {
		log.Printf("%d / %d %s", i+1, len(classesLinks), class)

		req, err := http.Get(class)
		if err != nil {
			log.Fatal(err)
		}

		doc, err := goquery.NewDocumentFromReader(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		abilities := make([]Ability, 0)

		doc.Find("h4:nth-of-type(2)").NextUntil("h4").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}

			text := strings.Trim(s.Text(), " \n\t")
			if i%2 != 0 {
				abilities = append(abilities, Ability{
					Name: text,
					Link: class+ "#" + s.AttrOr("id", ""),
				})
			} else {
				abilities[len(abilities)-1].Description = text
			}
		})

		classes = append(classes, Class{
			Name:      class,
			Abilities: abilities,
		})
	}

	log.Println(classes)

	j, err := json.Marshal(classes)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("special_abilities.json", j, 0644)

	formated, err := json.MarshalIndent(classes, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("formated_special_abilities.json", formated, 0644)
}
