package main

import (
	"climber/getHots"
	"climber/webPublish"
	"fmt"
	"time"
)

func main() {
	fmt.Println("climer server start ...")

	timer := time.NewTimer(60 * time.Second)
	go func() {
		getHots.Collect()
		for {
			timer.Reset(60 * 5 * time.Second) // 这样来复用 timer 和修改执行时间
			select {
			case <-timer.C:
				getHots.Collect()
			}
		}
	}()
	webPublish.Publish()

}
