package cli

type Command interface {
	Run(app *App, param map[string]string) int
}
