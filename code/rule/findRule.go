package rule

import (
	"github.com/PuerkitoBio/goquery"
)

type FindRule struct{}

func (p *FindRule) GetRule(doc *goquery.Document, f func(image string)) {
	doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
		if img, exist := s.Attr("href"); exist {
			f(img)
		}
	})
}
