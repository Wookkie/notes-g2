package main

import (
	"fmt"

	"github.com/Wookkie/notes-g2/internal"

	inmemory "github.com/Wookkie/notes-g2/internal/repository/in-memory"
	"github.com/Wookkie/notes-g2/internal/server"
)

func main() {
	cfg := internal.ReadConfig()
	fmt.Printf("Host: %s\nPort: %d\n", cfg.Host, cfg.Port)

	inMemoryRepo := inmemory.New()

	notesAPI := server.New(cfg, inMemoryRepo)

	notesAPI.Run()
}
