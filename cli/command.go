package cli

import (
	"fmt"
	"runtime"
)

type Command interface {
	Run(app *App, param map[string]string) error
	Help(app *App) Help
}

type CommandFunc func(app *App, param map[string]string) error

func (s CommandFunc) Help(app *App) Help {
	return Help{}
}

func (s CommandFunc) Run(app *App, param map[string]string) error {
	return s(app, param)
}

var versionCommand CommandFunc = func(app *App, param map[string]string) error {
	fmt.Printf("%s version %s %s/%s", app.Name, app.Version, runtime.GOOS, runtime.GOARCH)
	return nil
}
