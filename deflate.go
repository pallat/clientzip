package main

import (
	"compress/flate"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8080/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept-Encoding", "flate")

	c := http.DefaultClient

	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := flate.NewReader(res.Body)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}
