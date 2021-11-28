package main

import (
	"fmt"
	"github.com/romycode/go-api-template/cmd/api/bootstrap"
	"log"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(fmt.Errorf("%w", err))
	}
}
