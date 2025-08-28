package index

import (
	"WikipediaImage/parse"
	"WikipediaImage/tool"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type IndexResult struct {
	Year         int
	Month        int
	ImageResults []parse.ImageResult
}

func (index *IndexResult) WriteIndex() error {
	pwd, _ := os.Getwd()

	dir, _ := filepath.Abs(filepath.Join(pwd, tool.RootDir))
	indexFile, _ := filepath.Abs(filepath.Join(dir, fmt.Sprintf("index_%d_%d.json", index.Year, index.Month)))

	content, err := json.MarshalIndent(index.ImageResults, "", "	")
	if err != nil {
		return err
	}
	err = os.WriteFile(indexFile, content, 0755)
	if err != nil {
		return err
	}
	fmt.Printf("store index file: %s\n", indexFile)

	return nil
}
