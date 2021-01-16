package main

import (
	"fmt"
	"testing"
)

// workaround for error because of the use of flag.Parse()
// https://stackoverflow.com/a/58192326/1420669
var _ = func() bool {
	testing.Init()
	return true
}()

func TestGetItemFrom(t *testing.T) {
}

func TestToHL(t *testing.T) {
	ret := toHL("foo")
	if ret != fmt.Sprintf("\033]8;;%[1]s\a%[1]s\033]8;;\a", "foo") {
		t.Error("Error!")
		return
	}
}

func TestExtract(t *testing.T) {
	i := `<title>2020 Tied for Warmest Year on Record, NASA Analysis Shows</title>
  <link>http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows</link>
  <description>Earth’s global average surface temperature in 2020 tied with 2016 as the warmest year on record, according to an analysis by NASA.</description>
  <enclosure url="http://www.nasa.gov/sites/default/files/styles/1x1_cardfeed/public/thumbnails/image/2020temp_print.jpg?itok=FCoD74pJ" length="173482" type="image/jpeg" />
  <guid isPermaLink="false">http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows</guid>
  <pubDate>Thu, 14 Jan 2021 09:09 EST</pubDate>
  <source url="http://www.nasa.gov/rss/dyn/breaking_news.rss">NASA Breaking News</source>
  <dc:identifier>467581</dc:identifier>`

	if extract("title", i) != "2020 Tied for Warmest Year on Record, NASA Analysis Shows" {
		t.Error("Title is wrong!")
		return
	}

	if extract("link", i) != "http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows" {
		t.Error("Link is wrong!")
		return
	}

	if extract("pubDate", i) != "Thu, 14 Jan 2021 09:09 EST" {
		t.Error("pubDate is wrong!")
		return
	}
}

func TestExtract2(t *testing.T) {
	i := `<title><![CDATA[2020 Tied for Warmest Year on Record, NASA Analysis Shows]]></title>
  <link><![CDATA[http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows]]></link>
  <description><![CDATA[Earth’s global average surface temperature in 2020 tied with 2016 as the warmest year on record, according to an analysis by NASA.]]></description>
  <enclosure url="http://www.nasa.gov/sites/default/files/styles/1x1_cardfeed/public/thumbnails/image/2020temp_print.jpg?itok=FCoD74pJ" length="173482" type="image/jpeg" />
  <guid isPermaLink="false">http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows</guid>
  <pubDate><![CDATA[Thu, 14 Jan 2021 09:09 EST]]></pubDate>
  <source url="http://www.nasa.gov/rss/dyn/breaking_news.rss">NASA Breaking News</source>
  <dc:identifier>467581</dc:identifier>`

	if extract("title", i) != "2020 Tied for Warmest Year on Record, NASA Analysis Shows" {
		t.Error("Title is wrong with CDATA!")
		return
	}

	if extract("link", i) != "http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows" {
		t.Error("Link is wrong with CDATA!")
		return
	}

	if extract("pubDate", i) != "Thu, 14 Jan 2021 09:09 EST" {
		t.Error("pubDate is wrong with CDATA!")
		return
	}
}

func TestGetItem(t *testing.T) {
	i := `<title>2020 Tied for Warmest Year on Record, NASA Analysis Shows</title>
  <link>http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows</link>
  <description>Earth’s global average surface temperature in 2020 tied with 2016 as the warmest year on record, according to an analysis by NASA.</description>
  <enclosure url="http://www.nasa.gov/sites/default/files/styles/1x1_cardfeed/public/thumbnails/image/2020temp_print.jpg?itok=FCoD74pJ" length="173482" type="image/jpeg" />
  <guid isPermaLink="false">http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows</guid>
  <pubDate>Thu, 14 Jan 2021 09:09 EST</pubDate>
  <source url="http://www.nasa.gov/rss/dyn/breaking_news.rss">NASA Breaking News</source>
  <dc:identifier>467581</dc:identifier>`
	it := getItemFrom(i)

	if it.Title != extract("title", i) {
		t.Error("Title is wrong")
	}

	if it.Descr != extract("description", i) {
		t.Error("Descr is wrong")
	}

	if it.Date != extract("pubDate", i) {
		t.Error("Date is wrong")
	}

	if it.Link != extract("link", i) {
		t.Error("Link is wrong")
	}

}

