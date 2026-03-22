package main

import (
	"log"
	"os"

	"github.com/LeoYang92/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	sta := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registerCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(sta, command{Name: cmdName, Args: cmdArgs})

	if err != nil {
		log.Fatal(err)
	}

}
