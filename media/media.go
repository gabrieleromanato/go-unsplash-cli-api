package media

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Image struct {
	Filename string
	Data     []byte
}

func DownloadImage(url string, filename string, c chan Image, logger *log.Logger) (Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Image{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Image{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Image{}, err
	}
	logger.Println("Downloaded", url)
	c <- Image{Filename: filename, Data: body}
	return Image{Filename: filename, Data: body}, nil

}

func SaveImage(image Image, logger *log.Logger) error {
	err := os.WriteFile(image.Filename, image.Data, 0755)
	if err != nil {
		fmt.Println("Error saving image", image.Filename)
		return err
	}
	logger.Println("Image", image.Filename, "saved")
	return nil
}
