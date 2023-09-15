package main

import (
	"fmt"
	"os"
	"github.com/MarsOccupied/plasma-core/plasma/bench"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bench <command>")
		fmt.Println("Commands: send, keys")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "send":
		bench.Send()
	case "keys":
		bench.Keys()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
