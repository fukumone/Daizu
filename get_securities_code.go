package main

import (
    "os"
    "fmt"
    "github.com/PuerkitoBio/goquery"
)

func Exists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}

type Company struct {
    code  string
    brand string
    name  string
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    url := "https://indexes.nikkei.co.jp/nkave/index/component?idx=nk225"

    var file *os.File
    var err error

    if Exists("code.txt") {
        file, err = os.Open("code.txt")
        check(err)
    } else {
        file, err = os.Create("code.txt")
        check(err)
    }

    defer file.Close()

    doc, _ := goquery.NewDocument(url)
    doc.Find(".cmn-charcter").Each(func(_ int, ele *goquery.Selection) {
        code := ele.Find("td.cmn-stock_border").Text()
        brand := ele.Find("td.cmn-character cmn-stock_border a").Text()
        name := ele.Find("td.cmn-character cmn-stock_border").Last().Text()

        data := Company{ code, brand, name, }
        fmt.Println(data)
    })
}
