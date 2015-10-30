package rss

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Enclosure		Enclosure `xml:"enclosure"`
}
type Enclosure struct {
	Url					string `xml:"url,attr"`
	Type				string `xml:"type,attr"`
}
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

func ResponseToRss(r *http.Response) (Rss, error) {
	rss := Rss{}
	XMLdata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return rss, err
	}
	buffer := bytes.NewBuffer(XMLdata)
	decoded := xml.NewDecoder(buffer)
	err = decoded.Decode(&rss)
	return rss, err
}
