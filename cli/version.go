package cli

import (
	"fmt"
	"runtime"
)

type VersionCommand struct {
}

func (v VersionCommand) Run(app *App, param map[string]string) int {
	fmt.Printf("%s version %s %s/%s", app.Name, app.Version, runtime.GOOS, runtime.GOARCH)
	return 0
}
