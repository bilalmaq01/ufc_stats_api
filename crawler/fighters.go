package crawler

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func FighterCrawler(c *colly.Collector) {

	c.OnHTML("tr.b-statistics__table-row", func(e *colly.HTMLElement) {
		//These are pointers because the website im scraping from ufcstats.com some fighters don't have all of these
		var nickname *string
		var height *string
		var weight *string
		var reach_in *string

		//Combine first and last name and make it a single name
		name := fmt.Sprintf("%s %s", e.ChildText("td:nth-child(1)"), e.ChildText("td:nth-child(2)"))

		fmt.Println()
		fmt.Println(name)

		//Not all fighters have a nickname
		n := e.ChildText("td:nth-child(3)")
		if len(n) == 0 {
			nickname = nil
		} else {
			nickname = &n
		}
		if nickname != nil {
			fmt.Println(n)
		}

		h := e.ChildText("td:nth-child(4)")
		if h == "--" {
			height = nil
		} else {
			height = &h
		}
		if height != nil {
			fmt.Println("Height:", h)
		}

		w := e.ChildText("td:nth-child(5)")
		if w == "--" {
			weight = nil
		} else {
			weight = &w
		}
		if weight != nil {
			fmt.Println("Weight Class:", w)
		}
		r := e.ChildText("td:nth-child(6)")

		if r == "--" {
			reach_in = nil
		} else {
			reach_in = &r
		}
		r = strings.TrimSuffix(r, "\"")
		if r != "--" && r != "" {
			int_r, err := strconv.ParseFloat(r, 8)
			if err != nil {
				log.Fatal(err)
			}
			if reach_in != nil {
				fmt.Println("Reach:", int_r)
			}
		}

		wins := e.ChildText("td:nth-child(8)")
		fmt.Println("Wins:", wins)
		losses := e.ChildText("td:nth-child(9)")
		fmt.Println("Losses:", losses)
		draws := e.ChildText("td:nth-child(10)")
		fmt.Println("Draws:", draws)

	})

	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		link := fmt.Sprintf("http://www.ufcstats.com/statistics/fighters?char=%c&page=all", letter)
		c.Visit(link)
		link = fmt.Sprintf("Visited %s", link)
		fmt.Println(link)
	}

}
