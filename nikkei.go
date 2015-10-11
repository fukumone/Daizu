package main

import (
     "os"
     "github.com/PuerkitoBio/goquery"
)

func Exists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}

func main() {
    url := "https://indexes.nikkei.co.jp/nkave/index/component?idx=nk225"

    var file *os.File
    var err error

    if Exists("code.txt") {
        file, err = os.Open("code.txt")
    } else {
        file, err = os.Create("code.txt")
    }

    defer file.Close()

    if err != nil {
        panic(err)
    }

    doc, _ := goquery.NewDocument(url)
    doc.Find(".cmn-charcter").Text()
}
