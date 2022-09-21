package main

import (
	"fmt"

	integration "github.com/jearcila/hex-architecture/application/integration"
)

func main() {
	if err := integration.Run(); err != nil {
		fmt.Println(err)
	}
}
