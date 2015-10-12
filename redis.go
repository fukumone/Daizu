package main

import (
    "os"
    "fmt"
    "strings"
    "regexp"
    "io/ioutil"
    "github.com/PuerkitoBio/goquery"
    "github.com/garyburd/redigo/redis"
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
            fmt.Println(s2.Find("div.lineFi.clearfix").Text())
        })
    })
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

func main() {
    c, err := redis.Dial("tcp", ":6379")
    if err != nil {
        panic(err)
    }
    defer c.Close()


    // getNikkeiAve()
    date, err := ioutil.ReadFile("./code.txt")
    check(err)
    r := regexp.MustCompile(`[0-9]+`)

    for _, code := range strings.Split(string(date), "\n") {
        match_value := r.FindAllStringSubmatch(code, -1)
        securities_code := match_value[0][0]
        getPage(securities_code)
        //set
        c.Do("HSET", "bar", securities_code, "First bar")
        //get
        value, err := redis.String(c.Do("HGET", "bar", "1"))
        if err != nil {
            fmt.Println("key not found")
        }
        fmt.Println(value)
        //ENDINIT OMIT
    }
}

// To Do
// RedisにparseしたデータをHSETでsave
