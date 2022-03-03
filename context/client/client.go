package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const port string = "8085"

var client http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("ERR while creating cookie jar %s", err.Error())
	}

	client = http.Client{
		Jar: jar,
	}
}

func main() {
	// test clientContext
	Run()

	req, err := http.NewRequest("GET", "http://localhost:"+port+"/doc", nil)
	if err != nil {
		log.Fatalf("ERR %s", err.Error())
	}

	cookie := &http.Cookie{
		Name:   "token",
		Value:  "some_token",
		MaxAge: 300,
	}

	urlObj, _ := url.Parse("http://localhost:" + port)
	client.Jar.SetCookies(urlObj, []*http.Cookie{cookie})

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERR is %s", err.Error())
	}
	defer resp.Body.Close()

	fmt.Printf("StatusCode: %d\n", resp.StatusCode)

	req, err = http.NewRequest("GET", "http://localhost:"+port+"/doc/id", nil)
	if err != nil {
		log.Fatalf("Got error %s", err.Error())
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}
	defer resp.Body.Close()

	fmt.Printf("StatusCode: %d\n", resp.StatusCode)

}
