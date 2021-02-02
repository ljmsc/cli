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

func (a *App) parseArgs() error {
	a.parsedParameters = make(map[string]string)
	a.parsedCommands = make([]string, 0, len(a.args))
	isParamVal := false
	paramName := ""
	for _, arg := range a.args {
		// check if it is a parameter
		if strings.HasPrefix(arg, "-") {
			if equalsOp := strings.Split(arg, "="); len(equalsOp) == 2 {
				if len(equalsOp[0]) > 1 {
					a.parsedParameters[equalsOp[0][1:]] = equalsOp[1]
				}
				continue
			}

			paramName = arg[1:]
			isParamVal = true
			a.parsedParameters[paramName] = ""
			continue
		}

		// if previous one was a parameter this one might be the value
		if isParamVal {
			a.parsedParameters[paramName] = arg
			isParamVal = false
			continue
		}

		a.parsedCommands = append(a.parsedCommands, arg)
	}

	return nil
}

func (a *App) help() {
	for name, command := range a.registeredCommands {
		fmt.Printf("%s\t\t%s\n", name, command.Help(a))
	}
}

func (a *App) RegisterCommandFunc(name string, commandFunc CommandFunc) {
	a.RegisterCommand(name, &commandFunc)
}

func (a *App) RegisterCommand(name string, command Command) {
	if len(name) == 0 {
		panic("name can't be empty")
	}
	if a.registeredCommands == nil {
		a.registeredCommands = make(map[string]Command)
	}
	a.registeredCommands[name] = command
}

func (a *App) Run() error {
	if a.args == nil {
		a.args = os.Args[1:]
	}

	if _, ok := a.registeredCommands["version"]; !ok {
		a.RegisterCommand("version", versionCommand)
	}

	if err := a.parseArgs(); err != nil {
		return fmt.Errorf("could not parse parameters: %w", err)
	}

	if len(a.parsedCommands) == 0 {
		a.help()
		return nil
	}

	firstCommand := a.parsedCommands[0]
	if _, ok := a.registeredCommands[firstCommand]; !ok {
		// command not found
		return fmt.Errorf("unknown command: %s", firstCommand)
	}

	command := a.registeredCommands[firstCommand]
	return command.Run(a, a.parsedParameters)
}
