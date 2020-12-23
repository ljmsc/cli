package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ljmsc/cli/cli"
)

// Possible command options:
// $ myapp start -t 5
// $ myapp version

func main() {
	app := cli.App{
		Name:    "myapp",
		Version: "1.0.0",
	}
	// app.RegisterCommandWithFlag("start", startAppHandler, cli.Flag{})
	app.RegisterCommand("test", "this is for testing", TestCommand{})

	if err := app.Run(); err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}

type TestCommand struct {
}

func (v TestCommand) Run(app *cli.App, param map[string]string) error {
	loops := 1
	if loopsText, ok := param["loops"]; ok {
		loops, _ = strconv.Atoi(loopsText)
	}

	for i := 0; i < loops; i++ {
		fmt.Printf("This is a test command for %s : %d \n", app.Name, i)
	}

	return nil
}
