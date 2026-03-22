package main

import "errors"

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registerCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
