package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>image>caption"`
	Time []string `xml:"url>lastmod"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Time string
	Locations string
}

func main() {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://indianexpress.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
		// fmt.Println(n.Titles)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Time[idx], n.Locations[idx]}
			fmt.Println(idx)
		}
	}
	
	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Time)
		fmt.Println("\n", data.Locations)
	}

}