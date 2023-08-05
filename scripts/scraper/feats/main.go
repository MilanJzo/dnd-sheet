package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Feat struct {
	Name         string
	Description  string
	Prerequisite string
	Benefit      string
	Normal       string
	Special      string
	Link         string
}

func getLinks() {
	c := colly.NewCollector()
	links := make([]string, 0)

	c.OnHTML("div.mw-category-group > ul > li > a", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		link := e.Attr("href")
		if !strings.Contains(link, "Category:") {
			links = append(links, "https://www.dandwiki.com"+link)
			fmt.Println("Visiting", "https://www.dandwiki.com"+link)
		}
	})

	c.Visit("https://www.dandwiki.com/wiki/3.5e_Homebrew_Feats")

	links = links[1:]

	f, err := os.Create("links.json")
	if err != nil {
		log.Fatal(err)
	}

	j, err := json.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(j)
	defer f.Close()
}

func getFeats() {
	linksAsBytes, err := os.ReadFile("links.json")
	if err != nil {
		log.Fatal(err)
	}

	feats := make([]Feat, 0)

	var links []string
	err = json.Unmarshal(linksAsBytes, &links)
	if err != nil {
		log.Fatal(err)
	}

	// links = links[:5]

	for _, link := range links {
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		divs := doc.Find(".mw-parser-output > div > div")
		altDivs := doc.Find(".mw-parser-output div")

		feat := Feat{}
		feat.Link = link
		feat.Description = divs.First().Text()
		feat.Name = doc.Find(".mw-headline").Text()

		divs.Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			if strings.Contains(text, "Prerequisite") {
				feat.Prerequisite = strings.Replace(text, "Prerequisite: ", "", 1)
			} else if strings.Contains(text, "Benefit") {
				feat.Benefit = strings.Replace(text, "Benefit: ", "", 1)
			} else if strings.Contains(text, "Normal") {
				feat.Normal = strings.Replace(text, "Normal: ", "", 1)
			} else if strings.Contains(text, "Special") {
				feat.Special = strings.Replace(text, "Special: ", "", 1)
			}
		})

		altDivs.Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			if feat.Prerequisite == "" && strings.Contains(text, "Prerequisite") {
				feat.Prerequisite = strings.Replace(text, "Prerequisite: ", "", 1)
			} else if feat.Benefit == "" && strings.Contains(text, "Benefit") {
				feat.Benefit = strings.Replace(text, "Benefit: ", "", 1)
			} else if feat.Normal == "" && strings.Contains(text, "Normal") {
				feat.Normal = strings.Replace(text, "Normal: ", "", 1)
			} else if feat.Special == "" && strings.Contains(text, "Special") {
				feat.Special = strings.Replace(text, "Special: ", "", 1)
			}
		})

		feats = append(feats, feat)
		log.Println(len(feats), "/", len(links))
	}

	j, err := json.Marshal(feats)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("feats.json", j, 0644)

	formatted, err := json.MarshalIndent(feats, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("formatted_feats.json", formatted, 0644)
}

func main() {
	getLinks()
	getFeats()
}
