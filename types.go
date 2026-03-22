package main

import "github.com/LeoYang92/gator/internal/config"

type state struct {
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	registerCommands map[string]func(*state, command) error
}
