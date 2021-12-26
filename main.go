// concurrent.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var list []string
	start := time.Now()
	for i := 0; i < 25000; i++ {
		list = append(list, fmt.Sprintf("http://localhost:8000/test/%d", i+1))
	}
	request(list, 100, 1)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func call(url string, index int, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	secs := time.Since(start).Seconds()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ch <- fmt.Sprintf("%d : %.2f elapsed with response length: %d %s", index, secs, len(body), url)
}

func request(list []string, offset int, delay int) {
	page := 1
	for j := 0; j < len(list); j += offset {
		func() {
			end := j + offset
			if j+offset > len(list) {
				end = len(list)
			}
			slice := list[j:end]
			ch := make(chan string)
			for _, url := range slice {
				go call(url, j, ch)
			}
			for range slice {
				fmt.Println(<-ch)
			}
			fmt.Println("list page:", page)
		}()
		time.Sleep(time.Duration(delay) * time.Second)
		page++
	}
}
