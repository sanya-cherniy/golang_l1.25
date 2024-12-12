package main

import (
	"time"
)

func main() {
	my_sleep(time.Millisecond * 5000)
	println("gg")
}

func my_sleep(d time.Duration) {
	<-time.After(d)
}
