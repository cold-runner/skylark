package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(uuid.New())
	}
}
