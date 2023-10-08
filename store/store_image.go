package store

import (
	. "WikipediaImage/parse"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type StoreFile string

type Store struct {
}

var rootDir = "store_image"

func (store *Store) StoreImage(imageResults []ImageResult) error {
	pwd, _ := os.Getwd()

	for _, result := range imageResults {
		dir, _ := filepath.Abs(filepath.Join(pwd, rootDir, result.Date))
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
		thumbImageUrl := result.ThumbImageUrl
		thumbStoreFile, err := writeFileToDisk(thumbImageUrl, dir, true)
		if err != nil {
			panic(err)
		}
		fmt.Printf("store image : %s\n", thumbStoreFile)

		originalImageUrl := result.OriginalImageUrl
		originStoreFile, err := writeFileToDisk(originalImageUrl, dir, false)
		if err != nil {
			panic(err)
		}
		fmt.Printf("store image : %s\n", originStoreFile)

		desc := []byte(result.ImageDesc)
		descFile, _ := filepath.Abs(filepath.Join(dir, "description.txt"))
		_ = os.WriteFile(descFile, desc, 0755)
		fmt.Printf("store image desc: %s\n", descFile)
	}
	return nil
}

func writeFileToDisk(url string, dir string, thumb bool) (StoreFile, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	u := uuid.New()
	fileName := strings.ReplaceAll(u.String(), "-", "")
	if thumb {
		fileName = fileName + "_thumb"
	}
	fileName = fileName + filepath.Ext(url)

	name, _ := filepath.Abs(filepath.Join(dir, fileName))
	file, err := os.Create(name)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return StoreFile(name), nil
}