func TestParse(t *testing.T) {
	// courtesy of NASA
	xml := `<?xml version="1.0" encoding="utf-8" ?> <rss version="2.0" xml:base="http://www.nasa.gov/" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:media="http://search.yahoo.com/mrss/"> <channel>
   <title>NASA Breaking News</title>
   <description>A RSS news feed containing the latest NASA news articles and press releases.</description>
   <link>http://www.nasa.gov/</link>
   <atom:link rel="self" href="http://www.nasa.gov/rss/dyn/breaking_news.rss" />
   <language>en-us</language>
   <managingEditor>jim.wilson@nasa.gov</managingEditor>
   <webMaster>brian.dunbar@nasa.gov</webMaster>
   <docs>http://blogs.harvard.edu/tech/rss</docs>
  <item>
   <title>2020 Tied for Warmest Year on Record, NASA Analysis Shows</title>
   <link>http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows</link>
   <description>Earth’s global average surface temperature in 2020 tied with 2016 as the warmest year on record, according to an analysis by NASA.</description>
   <enclosure url="http://www.nasa.gov/sites/default/files/styles/1x1_cardfeed/public/thumbnails/image/2020temp_print.jpg?itok=FCoD74pJ" length="173482" type="image/jpeg" />
   <guid isPermaLink="false">http://www.nasa.gov/press-release/2020-tied-for-warmest-year-on-record-nasa-analysis-shows</guid>
   <pubDate>Thu, 14 Jan 2021 09:09 EST</pubDate>
   <source url="http://www.nasa.gov/rss/dyn/breaking_news.rss">NASA Breaking News</source>
   <dc:identifier>467581</dc:identifier>
  </item>
   <item> <title>NASA TV to Air Hot Fire Test of Rocket Core Stage for Artemis Moon Missions</title>
   <link>http://www.nasa.gov/press-release/nasa-tv-to-air-hot-fire-test-of-rocket-core-stage-for-artemis-moon-missions</link>
   <description>NASA is targeting a two-hour test window that opens at 5 p.m. EST Saturday, Jan. 16, for the hot fire test of NASA’s Space Launch System (SLS) rocket core stage at the agency’s Stennis Space Center near Bay St. Louis, Mississippi.</description>
   <enclosure url="http://www.nasa.gov/sites/default/files/styles/1x1_cardfeed/public/thumbnails/image/ssc_photo_4603_large.jpg?itok=PDsLpbb5" length="243520" type="image/jpeg" />
   <guid isPermaLink="false">http://www.nasa.gov/press-release/nasa-tv-to-air-hot-fire-test-of-rocket-core-stage-for-artemis-moon-missions</guid>
   <pubDate>Wed, 13 Jan 2021 10:11 EST</pubDate>
   <source url="http://www.nasa.gov/rss/dyn/breaking_news.rss">NASA Breaking News</source>
   <dc:identifier>467551</dc:identifier>
  </item>`
	parsed := parse(xml, 1)

	if parsed.Title != "NASA Breaking News" {
		t.Error("Wrong title!")
	}

	if parsed.Link != toHL("http://www.nasa.gov/") {
		t.Error("Wrong link!")
	}

	if parsed.Descr != "A RSS news feed containing the latest NASA news articles and press releases." {
		t.Error("Wrong description!")
	}

	if len(parsed.Items) != 1 {
		t.Error("Wrong items count 1")
	}

	parsed = parse(xml, -1)

	if len(parsed.Items) != 2 {
		t.Error("Wrong items count 2")
	}

	if parsed.Items[0].Title != "2020 Tied for Warmest Year on Record, NASA Analysis Shows" {
		t.Error("Items[0] is wrong!")
	}

	if parsed.Items[1].Title != "NASA TV to Air Hot Fire Test of Rocket Core Stage for Artemis Moon Missions" {
		t.Error("Items[1] is wrong")
	}

}

func TestMain(t *testing.T) {
	// TODO
}
