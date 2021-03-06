package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func getIDs() (ids []int, err error) {
	resp, err := http.Get("http://localhost:9874/people/ids")
	if err != nil {
		return nil, fmt.Errorf("error fetching ids: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	err = json.Unmarshal(body, &ids)
	return
}

func getPerson(id int) (person, error) {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return person{}, fmt.Errorf("error fetching person: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return person{}, fmt.Errorf("error fetching URL (%s) return status code %d", url, resp.StatusCode)
	}
	var p person
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return person{}, fmt.Errorf("error parsing person: %v", err)
	}
	return p, nil
}

type person struct {
	ID           int
	GivenName    string
	FamilyName   string
	StartDate    time.Time
	Rating       int
	FormatString string
}

func (p person) String() string {
	return fmt.Sprintf("%s %s", p.GivenName, p.FamilyName)
}

func main() {
	start := time.Now()
	ids, err := getIDs()
	if err != nil {
		log.Fatalf("getIDs failed: %v", err)
	}
	fmt.Println(ids)

	var wg sync.WaitGroup
	ch := make(chan person, 10)
	for _, n := range ids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			p, err := getPerson(id)
			if err != nil {
				log.Printf("%d: %v\n", id, err)
				return
			}
			ch <- p
		}(n)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for p := range ch {
		fmt.Printf("%d: %v\n", p.ID, p)
	}

	elapsed := time.Since(start)
	fmt.Printf("Total time: %s\n", elapsed)
}
