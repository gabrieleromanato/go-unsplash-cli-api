package utils

import (
	"errors"
	"math/rand"
	"time"
)

func RandomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomQuery(queries []string) string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	i := r.Intn(len(queries) - 1)
	return queries[i]
}

func GetURLsFromResponse(results map[string]interface{}) ([]string, error) {
	var urls []string
	if results["results"] == nil {
		return urls, errors.New("No results found")
	}
	for _, result := range results["results"].([]interface{}) {
		urls = append(urls, result.(map[string]interface{})["urls"].(map[string]interface{})["raw"].(string))
	}
	return urls, nil
}
