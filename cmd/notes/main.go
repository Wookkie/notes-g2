package main

import (
	"fmt"

	"github.com/Wookkie/notes-g2/internal/server"
)

func main() {

	server := server.New("0.0.0.0", "8080")
	server.Run()

	panic(fmt.Errorf("not implemented"))
}
