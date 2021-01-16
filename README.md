# CLI RSS feed reader
> https://www.codementor.io/projects/rss-feed-reader-in-terminal-atx32jp82q

## How to use

```bash
# Prints out only the last news for each feed
go run main.go --limit 1 https://www.nasa.gov/rss/dyn/breaking_news.rss https://www.repubblica.it/rss/homepage/rss2.0.xml

# Prints out all the news
go run main.go https://www.nasa.gov/rss/dyn/breaking_news.rss

# Prints out at most 10 news for each feed
go run main.go --limit 10 https://www.nasa.gov/rss/dyn/breaking_news.rss https://www.repubblica.it/rss/homepage/rss2.0.xml
```
