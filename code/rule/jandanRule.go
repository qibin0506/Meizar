package rule

import (
    // "fmt"
    "github.com/PuerkitoBio/goquery"
    // "github.com/robertkrimen/otto"
    "encoding/base64"
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
    // doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
    //     if img, exist := s.Attr("href"); exist {
    //         f(img)
    //     }
    // })

    doc.Find("span.img-hash").Each(func(i int, s *goquery.Selection) {
        hash_value := s.Text()
        decoded, err := base64.StdEncoding.DecodeString(hash_value)
        if err == nil {
            img := "http:" + string(decoded)
            // fmt.Println(hash_value + "->" + img)
            f(img)
        }
    })
}
