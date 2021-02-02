package cli

import (
	"fmt"
	"runtime"
)

type Command interface {
	Run(app *App, param map[string]string) error
	Help(app *App) string
}

type CommandFunc func(app *App, param map[string]string) error

func (s CommandFunc) Help(app *App) string {
	return ""
}

func (s CommandFunc) Run(app *App, param map[string]string) error {
	return s(app, param)
}

var versionCommand CommandFunc = func(app *App, param map[string]string) error {
	fmt.Printf("%s version %s %s/%s", app.Name, app.Version, runtime.GOOS, runtime.GOARCH)
	return nil
}
