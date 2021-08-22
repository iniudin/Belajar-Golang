package main

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
)

func getWeather(city string) {
	url := "https://www.google.com/search?q=cuaca+" + city
	soup.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0")
	res, err := soup.Get(url)

	if err != nil {
		log.Fatal(err.Error())
	}

	// html = Soup(resp.text, "lxml")
	// table = html.find(id="wob_wc")
	// loc = table.find(id="wob_loc").text
	// datetime = table.find(id="wob_dts").text
	// desc = table.find(id="wob_dc").text
	// temp = table.find(id="wob_tm").text
	// precipitation = table.find(id="wob_pp").text
	// humidity = table.find(id="wob_hm").text
	// wind = table.find(id="wob_ws").text

	doc := soup.HTMLParse(res)
	location := doc.Find("div", "id", "wob_loc").Text()
	datetime := doc.Find("div", "id", "wob_dts").Text()
	description := doc.Find("span", "id", "wob_dc").Text()
	// temperature := doc.Find("span")
	fmt.Println(location, datetime, description)

}

func main() {
	getWeather("pacet")
}
