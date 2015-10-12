package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    url := "https://indexes.nikkei.co.jp/nkave/index/component?idx=nk225"

    os.Remove("code.txt")
    file, err := os.OpenFile("code.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    check(err)

    defer file.Close()
    doc, _ := goquery.NewDocument(url)
    doc.Find(".cmn-table > tbody").Each(func(_ int, ele1 *goquery.Selection) {
        doc.Find("tr.cmn-charcter").Each(func(_ int, ele2 *goquery.Selection) {
        code := strings.TrimSpace(ele2.Find("td.cmn-stock_border").First().Text())
        brand := strings.TrimSpace(ele2.Find("td.cmn-stock_border a").Text())
        name := strings.TrimSpace(ele2.Find("td.cmn-stock_border").Last().Text())
        row := fmt.Sprintf("%s, %s, %s\n", code, brand, name)
        file.WriteString(row)
        })
    })
}
