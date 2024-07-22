package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/lilpipidron/share-botherer/internal/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := config.Load()
	fmt.Println(*cfg)
}
