package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://free-football-soccer-videos1.p.rapidapi.com/v1/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "free-football-soccer-videos1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "c90f9e04f1mshdacbcb552e04135p1a6cdajsn30aedbf99122")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
