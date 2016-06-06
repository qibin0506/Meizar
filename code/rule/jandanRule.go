package rule

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type JandanRule struct{}

func (p *JandanRule) UrlRule() (url string) {
	return "http://jandan.net/ooxx/"
}

func (p *JandanRule) PageRule(currentPage int) (page string) {
	return "page-" + strconv.Itoa(currentPage)
}

func (p *JandanRule) ImageRule(doc *goquery.Document, f func(image string)) {
	doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
		if img, exist := s.Attr("href"); exist {
			f(img)
		}
	})
}
