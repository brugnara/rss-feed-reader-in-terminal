package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type item struct {
	Title string
	Link  string
	Descr string
	Date  string
}

type data struct {
	Items []item
	Title string
	Descr string
	Link  string
}

var tpl *template.Template
var flagLimit int

func init() {
	fns := template.FuncMap{
		"incr": func(a, i int) int {
			return a + i
		},
	}
	tpl = template.Must(
		template.New("").Funcs(fns).ParseGlob("./tpls/*"),
	)
	if len(os.Args) < 2 {
		log.Fatalln("Missing url/s.")
	}
	//
	flag.IntVar(&flagLimit, "limit", -1, "Limit results to this number")
	flag.Parse()
}

func main() {
	for _, wantedURL := range os.Args[1:] {
		if !strings.HasPrefix(
			wantedURL, "http://") && !strings.HasPrefix(
			wantedURL, "https://") {
			continue
		}

		file, err := getFeed(wantedURL)
		if err != nil {
			log.Println(err)
			continue
		}

		parsed := parse(file, flagLimit)
		if parsed != nil {
			tpl.ExecuteTemplate(os.Stdout, "items.gohtml", parsed)
		}
	}
}

func extract(field, from string) (ret string) {
	tx := regexp.MustCompile(
		fmt.Sprintf("<%s>(.*?)</.*?>", field),
	).FindStringSubmatch(from)
	if len(tx) == 2 {
		ret = tx[1]
	}
	ret = strings.Replace(ret, "<![CDATA[", "", -1)
	ret = strings.Replace(ret, "]]>", "", -1)
	return
}

func toHL(s string) string {
	return fmt.Sprintf("\033]8;;%[1]s\a%[1]s\033]8;;\a", s)
}

func getItemFrom(s string) item {
	return item{
		Title: extract("title", s),
		Link:  extract("link", s),
		Descr: extract("description", s),
		Date:  extract("pubDate", s),
	}
}

func parse(feed string, flagLimit int) *data {
	feed = strings.ReplaceAll(feed, "\n", "")
	// header
	var hTitle string
	var hDescr string
	var hLink string
	// extracts the header in the XML
	hd := regexp.MustCompile(
		`<channel>(.*?)<item>`,
	).FindStringSubmatch(feed)

	if len(hd) == 2 {
		hTitle = extract("title", hd[1])
		hDescr = extract("description", hd[1])
		hLink = extract("link", hd[1])
		if hLink != "" {
			hLink = toHL(hLink)
		}
	}

	// items
	ix := regexp.MustCompile(
		`<item>(.*?)<\/item>`,
	).FindAllStringSubmatch(
		feed, -1)

	limit := len(ix)
	if flagLimit != -1 {
		limit = flagLimit
	}

	items := make([]item, limit)

	for i, it := range ix[0:limit] {
		items[i] = getItemFrom(it[1])
	}
	//
	return &data{
		Link:  hLink,
		Descr: hDescr,
		Title: hTitle,
		Items: items,
	}
}

func getFeed(uri string) (string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}
