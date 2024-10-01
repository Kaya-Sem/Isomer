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

// DefaultHandler represents a default action based on the number of arguments.
type DefaultHandler struct {
	Name        string
	Description string
	Action      func(args []string)
}

// Commander manages commands and dispatches them.
type Commander struct {
	commands        map[string]Command
	defaultHandlers map[int]DefaultHandler
}

// NewCommander creates a new Commander instance.
func NewCommander() *Commander {
	return &Commander{
		commands:        make(map[string]Command),
		defaultHandlers: make(map[int]DefaultHandler),
	}
}

// RegisterNamedCommand adds a new named command to the commander.
func (c *Commander) RegisterNamedCommand(name, description string, action func(args []string)) {
	c.commands[name] = Command{
		Name:        name,
		Description: description,
		Action:      action,
	}
}

// RegisterDefaultHandler registers a default handler based on the number of arguments.
func (c *Commander) RegisterDefaultHandler(argCount int, name, description string, action func(args []string)) {
	c.defaultHandlers[argCount] = DefaultHandler{
		Name:        name,
		Description: description,
		Action:      action,
	}
}

// Run parses and executes the appropriate command based on the input arguments.
func (c *Commander) Run(args []string) error {
	if len(args) < 1 {
		return errors.New("no arguments provided")
	}

	commandName := args[0]
	command, exists := c.commands[commandName]

	// If a command exists, execute it
	if exists {
		command.Action(args[1:])
		return nil
	}

	// If no command is found, treat args as a default operation
	return c.defaultOperation(args)
}

// defaultOperation handles arguments if no command matches
func (c *Commander) defaultOperation(args []string) error {
	argCount := len(args)

	// Check if a default handler exists for the given number of arguments
	if handler, exists := c.defaultHandlers[argCount]; exists {
		handler.Action(args)
		return nil
	}

	return fmt.Errorf("no command or default handler for %d arguments", argCount)
}

// ListCommands lists all available commands with their descriptions.
func (c *Commander) ListCommands() {
	fmt.Println("Available commands:")
	for _, cmd := range c.commands {
		fmt.Printf("  %s: %s\n", cmd.Name, cmd.Description)
	}

	fmt.Println("Default handlers:")
	for argCount, handler := range c.defaultHandlers {
		fmt.Printf("  Default for %d args: %s - %s\n", argCount, handler.Name, handler.Description)
	}
}

// ExecuteCommand handles command-line arguments directly from os.Args.
func (c *Commander) ExecuteCommand() {
	if err := c.Run(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
