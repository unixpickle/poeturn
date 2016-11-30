package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "poems.json count")
		os.Exit(1)
	}

	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid count:", os.Args[2])
		os.Exit(1)
	}

	poems := loadExisting(os.Args[1])
	m := map[string]bool{}
	for _, p := range poems {
		m[p] = true
	}

	log.Println("Fetching poem URLs...")
	urls := allPoemURLs()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error fetching URLs:", err)
		os.Exit(1)
	}

	for i := 0; i < count && len(urls) > 0; i++ {
		idx := rand.Intn(len(urls))
		url := urls[idx]
		urls[idx] = urls[len(urls)-1]
		urls = urls[:len(urls)-1]
		contents, err := PoemText(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Fetch failed:", err)
			break
		}
		if m[contents] {
			log.Println("Got duplicate.")
			continue
		}
		m[contents] = true
		poems = append(poems, contents)
		log.Println("Have", len(poems), "poems")
	}

	data, _ := json.Marshal(poems)
	if err := ioutil.WriteFile(os.Args[1], data, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to save:", err)
		os.Exit(1)
	}
}

func loadExisting(file string) []string {
	contents, err := ioutil.ReadFile(file)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		fmt.Fprintln(os.Stderr, "Read existing data:", err)
		os.Exit(1)
	}
	var res []string
	if err := json.Unmarshal(contents, &res); err != nil {
		fmt.Fprintln(os.Stderr, "Unmarshal existing data:", err)
		os.Exit(1)
	}
	return res
}

func allPoemURLs() []string {
	poets, err := TopPoetNames()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fetch poets:", err)
		os.Exit(1)
	}

	resChan := make(chan string, 1)
	wg := sync.WaitGroup{}
	go func() {
		wg.Wait()
		close(resChan)
	}()

	for _, p := range poets {
		wg.Add(1)
		go func(poet string) {
			defer wg.Done()
			urls, err := Poems(poet)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Fetch poems:", err)
				os.Exit(1)
			}
			for _, u := range urls {
				resChan <- u
			}
		}(p)
	}

	var res []string
	for r := range resChan {
		res = append(res, r)
	}
	return res
}
