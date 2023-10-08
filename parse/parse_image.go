package parse

import (
	. "WikipediaImage/tool"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"time"
)

type Parse struct {
	Year  int
	Month int
}

type ImageResult struct {
	Date              string
	ThumbImageUrl     string
	OriginalImageUrl  string
	OriginalImageLink string
	ImageDesc         string
}

const websiteUrl = "https://zh.wikipedia.org"
const websiteProtocol = "https"
const baseUrl = websiteUrl + "/wiki/Wikipedia:%E6%AF%8F%E6%97%A5%E5%9B%BE%E7%89%87/"

func (parse *Parse) ParseImage() ([]ImageResult, error) {
	url := baseUrl + FormatDate(parse.Year, parse.Month)

	client := http.Client{Timeout: time.Second * 15}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var rets []ImageResult
	doc.Find("td[width='50%']").Each(func(i int, selection *goquery.Selection) {
		node := selection.Find("div.thumb a")
		href, _ := node.Attr("href")
		title, _ := node.Attr("title")
		src, _ := node.Find("img").Attr("src")

		date := fmt.Sprintf("%d-%d-%d", parse.Year, parse.Month, i+1)
		ret := ImageResult{
			Date:              date,
			ThumbImageUrl:     websiteProtocol + ":" + src,
			ImageDesc:         title,
			OriginalImageLink: websiteUrl + href,
			OriginalImageUrl:  "",
		}

		foo, _ := parseOriginalImageUrl(ret.OriginalImageLink)
		ret.OriginalImageUrl = foo
		rets = append(rets, ret)

		fmt.Printf("parse image : %s\n", date)
	})

	return rets, nil
}

func parseOriginalImageUrl(url string) (string, error) {
	client := http.Client{Timeout: time.Second * 15}
	res, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	src, _ := doc.Find("a.internal").Attr("href")
	return websiteProtocol + ":" + src, nil
}
