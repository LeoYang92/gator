package main

import (
	"fmt"
	"log"

	"github.com/LeoYang92/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	cfg.SetUser("lane")
	fmt.Println(cfg.DBURL)
}
