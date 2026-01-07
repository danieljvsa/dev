package main

import (
	"log"
	"path/filepath"
	"time"

	"yourproject/sources"
	"yourproject/storage"
)

func main() {
	date := time.Now().Format("2006-01-02")
	path := filepath.Join("data", "france", date+".jsonl")

	articles, err := sources.FetchRSS(
		"https://www.gouvernement.fr/rss",
		"French Government",
		"France",
		"Politics",
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range articles {
		err := storage.SaveJSONL(path, article)
		if err != nil {
			log.Println("save error:", err)
		}
	}
}
