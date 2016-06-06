package rule

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type MeiRule struct{}

func (p *MeiRule) UrlRule() string {
	return "https://meizi.us/"
}

func (p *MeiRule) PageRule(currentPage int) (page string) {
	return "?page=" + strconv.Itoa(currentPage)
}

func (p *MeiRule) ImageRule(doc *goquery.Document, f func(image string)) {
	doc.Find("img.img-responsive").Each(func(i int, s *goquery.Selection) {
		if img, exist := s.Attr("src"); exist {
			f(img)
		}
	})
}
