package stat

import (
	"fmt"
	"time"
)

var StatChan = make(chan (int))
var ticker = time.NewTicker(time.Second)

func CollectStat() {
	var requestCount int
	for {
		select {
		case <-StatChan:
			requestCount++
		case <-ticker.C:
			fmt.Println("Requests per second:", requestCount)
			requestCount = 0
		}
	}
}
