package rule

import (
	"github.com/PuerkitoBio/goquery"
)

type Rule interface {
	UrlRule() (url string)
	PageRule(currentPage int) (page string)
	ImageRule(doc *goquery.Document, f func(image string))
}
