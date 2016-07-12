package main

import (
	"fmt"

	"golang.org/x/net/context"
	_ "golang.org/x/net/context/ctxhttp"
)

func main() {
	fmt.Printf("Hello, World from %v\n", context.TODO())
}
