package tests

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func fire_api(count int, wg *sync.WaitGroup) {
	hostname := "http://localhost:1729"
	urls := []string{"hello", "wrongurl"}
	url := urls[rand.Intn(2)]
	resp, err := http.Get(hostname + "/" + url)
	if err != nil {
		fmt.Println("Error while calling /", url)
		panic(err)
	}
	fmt.Printf("Called %d api and successful %s", count, resp.Status)
	wg.Done()
}
func Test() {
	start := time.Now()
	var wg sync.WaitGroup
	for j := 1; j <= 100; j++ {
		wg.Add(1)
		go fire_api(j, &wg)
	}
	wg.Wait()
	fmt.Println("\n\n\nTotal time taken", time.Since(start))
}
