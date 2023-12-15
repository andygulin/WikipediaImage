package index

import (
	"WikipediaImage/parse"
	"WikipediaImage/tool"
	"fmt"
	"strings"
	"testing"
)

func TestIndexResult_WriteIndex(t *testing.T) {
	var rets []parse.ImageResult
	for i := 0; i <= 10; i++ {
		idx := i + 1
		rets = append(rets, parse.ImageResult{
			Date:              fmt.Sprintf("%d-%d-%d", 2023, 1, idx),
			ThumbImageUrl:     fmt.Sprintf("ThumbImageUrl-%d", idx),
			ThumbImageFile:    fmt.Sprintf("ThumbImageFile-%d", idx),
			OriginalImageUrl:  fmt.Sprintf("OriginalImageUrl-%d", idx),
			OriginalImageFile: fmt.Sprintf("OriginalImageFile-%d", idx),
			OriginalImageLink: fmt.Sprintf("OriginalImageLink-%d", idx),
			ImageDesc:         fmt.Sprintf("ImageDesc-%d", idx),
		})
	}

	index := IndexResult{Year: 2023, Month: 1, ImageResults: rets}
	err := index.WriteIndex()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestIndexResult_WriteIndex2(t *testing.T) {
	str := "/Users/gulin/Documents/project/gulin/WikipediaImage/store_image/2023-11-1/xxxxxx.jpg"
	idx := strings.Index(str, tool.RootDir) + len(tool.RootDir) + 1
	t.Log(idx)
	t.Log(str[idx:])
}
