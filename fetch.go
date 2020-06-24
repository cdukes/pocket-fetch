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

func populateArticles(aa map[string]Article, page int) error {
	var (
		rr struct {
			List map[string]Article `json:"list,omitempty"`
		}
		err error
	)

	count := 1500

	req, err := http.NewRequest("GET", "https://getpocket.com/v3/get", nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("consumer_key", os.Getenv("POCKET_CONSUMER_KEY"))
	q.Add("access_token", os.Getenv("POCKET_ACCESS_TOKEN"))
	q.Add("detailType", "complete")
	q.Add("count", strconv.Itoa(count))
	q.Add("offset", strconv.Itoa(count*(page-1)))
	q.Add("sort", "oldest")
	q.Add("state", "all")

	req.URL.RawQuery = q.Encode()

	log.Printf("Requesting page %d", page)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Invalid status")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &rr)
	if err != nil {
		return err
	}

	for _, a := range rr.List {
		aa[a.ID] = a
	}

	if len(rr.List) < count {
		return nil
	}

	return populateArticles(aa, page+1)
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
	aa := make(map[string]Article)

	err := populateArticles(aa, 1)
	if err != nil {
		log.Print(err)
		return
	}

	for _, a := range aa {
		count := 0

		db.Model(Article{}).Where("id = ?", a.ID).Count(&count)

		if count > 0 {
			continue
		}

		log.Print("Inserting: ", a.GivenTitle)
		insert(a)
	}
}
