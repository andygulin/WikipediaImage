package parse

import (
	"testing"
)

func TestParse_ParseImage(t *testing.T) {
	parse := Parse{Year: 2023, Month: 1}
	result, err := parse.ParseImage()
	if err != nil {
		panic(err)
	}
	for _, imageResult := range result {
		t.Log(imageResult.Date)
		t.Log(imageResult.ThumbImageUrl)
		t.Log(imageResult.OriginalImageUrl)
		t.Log()
	}
}

func TestParse_ParseImage2(t *testing.T) {
	url, err := parseOriginalImageUrl("https://zh.wikipedia.org/wiki/File:%D7%94%D7%A7%D7%A8%D7%95%D7%A1%D7%9C%D7%94_%D7%94%D7%A2%D7%91%D7%A8%D7%99%D7%AA_%D7%94%D7%A8%D7%90%D7%A9%D7%95%D7%A0%D7%94.jpg")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(url)
}
