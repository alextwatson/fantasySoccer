package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
)
// this thing working???

func main() {

        url := "https://stroccoli-futbol-science-v1.p.rapidapi.com/s1/stats/6cf4a98226ebece13a9fa993cd8bf1a6a1011e27cf1f47862ccc3dcc"

        req, _ := http.NewRequest("GET", url, nil)

        req.Header.Add("x-rapidapi-host", "stroccoli-futbol-science-v1.p.rapidapi.com")
        req.Header.Add("x-rapidapi-key", "c90f9e04f1mshdacbcb552e04135p1a6cdajsn30aedbf99122")

        res, _ := http.DefaultClient.Do(req)

        defer res.Body.Close()
        body, _ := ioutil.ReadAll(res.Body)

        fmt.Println(res)
        fmt.Println(string(body))

}
