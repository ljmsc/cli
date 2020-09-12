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
	app.RegisterCommand("test", "this is for testing", TestCommmandBuilder)

	os.Exit(app.Run())
}

type TestCommand struct {
}

func TestCommmandBuilder() cli.Command {
	return TestCommand{}
}

func (v TestCommand) Run(param map[string]string) int {
	loops := 1
	if loopsText, ok := param["loops"]; ok {
		loops, _ = strconv.Atoi(loopsText)
	}

	for i := 0; i < loops; i++ {
		fmt.Printf("This is a test command %d \n", i)
	}

	return 0
}
