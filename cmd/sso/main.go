package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("%+v\n", cfg)
}
