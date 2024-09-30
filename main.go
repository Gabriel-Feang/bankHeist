package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	isHeistOn := true
	guards := rand.Intn(100)

	if guards >= 50 {
		fmt.Println("Heh, escaped guards.")
	} else {
		isHeistOn = false
	}
	fmt.Println(isHeistOn)
}
