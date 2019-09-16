package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func getArticles() (map[string]Article, error) {

	var (
		rr struct {
			List map[string]Article `json:"list"`
		}
		err error
	)

	count := 1500
	offset := 0
	page := 1

	req, err := http.NewRequest("GET", "https://getpocket.com/v3/get", nil)
	if err != nil {
		return rr.List, err
	}

	q := req.URL.Query()
	q.Add("consumer_key", os.Getenv("POCKET_CONSUMER_KEY"))
	q.Add("access_token", os.Getenv("POCKET_ACCESS_TOKEN"))
	q.Add("detailType", "complete")
	q.Add("count", strconv.Itoa(count))
	q.Add("offset", strconv.Itoa(offset))
	q.Add("sort", "oldest")
	q.Add("state", "all")

	req.URL.RawQuery = q.Encode()

	log.Printf("Requesting page %d", page)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return rr.List, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return rr.List, errors.New("Invalid status")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rr.List, err
	}

	err = json.Unmarshal(body, &rr)
	if err != nil {
		return rr.List, err
	}

	if len(rr.List) == 0 {
		return rr.List, errors.New("No results")
	}

	return rr.List, nil

}

func getDomain(s string) (string, error) {

	u, err := url.Parse(s)
	return u.Hostname(), err

}

func insert(a Article) {

	var (
		domain string
		err    error
	)

	if a.ResolvedURL != "" {
		domain, err = getDomain(a.ResolvedURL)
		if err == nil {
			a.Domain = domain
		}
	}

	if a.Domain == "" {
		domain, err = getDomain(a.GivenURL)
		if err == nil {
			a.Domain = domain
		}
	}

	a.Domain = strings.ToLower(a.Domain)

	db.NewRecord(a)
	db.Create(&a)

}

func updateArticles() {

	aa, err := getArticles()
	if err != nil {
		log.Print(err)
		return
	}

	for _, a := range aa {
		log.Print("Inserting: ", a.GivenTitle)
		insert(a)
	}

	updateArticles()

}
