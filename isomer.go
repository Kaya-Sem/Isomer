package isomer

import (
	"errors"
	"fmt"
	"os"
)

// Command represents a CLI command with a name, description, and an action function.
type Command struct {
	Name        string
	Description string
	Action      func(args []string)
}

// Commander manages commands and dispatches them.
type Commander struct {
	commands map[string]Command
}

// NewCommander creates a new Commander instance.
func NewCommander() *Commander {
	return &Commander{
		commands: make(map[string]Command),
	}
}

// Register adds a new command to the commander.
func (c *Commander) Register(name, description string, action func(args []string)) {
	c.commands[name] = Command{
		Name:        name,
		Description: description,
		Action:      action,
	}
}

// Run parses and executes the appropriate command based on the input arguments.
func (c *Commander) Run(args []string) error {
	if len(args) < 1 {
		return errors.New("no command provided")
	}

	commandName := args[0]
	command, exists := c.commands[commandName]

	if !exists {
		return fmt.Errorf("unknown command: %s", commandName)
	}

	// Execute the command action with the provided arguments
	command.Action(args[1:])
	return nil
}

// ListCommands lists all available commands with their descriptions.
func (c *Commander) ListCommands() {
	fmt.Println("Available commands:")
	for _, cmd := range c.commands {
		fmt.Printf("  %s: %s\n", cmd.Name, cmd.Description)
	}
}

func (c *Commander) getCommands() map[string]Command {
	return c.commands
}

// ExecuteCommand handles command-line arguments directly from os.Args.
func (c *Commander) ExecuteCommand() {
	if err := c.Run(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
