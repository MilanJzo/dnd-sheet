package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// https://srd.dndtools.org/srd/magic/spells/spells/spellsAllCore.html

type Spell struct {
	Name            string
	Level           string
	Users           []string
	School          string
	Components      string
	Range           string
	Duration        string
	Area            string
	Effect          string
	SavingThrow     string
	CastingTime     string
	SpellResistance string
	Description     string
	Link            string
}

// "Ranger", "Cleric (Plant Domain)"

func main() {
	reqest, err := http.Get("https://srd.dndtools.org/srd/magic/spells/spells/spellsAllCore.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(reqest.Body)
	if err != nil {
		log.Fatal(err)
	}

	spells := make([]Spell, 0)

	selector := "body > table > tbody > tr > td > h6"
	doc.Find(selector).NextAllFiltered("h6").Add(selector).Each(func(i int, s *goquery.Selection) {
		spell := Spell{
			Users: make([]string, 0),
		}
		spell.Name = s.Text()
		link, _ := s.Children().First().Attr("id")
		spell.Link = "https://srd.dndtools.org/srd/magic/spells/spells/spellsAllCore.html#" + link

		s.NextUntil("h6").Each(func(i int, s *goquery.Selection) {
			text := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(s.Text()), "\n", " "), "  ", " ")

			if i == 1 {
				spell.School = text
			} else if i == 2 {
				spell.Level = strings.ReplaceAll(text, "Level: ", "")
			} else if i == 3 {
				spell.Components = strings.ReplaceAll(text, "Components: ", "")
			} else if i == 4 {
				spell.CastingTime = strings.ReplaceAll(text, "Casting Time: ", "")
			} else if i == 5 {
				spell.Range = strings.ReplaceAll(text, "Range: ", "")
			} else if strings.HasPrefix(text, "Effect:") {
				spell.Effect = strings.ReplaceAll(text, "Effect: ", "")
			} else if strings.HasPrefix(text, "Area:") {
				spell.Area = strings.ReplaceAll(text, "Area: ", "")
			} else if strings.HasPrefix(text, "Duration:") {
				spell.Duration = strings.ReplaceAll(text, "Duration: ", "")
			} else if strings.HasPrefix(text, "Saving Throw:") {
				spell.SavingThrow = strings.ReplaceAll(text, "Saving Throw: ", "")
			} else if strings.HasPrefix(text, "Spell Resistance:") {
				spell.SpellResistance = strings.ReplaceAll(text, "Spell Resistance: ", "")
			} else if !s.HasClass("stat-block") && !strings.Contains(text, "Target, Effect, or Area") && text != "" {
				spell.Description += text + "\n"
			}
		})

		spell.Description = strings.Trim(spell.Description, "\n")

		spells = append(spells, spell)
	})

	bardSpells := doc.Find("body > table > tbody > tr > td > table:nth-child(5) > tbody > tr > td:nth-child(1) > h4:nth-child(1)").NextUntil("h4")
	bardSpellsString := strings.Join(bardSpells.Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	}), "\n")

	druidSpells := doc.Find("body > table > tbody > tr > td > table:nth-child(5) > tbody > tr > td:nth-child(2) > h4:nth-child(1)").NextUntil("h4")
	druidSpellsString := strings.Join(druidSpells.Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	}), "\n")

	clericSpells := doc.Find("body > table > tbody > tr > td > table:nth-child(5) > tbody > tr > td:nth-child(1) > h4:nth-child(173)").NextUntil("h4")
	clericSpellsString := strings.Join(clericSpells.Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	}), "\n")

	paladinSpells := doc.Find("body > table > tbody > tr > td > table:nth-child(5) > tbody > tr > td:nth-child(2) > h4:nth-child(181)").NextUntil("h4")
	paladinSpellsString := strings.Join(paladinSpells.Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	}), "\n")

	rangerSpells := doc.Find("body > table > tbody > tr > td > table:nth-child(5) > tbody > tr > td:nth-child(2) > h4:nth-child(230)").NextUntil("h4")
	rangerSpellsString := strings.Join(rangerSpells.Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	}), "\n")

	wizardSpells := doc.Find("body > table > tbody > tr > td > table:nth-child(5) > tbody > tr > td:nth-child(2) > h4:nth-child(286)").NextUntil("h4")
	wizardSpellsString := strings.Join(wizardSpells.Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	}), "\n")

	for i, spell := range spells {
		if strings.Contains(wizardSpellsString, spell.Name) {
			spells[i].Users = append(spells[i].Users, "Wizard")
		}
		if strings.Contains(bardSpellsString, spell.Name) {
			spells[i].Users = append(spells[i].Users, "Bard")
		}
		if strings.Contains(druidSpellsString, spell.Name) {
			spells[i].Users = append(spells[i].Users, "Druid")
		}
		if strings.Contains(clericSpellsString, spell.Name) {
			spells[i].Users = append(spells[i].Users, "Cleric")
		}
		if strings.Contains(paladinSpellsString, spell.Name) {
			spells[i].Users = append(spells[i].Users, "Paladin")
		}
		if strings.Contains(rangerSpellsString, spell.Name) {
			spells[i].Users = append(spells[i].Users, "Ranger")
		}
	}

	formatted, err := json.MarshalIndent(spells, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("formatted_spells.json", formatted, 0644)
	if err != nil {
		log.Fatal(err)
	}

	clean, err := json.Marshal(spells)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("spells.json", clean, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
