package main

import (
	"fmt"
	"time"

	"github.com/jearcila/hex-architecture/application/integration"
)

func main() {
	if err := integration.Run(); err != nil {
		for {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		}
	}
}
