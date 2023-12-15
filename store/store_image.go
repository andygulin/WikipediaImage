package store

import (
	. "WikipediaImage/parse"
	. "WikipediaImage/tool"
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

func (store *Store) StoreImage(imageResults []ImageResult) ([]ImageResult, error) {
	var rets []ImageResult

	pwd, _ := os.Getwd()

	for _, result := range imageResults {
		dir, _ := filepath.Abs(filepath.Join(pwd, RootDir, result.Date))
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

		ret := result

		idx := strings.Index(string(thumbStoreFile), RootDir) + len(RootDir) + 1
		ret.ThumbImageFile = string(thumbStoreFile)[idx:]

		idx = strings.Index(string(originStoreFile), RootDir) + len(RootDir) + 1
		ret.OriginalImageFile = string(originStoreFile)[idx:]
		rets = append(rets, ret)

		fmt.Printf("store image : %s\n", originStoreFile)
	}
	return rets, nil
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
