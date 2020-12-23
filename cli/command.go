package cli

type Command interface {
	Run(app *App, param map[string]string) error
}

type SimpleCommand struct {
	Command func(app *App, param map[string]string) error
}

func (s *SimpleCommand) Run(app *App, param map[string]string) error {
	return s.Command(app, param)
}
