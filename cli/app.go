package cli

import (
	"fmt"
	"os"
	"strings"
)

type App struct {
	Name               string
	Version            string
	args               []string
	registeredCommands map[string]Command
	parsedParameters   map[string]string
	parsedCommands     []string
}

func (c *App) parseArgs() error {
	c.parsedParameters = make(map[string]string)
	c.parsedCommands = make([]string, 0, len(c.args))
	isParamVal := false
	paramName := ""
	for _, arg := range c.args {
		// check if it is a parameter
		if strings.HasPrefix(arg, "-") {
			if equalsOp := strings.Split(arg, "="); len(equalsOp) == 2 {
				c.parsedParameters[equalsOp[0]] = equalsOp[1]
				continue
			}

			paramName = arg[1:]
			isParamVal = true
			c.parsedParameters[paramName] = ""
			continue
		}

		// if previous one was a parameter this one might be the value
		if isParamVal {
			c.parsedParameters[paramName] = arg
			isParamVal = false
			continue
		}

		c.parsedCommands = append(c.parsedCommands, arg)
	}

	return nil
}

func (c *App) RegisterCommand(name string, desc string, comBuilder func() Command) {
	if c.registeredCommands == nil {
		c.registeredCommands = make(map[string]Command)
	}
	c.registeredCommands[name] = comBuilder()
}

func (c *App) Run() int {
	if c.args == nil {
		c.args = os.Args[1:]
	}
	if err := c.parseArgs(); err != nil {
		fmt.Printf("could not parse parameters: %s \n", err.Error())
		return 1
	}

	if len(c.parsedCommands) == 0 {
		// render help text
		return 0
	}

	firstCommand := c.parsedCommands[0]
	if _, ok := c.registeredCommands[firstCommand]; !ok {
		// command not found
		fmt.Printf("unknown command: %s", firstCommand)
		return 1
	}

	command := c.registeredCommands[firstCommand]
	return command.Run(c.parsedParameters)
}
