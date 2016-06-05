package rule

import (
	"github.com/PuerkitoBio/goquery"
)

type Rule interface {
	GetRule(doc *goquery.Document, f func(image string))
}
