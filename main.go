package main

import (
	"flag"
	"fmt"
	"gabrieleromanato/unsplash/api"
	"gabrieleromanato/unsplash/media"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(1)
	}
	defer logFile.Close()
	logger := log.New(logFile, "unsplash ", log.Lshortfile|log.LstdFlags)

	queryPtr := flag.String("query", "cats", "Search query for Unsplash API")

	flag.Parse()

	query := *queryPtr
	imageURLs, err := api.SearchImages(query)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	downloadChan := make(chan media.Image)
	for i, url := range imageURLs {
		n := i + 1
		go media.DownloadImage(url, fmt.Sprintf("image-%d.jpg", n), downloadChan, logger)
	}
	for i := 0; i < len(imageURLs); i++ {
		image := <-downloadChan
		err := media.SaveImage(image, logger)
		if err != nil {
			logger.Println(err)
		}
	}
	os.Exit(0)
}
