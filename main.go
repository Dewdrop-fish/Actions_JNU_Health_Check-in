package main

import (
	"fmt"
	"os"
)

func main() {
	sckey := os.Getenv("SCKEY")
	fmt.Println("sckey is "+sckey)
	}
