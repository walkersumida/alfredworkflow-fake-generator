package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/walkersumida/jaker"
)

// Item is Alfred's item struct.
type Item struct {
	Type     string `json:"type"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

// Menu is Alfred's menu struct.
type Menu struct {
	Items []Item `json:"items"`
}

type Row struct {
	Title string
	Value string
}

func outputFormat(items []Item) {
	var menu Menu
	menu.Items = items

	menuJSON, _ := json.Marshal(menu)
	fmt.Println(string(menuJSON))
}

func main() {
	var items []Item
	var rows []Row
	var item Item
	item.Icon = "./icon.png"

	flag.Parse()
	fakeType := flag.Arg(0)
	fakeType = strings.ToLower(fakeType)

	if fakeType == "text" {
		baseText := flag.Arg(1)
		if baseText == "" {
			item.Title = "Please enter base strings."
			items = append(items, item)
			outputFormat(items)
			return
		}

		strTextSize := flag.Arg(2)
		if strTextSize == "" {
			item.Title = "Please enter a text size."
			items = append(items, item)
			outputFormat(items)
			return
		}

		textSize, _ := strconv.Atoi(strTextSize)
		value:= jaker.Text(baseText, textSize)

		item.Title = "Text"
		item.Subtitle = value
		item.Arg = value

		items = append(items, item)
		outputFormat(items)
		return
	}

	if fakeType == "uuid" {
		value := jaker.Uuid

		item.Title = "Uuid"
		item.Subtitle = value
		item.Arg = value

		items = append(items, item)
		outputFormat(items)
		return
	}

	profile := jaker.Profile

	rows = append(rows,
		Row{ Title: "Name", Value: profile.JaKanjiFullName },
		Row{ Title: "Email", Value: profile.Email },
		Row{ Title: "Name + Email", Value: profile.JaKanjiFullName + "," + profile.Email },
		Row{ Title: "Company + Name + Email", Value: profile.JaCompany + "," + profile.JaKanjiFullName + "," + profile.Email },
	)

	for _, row := range rows {
		item.Title = row.Title
		item.Subtitle = row.Value
		item.Arg = row.Value

		items = append(items, item)
	}

	outputFormat(items)
}
