package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(fmt.Sprintf("execution at %s", now))
}
