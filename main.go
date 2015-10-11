package main

import (
     "fmt"
     "github.com/PuerkitoBio/goquery"
)

func main() {
    url := "http://stocks.finance.yahoo.co.jp/stocks/detail/?code=2001.T"
    doc, _ := goquery.NewDocument(url)
    doc.Find("div#detail.marB6").Each(func(_ int, s1 *goquery.Selection) {
        s1.Find("div.innerDate").Each(func(_ int, s2 *goquery.Selection) {
            s2.Find("div.lineFi.clearfix").Each(func(_ int, s3 *goquery.Selection) {
                fmt.Print(s3.Text())
            })
        })
    })
}
