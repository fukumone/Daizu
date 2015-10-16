package main

import (
    "os"
    "fmt"
    "strings"
    "regexp"
    "io/ioutil"
    "github.com/PuerkitoBio/goquery"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func getNikkeiAve() {
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

func getPage(securities_code string, file *os.File) {
    url := fmt.Sprintf("http://stocks.finance.yahoo.co.jp/stocks/detail/?code=%s", securities_code)
    doc, _ := goquery.NewDocument(url)
    name := fmt.Sprintf("%s\n", doc.Find("th.symbol").Text())
    file.WriteString(name)
    doc.Find("div#detail.marB6").Each(func(_ int, s1 *goquery.Selection) {
        s1.Find("div.innerDate").Each(func(_ int, s2 *goquery.Selection) {
            s2.Find("div.lineFi.clearfix > dl.tseDtlDelay").Each(func(_ int, s3 *goquery.Selection) {
                row := fmt.Sprintf("%s: %s\n", s3.Find("dt.title").Text(), s3.Find("dd.ymuiEditLink.mar0 > strong").Text())
                file.WriteString(row)
            })
        })
    })
    file.WriteString("\n")
}

func main() {
    getNikkeiAve()

    os.Remove("nikkei_info.txt")
    file, err := os.OpenFile("nikkei_info.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    check(err)
    defer file.Close()

    date, err := ioutil.ReadFile("./code.txt")
    check(err)

    r := regexp.MustCompile(`[0-9]+`)

    for _, code := range strings.Split(string(date), "\n") {
        match_value := r.FindAllStringSubmatch(code, -1)
        securities_code := match_value[0][0]
        getPage(securities_code, file)
    }
}
