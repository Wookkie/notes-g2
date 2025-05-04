package main

import (
	"fmt"

	"github.com/Wookkie/notes-g2/internal"
)

func main() {
	cfg := internal.ReadConfig()

	fmt.Printf("Host: %s\nPort: %d\n", cfg.Host, cfg.Port)
}
