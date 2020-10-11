package cli

import "fmt"

type Command interface {
	Run(app *App, param map[string]string) int
}

type SimpleCommand struct {
	Command func(app *App, param map[string]string) error
}

func (s *SimpleCommand) Run(app *App, param map[string]string) int {
	if err := s.Command(app, param); err != nil {
		fmt.Printf("%s", err)
		return 1
	}
	return 0
}
