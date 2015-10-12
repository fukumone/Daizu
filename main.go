package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "regexp"
    "github.com/PuerkitoBio/goquery"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func getPage(securities_code string) {
    url := fmt.Sprintf("http://stocks.finance.yahoo.co.jp/stocks/detail/?code=%s", securities_code)
    doc, _ := goquery.NewDocument(url)
    doc.Find("div#detail.marB6").Each(func(_ int, s1 *goquery.Selection) {
        s1.Find("div.innerDate").Each(func(_ int, s2 *goquery.Selection) {
            s2.Find("div.lineFi.clearfix").Each(func(_ int, s3 *goquery.Selection) {
                fmt.Print(s3.Text())
            })
        })
    })
}

func main() {
    date, err := ioutil.ReadFile("./code.txt")
    check(err)
    r := regexp.MustCompile(`[0-9]+`)

    for _, code := range strings.Split(string(date), "\n") {
        match_value := r.FindAllStringSubmatch(code, -1)
        securities_code := match_value[0][0]
        getPage(securities_code)
    }
}
