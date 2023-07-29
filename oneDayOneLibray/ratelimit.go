package oneDayOneLibray

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

func RateLimitTest() {
	rl := ratelimit.New(1000) // per second

	prev := time.Now()
	for i := 0; i < 1000; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
