# Concurrent News Scraper

This project demonstrates concurrent web scraping from two news sites using Golang. The scraper collects article links based on a user-provided keyword, using the [Colly](https://github.com/gocolly/colly) library, and saves the results in a JSON file.

## Features

- Scrapes article links from **Times of India** and **NDTV** concurrently.
- Uses Go’s goroutines and WaitGroups for parallelism.
- Stores scraped links in a JSON file (`articles.json`).
- Thread-safe data collection with mutex.
- Customizable via user input keyword.

## Requirements

- Go 1.18+
- [Colly](https://github.com/gocolly/colly) package

Install Colly:
```bash
go get -u github.com/gocolly/colly
```

## Usage

1. Clone the repository:
    ```bash
    git clone https://github.com/SafwanAhmed08/golang.git
    cd golang
    ```

2. Run the scraper:
    ```bash
    go run ConcurrectNewsScraper.go
    ```

3. Enter the keyword when prompted (e.g., `technology`, `sports`, etc.).

4. The results are saved in `articles.json`.

## Example

```bash
$ go run ConcurrectNewsScraper.go
Enter the name of the News link page: technology
Visiting https://timesofindia.indiatimes.com/topic/technology
Visiting https://www.ndtv.com/search?searchtext=technology
Scraping finished.
Scraping finished.
```

## Output

The scraped article links are saved in `articles.json`:
```json
[
  {
    "links": [
      "https://timesofindia.indiatimes.com/...",
      ...
    ]
  },
  {
    "links": [
      "https://www.ndtv.com/...",
      ...
    ]
  }
]
```

## Notes

- The CSS selectors used (`.crmK8` for Times of India and `.src_tab-cnt` for NDTV) may need updating if the websites change their layouts.
- Error handling is implemented to exit on failed requests.
